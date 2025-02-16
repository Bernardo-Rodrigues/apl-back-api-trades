package generate_report_test

import (
	"app/core/domain/enum"
	dto2 "app/core/use-case/dto"
	"app/core/use-case/generate-report"
	"testing"
	"time"
)

func TestGenerateReport_Execute(t *testing.T) {
	trades := dto2.TradeDtos{
		{
			Date:          time.Date(2025, 02, 15, 10, 0, 0, 0, time.UTC),
			AssetName:     "BTC",
			AssetQuantity: 1,
			AssetPrice:    50000.0,
			TradeType:     string(enum.BUY),
		},
		{
			Date:          time.Date(2025, 02, 15, 11, 0, 0, 0, time.UTC),
			AssetName:     "BTC",
			AssetQuantity: 1,
			AssetPrice:    52000.0,
			TradeType:     string(enum.SELL),
		},
	}

	prices := dto2.PricesDto{
		time.Date(2025, 02, 15, 10, 0, 0, 0, time.UTC): {
			"BTC": 50000.0,
		},
		time.Date(2025, 02, 15, 11, 0, 0, 0, time.UTC): {
			"BTC": 52000.0,
		},
	}

	startDate := time.Date(2025, 02, 15, 9, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 02, 15, 12, 0, 0, 0, time.UTC)
	intervalMinutes := 60
	initialBalance := 100000.0

	generateReportDto := dto2.New(trades, prices, startDate, endDate, intervalMinutes, initialBalance)

	useCase := generate_report.New()
	useCase.Execute(*generateReportDto)
}
