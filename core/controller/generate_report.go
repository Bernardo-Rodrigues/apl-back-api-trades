package controller

import (
	"app/core/use-case/dto"
	"time"
)

func (c *reportController) GenerateReport(startDate, endDate time.Time, intervalMinutes int, initialBalance float64) ([]byte, error) {
	trades, prices, err := c.filesHandler.LoadValuesInInterval(startDate, endDate)
	if err != nil {
		return nil, err
	}
	generateReportDto := dto.New(trades, prices, startDate, endDate, intervalMinutes, initialBalance)

	report := c.usecase.Execute(*generateReportDto)

	return c.filesHandler.BuildByteArray(report)
}
