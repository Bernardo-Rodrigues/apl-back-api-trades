package rest

import (
	"app/infra/adapters/rest/services/report"
)

type restServer struct {
	reportService rest_services_report.ReportService
}

func NewServer(reportService rest_services_report.ReportService) *restServer {
	return &restServer{
		reportService: reportService,
	}
}
