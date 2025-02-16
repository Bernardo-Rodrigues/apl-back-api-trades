package grpc

import (
	"app/infra/adapters/grpc/services/report/gen"
	"google.golang.org/grpc"
)

type grpcServer struct {
	server        *grpc.Server
	reportService gen.ReportServiceServer
}

func NewServer(reportService gen.ReportServiceServer) *grpcServer {
	return &grpcServer{
		reportService: reportService,
	}
}
