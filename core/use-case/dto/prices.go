package dto

import (
	"app/core/domain"
	"time"
)

type PricesDto map[time.Time]map[string]float64

func (pd PricesDto) ToDomain() domain.PricesPerInstant {
	domainPrices := domain.PricesPerInstant{}

	for date, prices := range pd {
		assetsPrice := domain.AssetsPrice{}
		for assetName, price := range prices {
			assetsPrice[assetName] = price
		}
		domainPrices[date] = assetsPrice
	}

	return domainPrices
}
