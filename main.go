package main

import (
	"fmt"
	"time"
)

func main() {
	StartGrpcServer()
}

func generateReport(trades Trades, prices map[time.Time]map[string]float64, startDate, endDate time.Time, minutesInterval int, initialBalance float64) {
	cashBalance := initialBalance
	fmt.Printf("%s\t%.4f\t%.5f\n", startDate.Format("2006-01-02 15:04:05"), cashBalance, 0.0)

	intervalDuration := time.Duration(minutesInterval) * time.Minute
	assets := make(map[string]int)

	for intervalStart := startDate; intervalStart.Before(endDate); intervalStart = intervalStart.Add(intervalDuration) {
		intervalEnd := intervalStart.Add(intervalDuration)
		if intervalEnd.After(endDate) {
			intervalEnd = endDate
		}

		intervalTrades := trades.FilterInInterval(intervalStart, intervalEnd)

		cashInterval := calculateCashBalancePerInterval(intervalTrades, assets)
		assetsValue := calculateAssetsValueAtIntervalEnd(prices, intervalEnd, assets)

		cashBalance += cashInterval
		totalBalance := cashBalance + assetsValue

		accumulatedProfit := (totalBalance / initialBalance) - 1

		fmt.Printf("%s\t%.4f\t%.5f\n", intervalEnd.Format("2006-01-02 15:04:05"), totalBalance, accumulatedProfit)
	}
}
