package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
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

func loadValues(startDate, endDate time.Time, tradesFile io.Reader, assetsFiles map[string]io.Reader) ([]Trade, map[time.Time]map[string]float64) {
	trades := loadTrades(tradesFile, startDate, endDate)

	prices := make(map[time.Time]map[string]float64)
	for assetName, assetFile := range assetsFiles {
		loadPrices(prices, assetFile, assetName, startDate, endDate)
	}

	return trades, prices
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

func calculateCashBalancePerInterval(trades []Trade, assets map[string]int) float64 {
	cash := 0.0

	for _, t := range trades {
		if t.TradesType == "BUY" {
			assets[t.AssetName] += t.AssetQuantity
			cash -= float64(t.AssetQuantity) * t.AssetPrice
			fmt.Printf("[%s] Bought %d units of %s, spent $%.2f\n", t.Date.Format("2006-01-02 15:04:05"), t.AssetQuantity, t.AssetName, float64(t.AssetQuantity)*t.AssetPrice)
		} else if t.TradesType == "SELL" {
			assets[t.AssetName] -= t.AssetQuantity
			cash += float64(t.AssetQuantity) * t.AssetPrice
			fmt.Printf("[%s] Sold %d units of %s, recovered $%.2f\n", t.Date.Format("2006-01-02 15:04:05"), t.AssetQuantity, t.AssetName, float64(t.AssetQuantity)*t.AssetPrice)
		}
	}

	return cash
}

func calculateAssetsValueAtIntervalEnd(prices map[time.Time]map[string]float64, end time.Time, assets map[string]int) float64 {
	assetsValue := 0.0
	for asset, quantity := range assets {
		assetValue := getInstantPrice(prices, asset, end)
		assetsValue += float64(quantity) * assetValue
		fmt.Printf("[%s] Asset %s with %d units valued at $%.2f each, total $%.2f\n", end.Format("2006-01-02 15:04:05"), asset, quantity, assetValue, float64(quantity)*assetValue)
	}

	return assetsValue
}

func getInstantPrice(prices map[time.Time]map[string]float64, asset string, instante time.Time) float64 {
	truncateInstant := time.Date(instante.Year(), instante.Month(), instante.Day(), instante.Hour(), instante.Minute(), 0, 0, instante.Location())

	assets := prices[truncateInstant]
	return assets[asset]
}

func loadTrades(file io.Reader, start, end time.Time) []Trade {
	var trades []Trade

	reader := csv.NewReader(file)
	layout := "2006-01-02 15:04:05"

	_, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading trades file headers:", err)
		return nil
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading trades file line:", err)
			break
		}

		date, _ := time.Parse(layout, line[0])
		if date.Before(start) {
			continue
		}
		if date.After(end) {
			break
		}

		quantity, _ := strconv.Atoi(line[2])
		price, _ := strconv.ParseFloat(line[3], 64)

		trade := Trade{
			Date:          date,
			AssetName:     line[1],
			AssetQuantity: quantity,
			AssetPrice:    price,
			TradesType:    line[4],
		}

		trades = append(trades, trade)
	}
	return trades
}

func loadPrices(prices map[time.Time]map[string]float64, file io.Reader, asset string, start, end time.Time) {
	reader := csv.NewReader(file)

	layout := "2006-01-02 15:04:05"

	_, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading prices file headers:", err)
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading prices file line:", err)
			break
		}

		date, _ := time.Parse(layout, line[0])

		if date.Before(start) {
			continue
		}
		if date.After(end) {
			break
		}

		price, _ := strconv.ParseFloat(line[1], 64)

		if _, ok := prices[date]; !ok {
			prices[date] = make(map[string]float64)
		}
		prices[date][asset] = price
	}
}

type Trade struct {
	Date          time.Time
	AssetName     string
	AssetQuantity int
	AssetPrice    float64
	TradesType    string
}

type Trades []Trade

func (trades Trades) FilterInInterval(start, end time.Time) []Trade {
	filteredTrades := make([]Trade, 0)

	for _, t := range trades {
		if t.Date.Before(start) {
			continue
		}
		if t.Date.After(end) {
			break
		}
		filteredTrades = append(filteredTrades, t)
	}

	return filteredTrades
}
