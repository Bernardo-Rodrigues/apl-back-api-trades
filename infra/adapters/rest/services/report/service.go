package rest_services_report

import (
	"app/core/controller"
	"github.com/valyala/fasthttp"
)

type ReportService interface {
	GenerateReport(ctx *fasthttp.RequestCtx)
}

type reportService struct {
	reportController controller.ReportController
}

func New(reportController controller.ReportController) *reportService {
	return &reportService{
		reportController: reportController,
	}
}
