package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	start := time.Now()
	layout := "2006-01-02 15:04:05"
	startDate, _ := time.Parse(layout, "2021-03-01 10:00:00")
	endDate, _ := time.Parse(layout, "2021-03-01 10:16:00")

	tradesFile, err := os.Open("./arquivos-exemplo/march_2021_trades.csv")
	if err != nil {
		fmt.Println("Error loading trades:", err)
		return
	}
	defer tradesFile.Close()

	assetsFiles := make(map[string]io.Reader)

	assetAFile, err := os.Open("./arquivos-exemplo/march_2021_pricesA.csv")
	if err != nil {
		fmt.Println("Error loading asset A file:", err)
		return
	}
	defer assetAFile.Close()
	assetsFiles["A"] = assetAFile

	assetBFile, err := os.Open("./arquivos-exemplo/march_2021_pricesB.csv")
	if err != nil {
		fmt.Println("Error loading asset B file:", err)
		return
	}
	defer assetBFile.Close()
	assetsFiles["B"] = assetBFile

	trades, prices := loadValues(startDate, endDate, tradesFile, assetsFiles)
	generateReport(trades, prices, startDate, endDate, 10, 100_000)
	fmt.Printf("Execution time: %v µs\n", time.Since(start).Microseconds())
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
