package files

import (
	"app/core/use-case/dto"
	"app/infra/adapters/files/csv"
	"errors"
	"io"
	"time"
)

type FilesHandler interface {
	LoadValuesInInterval(startDate, endDate time.Time) (dto.TradeDtos, dto.PricesDto, error)
	BuildByteArray(report dto.ReportDto) ([]byte, error)
}

func NewHandler(fileType string, tradesFile io.Reader, assetsFiles map[string]io.Reader) (FilesHandler, error) {
	if fileType == "csv" {
		return csv.NewHandler(tradesFile, assetsFiles), nil
	}
	return nil, errors.New("unsupported file type: " + fileType)
}
