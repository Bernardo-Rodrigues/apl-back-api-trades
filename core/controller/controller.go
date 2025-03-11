package controller

import (
	"app/core/use-case/generate-report"
	"app/infra/adapters/files"
	"time"
)

type ReportController interface {
	GenerateReport(startDate, endDate time.Time, intervalMinutes int, initialBalance float64) ([]byte, error)
	SetFilesHandler(filesHandler files.FilesHandler)
}

type reportController struct {
	usecase      generate_report.GenerateReportUsecase
	filesHandler files.FilesHandler
}

func New(usecase generate_report.GenerateReportUsecase) *reportController {
	return &reportController{
		usecase: usecase,
	}
}

func (c *reportController) SetFilesHandler(filesHandler files.FilesHandler) {
	c.filesHandler = filesHandler
}
