package rest_services_report

import (
	"app/infra/adapters/files"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"strconv"
	"strings"
	"time"
)

func (s *reportService) GenerateReport(ctx *fasthttp.RequestCtx) {
	start := time.Now()

	intervalMinutes, _ := strconv.Atoi(string(ctx.FormValue("interval_minutes")))
	initialBalance, _ := strconv.ParseFloat(string(ctx.FormValue("initial_balance")), 64)
	startDate, endDate, err := getDates(ctx)
	if err != nil {
		return
	}
	tradesFile, assetsFiles, fileType, err := getFiles(ctx)
	if err != nil {
		return
	}

	filesHandler, err := files.NewHandler(fileType, tradesFile, assetsFiles)
	if err != nil {
		return
	}
	s.reportController.SetFilesHandler(filesHandler)

	fileData, err := s.reportController.GenerateReport(*startDate, *endDate, intervalMinutes, initialBalance)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte(fmt.Sprintf(`{"error": "error generating report: %v"}`, err)))
		return
	}

	ctx.Response.Header.Set("Content-Type", "text/csv")
	ctx.Response.Header.Set("Content-Disposition", `attachment; filename="report.csv"`)
	ctx.Response.SetBody(fileData)
	ctx.SetStatusCode(fasthttp.StatusOK)

	fmt.Printf("Execution time (microseconds): %d µs\n", time.Since(start).Microseconds())
}

func getDates(ctx *fasthttp.RequestCtx) (*time.Time, *time.Time, error) {
	layout := "2006-01-02 15:04:05"

	startDateStr := string(ctx.FormValue("start_date"))
	endDateStr := string(ctx.FormValue("end_date"))

	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte(`{"error": "invalid start date format"}`))
		return nil, nil, err
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte(`{"error": "invalid end date format"}`))
		return nil, nil, err
	}

	return &startDate, &endDate, nil
}

func getFiles(ctx *fasthttp.RequestCtx) (io.Reader, map[string]io.Reader, string, error) {
	tradesHeader, err := ctx.FormFile("trades_file")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte(`{"error": "trades_file is required"}`))
		return nil, nil, "", err
	}
	tradesFile, err := tradesHeader.Open()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte(`{"error": "failed to read trades file"}`))
		return nil, nil, "", err
	}

	assetsFiles := make(map[string]io.Reader)
	if form, err := ctx.MultipartForm(); err == nil {
		for fieldName, fileHeaders := range form.File {
			if strings.HasPrefix(fieldName, "assets_files[") {
				assetName := strings.TrimSuffix(strings.TrimPrefix(fieldName, "assets_files["), "]")

				fileHeader := fileHeaders[0]
				file, _ := fileHeader.Open()
				assetsFiles[assetName] = file
			}
		}
	}

	fiilesType := files.DetectFileType(tradesFile)
	tradesFile.Seek(0, io.SeekStart)

	return tradesFile, assetsFiles, fiilesType, nil
}
