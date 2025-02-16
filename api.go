package main

import (
	"app/report/gen"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"time"
)

func StartGrpcServer() {
	server := grpc.NewServer()

	gen.RegisterReportServiceServer(server, &reportService{})

	reflection.Register(server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening: %v\n", err)
	}

	fmt.Println("gRPC Server running at 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

type reportService struct {
	gen.UnimplementedReportServiceServer
}

func (s *reportService) GenerateReport(ctx context.Context, req *gen.ReportRequest) (*gen.ReportResponse, error) {
	start := time.Now()
	layout := "2006-01-02 15:04:05"

	startDate, _ := time.Parse(layout, req.StartDate)
	endDate, _ := time.Parse(layout, req.EndDate)

	tradesFile := bytes.NewReader(req.TradesFile)
	assetsFiles := make(map[string]io.Reader)
	for fileName, content := range req.AssetsFiles {
		assetsFiles[fileName] = bytes.NewReader(content)
	}

	trades, prices := loadValues(startDate, endDate, tradesFile, assetsFiles)
	generateReport(trades, prices, startDate, endDate, int(req.IntervalMinutes), float64(req.InitialBalance))

	fmt.Printf("Execution time (microseconds): %d µs\n", time.Since(start).Microseconds())
	return &gen.ReportResponse{
		Message: "Report generated successfully!",
	}, nil
}
