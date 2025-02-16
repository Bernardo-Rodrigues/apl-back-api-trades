package grpc

import (
	"app/infra/adapters/grpc/services/report/gen"
)

func (s *grpcServer) addServices() {
	gen.RegisterReportServiceServer(s.server, s.reportService)
}
