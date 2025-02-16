package generate_report

import (
	"app/core/use-case/dto"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func (u generateReportUsecase) Execute(dto dto.GenerateReportDto) (string, error) {
	tempDir := os.TempDir()

	tempFile, err := os.CreateTemp(tempDir, "report-*.csv")
	if err != nil {
		return "", fmt.Errorf("error creating temporary file: %w", err)
	}
	defer tempFile.Close()

	writer := csv.NewWriter(tempFile)
	defer writer.Flush()

	writer.Write([]string{"timestamp", "Patrimônio Total", "Rentabilidade Acumulada"})

	cashBalance := dto.GetInitialBalance()
	intervalDuration := time.Duration(dto.GetMinutesInterval()) * time.Minute
	assets := make(map[string]int)

	for intervalStart := dto.GetStartDate(); intervalStart.Before(dto.GetEndDate()); intervalStart = intervalStart.Add(intervalDuration) {
		intervalEnd := intervalStart.Add(intervalDuration)
		if intervalEnd.After(dto.GetEndDate()) {
			intervalEnd = dto.GetEndDate()
		}

		intervalTrades := dto.GetTrades().FilterInInterval(intervalStart, intervalEnd)
		cashInterval := intervalTrades.CalculateCashBalancePerInterval(assets)
		assetsValue := dto.GetPrices().CalculateAssetsValueAtIntervalEnd(intervalEnd, assets)

		cashBalance += cashInterval
		totalBalance := cashBalance + assetsValue
		accumulatedProfit := (totalBalance / dto.GetInitialBalance()) - 1

		writer.Write([]string{
			intervalEnd.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%.4f", totalBalance),
			fmt.Sprintf("%.5f", accumulatedProfit),
		})
	}

	return tempFile.Name(), nil
}
