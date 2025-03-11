package grpc_services_report

import (
	"app/core/controller"
	"app/infra/adapters/grpc/services/report/gen"
)

type grpcReportService struct {
	gen.UnimplementedReportServiceServer
	reportController controller.ReportController
}

func New(reportController controller.ReportController) *grpcReportService {
	return &grpcReportService{
		reportController: reportController,
	}
}
