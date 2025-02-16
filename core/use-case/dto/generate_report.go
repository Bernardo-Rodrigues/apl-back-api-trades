package dto

import (
	"app/core/domain"
	"time"
)

type GenerateReportDto struct {
	trades          domain.Trades
	prices          domain.PricesPerInstant
	startDate       time.Time
	endDate         time.Time
	intervalMinutes int
	initialBalance  float64
}

func New(trades TradeDtos, prices PricesDto, startDate, endDate time.Time, intervalMinutes int, initialBalance float64) *GenerateReportDto {
	return &GenerateReportDto{
		trades:          trades.ToDomain(),
		prices:          prices.ToDomain(),
		startDate:       startDate,
		endDate:         endDate,
		intervalMinutes: intervalMinutes,
		initialBalance:  initialBalance,
	}
}

func (g *GenerateReportDto) GetTrades() domain.Trades {
	return g.trades
}

func (g *GenerateReportDto) GetPrices() domain.PricesPerInstant {
	return g.prices
}

func (g *GenerateReportDto) GetStartDate() time.Time {
	return g.startDate
}

func (g *GenerateReportDto) GetEndDate() time.Time {
	return g.endDate
}

func (g *GenerateReportDto) GetMinutesInterval() int {
	return g.intervalMinutes
}

func (g *GenerateReportDto) GetInitialBalance() float64 {
	return g.initialBalance
}
