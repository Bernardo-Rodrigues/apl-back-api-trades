package generate_report

import (
	"app/core/use-case/dto"
	"time"
)

func (u generateReportUsecase) Execute(input dto.GenerateReportDto) (dto.ReportDto, error) {
	report := make(dto.ReportDto, 0, len(input.GetTrades()))

	intervalDuration := time.Duration(input.GetMinutesInterval()) * time.Minute
	assets := make(map[string]int)
	cashBalance := input.GetInitialBalance()

	report = append(report, *dto.NewReportLine(input.GetStartDate(), cashBalance, 0.0))

	for intervalStart := input.GetStartDate(); intervalStart.Before(input.GetEndDate()); intervalStart = intervalStart.Add(intervalDuration) {
		intervalEnd := intervalStart.Add(intervalDuration)
		if intervalEnd.After(input.GetEndDate()) {
			intervalEnd = input.GetEndDate()
		}

		intervalTrades := input.GetTrades().FilterInInterval(intervalStart, intervalEnd, input.GetEndDate())
		cashInterval := intervalTrades.CalculateCashBalancePerInterval(assets)
		assetsValue := input.GetPrices().CalculateAssetsValueAtIntervalEnd(intervalEnd, assets)

		cashBalance += cashInterval
		totalBalance := cashBalance + assetsValue
		accumulatedProfit := (totalBalance / input.GetInitialBalance()) - 1

		report = append(report, *dto.NewReportLine(intervalEnd, totalBalance, accumulatedProfit))
	}

	return report, nil
}
