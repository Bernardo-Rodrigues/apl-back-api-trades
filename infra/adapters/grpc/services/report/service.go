package report

import (
	"app/core/controller"
	"app/infra/adapters/grpc/services/report/gen"
)

type reportService struct {
	gen.UnimplementedReportServiceServer
	reportController controller.ReportController
}

func New(reportController controller.ReportController) *reportService {
	return &reportService{
		reportController: reportController,
	}
}
