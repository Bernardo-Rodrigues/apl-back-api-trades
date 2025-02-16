package report

import (
	gen "app/infra/adapters/grpc/services/report/gen"
	"bytes"
	"context"
	"fmt"
	"io"
	"time"
)

func (s *reportService) GenerateReport(ctx context.Context, req *gen.ReportRequest) (*gen.ReportResponse, error) {
	start := time.Now()
	layout := "2006-01-02 15:04:05"

	startDate, _ := time.Parse(layout, req.StartDate)
	endDate, _ := time.Parse(layout, req.EndDate)

	tradesFile := bytes.NewReader(req.TradesFile)
	assetsFiles := make(map[string]io.Reader)
	for fileName, content := range req.AssetsFiles {
		assetsFiles[fileName] = bytes.NewReader(content)
	}

	s.reportController.GenerateReport(startDate, endDate, int(req.IntervalMinutes), float64(req.InitialBalance), tradesFile, assetsFiles)

	fmt.Printf("Execution time (microseconds): %d µs\n", time.Since(start).Microseconds())
	return &gen.ReportResponse{
		Message: "Report generated successfully!",
	}, nil
}
