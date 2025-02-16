package middleware

import (
	"app/infra/adapters/grpc/services/report/gen"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var logFilePath = getLogFilePath()

func getLogFilePath() string {
	path := os.Getenv("LOG_PATH")
	if path == "" {
		path = "./logs/audit.log"
	}
	return path
}

func AuditInterceptor() grpc.UnaryServerInterceptor {
	logDir := logFilePath[:len(logFilePath)-len("/audit.log")]
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating log directory: %v", err)
	}

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	log.SetOutput(logFile)

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		p, ok := peer.FromContext(ctx)
		clientIP, clientPort := "unknown", "unknown"
		if ok {
			clientIP, clientPort = parseIPPort(p.Addr.String())
		}

		resp, err := handler(ctx, req)
		status := "Success"
		errMessage := ""
		fileSize := 0
		fileBase64 := "No file response"

		if err != nil {
			status = "Error"
			errMessage = err.Error()
		} else {
			if reportResp, ok := resp.(*gen.ReportResponse); ok {
				fileSize = len(reportResp.File)
				fileBase64 = base64.StdEncoding.EncodeToString(reportResp.File)
			}
		}

		auditLog := map[string]interface{}{
			"timestamp":     time.Now(),
			"client_ip":     clientIP,
			"client_port":   clientPort,
			"method":        info.FullMethod,
			"request":       req,
			"status":        status,
			"error_message": errMessage,
			"file_size":     fileSize,
			"file_base64":   fileBase64,
		}

		logData, _ := json.Marshal(auditLog)
		log.Println(string(logData))

		return resp, err
	}
}

func parseIPPort(addr string) (string, string) {
	for i := len(addr) - 1; i >= 0; i-- {
		if addr[i] == ':' {
			return addr[:i], addr[i+1:]
		}
	}
	return addr, "unknown"
}
