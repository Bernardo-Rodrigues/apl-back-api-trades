package controller

import (
	"app/core/use-case/dto"
	"io"
	"time"
)

func (c reportController) GenerateReport(startDate, endDate time.Time, intervalMinutes int, initialBalance float64, tradesFile io.Reader, assetsFiles map[string]io.Reader) (string, error) {
	trades, prices := loadValues(startDate, endDate, tradesFile, assetsFiles)
	generateReportDto := dto.New(trades, prices, startDate, endDate, intervalMinutes, initialBalance)

	return c.usecase.Execute(*generateReportDto)
}
