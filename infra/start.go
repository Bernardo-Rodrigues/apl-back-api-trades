package infra

import (
	"app/core/controller"
	"app/core/use-case/generate-report"
	"app/infra/adapters/grpc"
	"app/infra/adapters/grpc/services/report"
	"app/infra/adapters/rest"
	"app/infra/adapters/rest/services/report"
)

func Start() {
	//DI
	reportUsecase := generate_report.New()
	reportController := controller.New(reportUsecase)

	grpcReportService := grpc_services_report.New(reportController)
	grpcServer := grpc.NewServer(grpcReportService)
	go grpcServer.Serve()

	restReportService := rest_services_report.New(reportController)
	restServer := rest.NewServer(restReportService)
	restServer.Serve()
}
