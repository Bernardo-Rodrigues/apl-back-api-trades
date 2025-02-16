package infra

import (
	"app/core/controller"
	"app/core/use-case/generate-report"
	"app/infra/adapters/grpc"
	"app/infra/adapters/grpc/services/report"
)

func Start() {
	//DI
	reportUsecase := generate_report.New()
	reportController := controller.New(reportUsecase)
	reportService := report.New(reportController)
	server := grpc.NewServer(reportService)

	server.Serve()
}
