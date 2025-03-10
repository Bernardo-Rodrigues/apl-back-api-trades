package grpc_services_report

import (
	"app/infra/adapters/files"
	gen "app/infra/adapters/grpc/services/report/gen"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"time"
)

func (s *grpcReportService) GenerateReport(ctx context.Context, req *gen.ReportRequest) (*gen.ReportResponse, error) {
	start := time.Now()

	var startDate, endDate time.Time

	if err := validateRequest(req, &startDate, &endDate); err != nil {
		return nil, fmt.Errorf("error validating request: %w", err)
	}

	tradesFile := bytes.NewReader(req.TradesFile)

	assetsFiles := make(map[string]io.Reader)
	for fileName, content := range req.AssetsFiles {
		if len(content) == 0 {
			return nil, fmt.Errorf("assets_file %s is empty", fileName)
		}
		assetsFiles[fileName] = bytes.NewReader(content)
	}

	fileType := files.DetectFileType(tradesFile)
	tradesFile.Seek(0, io.SeekStart)

	filesHandler, err := files.NewHandler(fileType, tradesFile, assetsFiles)
	if err != nil {
		return nil, fmt.Errorf("error creating file handler: %w", err)
	}
	s.reportController.SetFilesHandler(filesHandler)

	fileData, err := s.reportController.GenerateReport(startDate, endDate, int(req.IntervalMinutes), float64(req.InitialBalance))
	if err != nil {
		return nil, fmt.Errorf("error generating report: %w", err)
	}

	fmt.Printf("Execution time (microseconds): %d µs\n", time.Since(start).Microseconds())

	return &gen.ReportResponse{
		Message: "Report generated successfully!",
		File:    fileData,
	}, nil
}

func validateRequest(req *gen.ReportRequest, startDate, endDate *time.Time) error {
	var err error

	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.StartDate == "" || req.EndDate == "" {
		return errors.New("start_date and end_date are required")
	}

	layout := "2006-01-02 15:04:05"
	*startDate, err = time.Parse(layout, req.StartDate)
	if err != nil {
		return errors.New("invalid start_date format, expected 'YYYY-MM-DD HH:MM:SS'")
	}

	*endDate, err = time.Parse(layout, req.EndDate)
	if err != nil {
		return errors.New("invalid end_date format, expected 'YYYY-MM-DD HH:MM:SS'")
	}

	if !endDate.After(*startDate) {
		return errors.New("end_date must be after start_date")
	}

	if req.IntervalMinutes <= 0 {
		return errors.New("interval_minutes must be greater than zero")
	}

	if req.InitialBalance < 0 {
		return errors.New("initial_balance cannot be negative")
	}

	if len(req.TradesFile) == 0 {
		return errors.New("trades_file is required")
	}

	return nil
}
