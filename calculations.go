package main

import (
	"fmt"
	"time"
)

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
