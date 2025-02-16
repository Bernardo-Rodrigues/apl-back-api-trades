package generate_report

import (
	"app/core/use-case/dto"
	"fmt"
	"time"
)

func (u generateReportUsecase) Execute(dto dto.GenerateReportDto) {
	cashBalance := dto.GetInitialBalance()
	fmt.Printf("%s\t%.4f\t%.5f\n", dto.GetStartDate().Format("2006-01-02 15:04:05"), cashBalance, 0.0)

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

		fmt.Printf("%s\t%.4f\t%.5f\n", intervalEnd.Format("2006-01-02 15:04:05"), totalBalance, accumulatedProfit)
	}
}
