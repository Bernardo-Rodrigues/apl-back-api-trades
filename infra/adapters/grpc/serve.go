package grpc

import (
	"app/infra/adapters/grpc/middleware"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func (s *grpcServer) Serve() {
	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuditInterceptor()),
	)

	s.addServices()

	reflection.Register(s.server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening: %v\n", err)
	}

	fmt.Println("gRPC Server running at 50051...")
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
