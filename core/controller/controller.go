package controller

import (
	"app/core/use-case/generate-report"
	"io"
	"time"
)

type ReportController interface {
	GenerateReport(startDate, endDate time.Time, intervalMinutes int, initialBalance float64, tradesFile io.Reader, assetsFiles map[string]io.Reader) (string, error)
}

type reportController struct {
	usecase generate_report.GenerateReportUsecase
}

func New(usecase generate_report.GenerateReportUsecase) *reportController {
	return &reportController{
		usecase: usecase,
	}
}
