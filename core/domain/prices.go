package domain

import (
	"time"
)

type PricesPerInstant map[time.Time]AssetsPrice

func (prices PricesPerInstant) CalculateAssetsValueAtIntervalEnd(end time.Time, assetsQuantity AssetsQuantity) float64 {
	totalValue := 0.0
	for name, quantity := range assetsQuantity {
		assetValue := prices.GetInstantPrice(name, end)
		totalValue += float64(quantity) * assetValue
	}

	return totalValue
}

func (prices PricesPerInstant) GetInstantPrice(assetName string, instant time.Time) float64 {
	truncateInstant := time.Date(instant.Year(), instant.Month(), instant.Day(), instant.Hour(), instant.Minute(), 0, 0, instant.Location())

	assetsPrice := prices[truncateInstant]
	return assetsPrice[assetName]
}
