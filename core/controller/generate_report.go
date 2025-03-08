package controller

import (
	"app/core/use-case/dto"
	"fmt"
	"time"
)

func (c *reportController) GenerateReport(startDate, endDate time.Time, intervalMinutes int, initialBalance float64) ([]byte, error) {
	trades, prices := c.filesHandler.LoadValuesInInterval(startDate, endDate)
	generateReportDto := dto.New(trades, prices, startDate, endDate, intervalMinutes, initialBalance)

	report, err := c.usecase.Execute(*generateReportDto)
	if err != nil {
		return nil, fmt.Errorf("error executing report use case: %w", err)
	}

	return c.filesHandler.BuildByteArray(report)
}
