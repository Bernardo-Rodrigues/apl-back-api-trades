package domain

import (
	"fmt"
	"time"
)

type PricesPerInstant map[time.Time]AssetsPrice

func (prices PricesPerInstant) CalculateAssetsValueAtIntervalEnd(end time.Time, assetsQuantity AssetsQuantity) float64 {
	totalValue := 0.0
	for name, quantity := range assetsQuantity {
		assetValue := prices.GetInstantPrice(name, end)
		totalValue += float64(quantity) * assetValue
		fmt.Printf("[%s] Asset %s with %d units valued at $%.2f each, total $%.2f\n", end.Format("2006-01-02 15:04:05"), name, quantity, assetValue, float64(quantity)*assetValue)
	}

	return totalValue
}

func (prices PricesPerInstant) GetInstantPrice(assetName string, instant time.Time) float64 {
	truncateInstant := time.Date(instant.Year(), instant.Month(), instant.Day(), instant.Hour(), instant.Minute(), 0, 0, instant.Location())

	assetsPrice := prices[truncateInstant]
	return assetsPrice[assetName]
}
