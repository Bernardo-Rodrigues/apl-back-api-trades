package dto

import "time"

type ReportDto []reportLineDto

type reportLineDto struct {
	date              time.Time
	totalBalance      float64
	accumulatedProfit float64
}

func NewReportLine(date time.Time, totalBalance float64, accumulatedProfit float64) *reportLineDto {
	return &reportLineDto{date, totalBalance, accumulatedProfit}
}

func (rl *reportLineDto) GetDate() time.Time {
	return rl.date
}

func (rl *reportLineDto) GetTotalBalance() float64 {
	return rl.totalBalance
}

func (rl *reportLineDto) GetAccumulatedProfit() float64 {
	return rl.accumulatedProfit
}
