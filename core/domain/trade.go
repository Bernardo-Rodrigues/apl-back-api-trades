package domain

import (
	"app/core/domain/enum"
	"time"
)

type Trade struct {
	Date          time.Time
	AssetName     string
	AssetQuantity int
	AssetPrice    float64
	TradeType     enum.TradeType
}

type Trades []Trade

func (trades Trades) FilterInInterval(intervalStart, intervalEnd, endDate time.Time) Trades {
	filteredTrades := make(Trades, 0)
	
	for _, t := range trades {
		if t.Date.Before(intervalStart) {
			continue
		}

		if t.Date.Equal(intervalEnd) && !intervalEnd.Equal(endDate) {
			continue
		}

		if t.Date.After(intervalEnd) {
			break
		}

		filteredTrades = append(filteredTrades, t)
	}

	return filteredTrades
}

func (trades Trades) CalculateCashBalancePerInterval(assetsQuantity AssetsQuantity) float64 {
	cash := 0.0

	for _, t := range trades {
		if t.TradeType == enum.BUY {
			assetsQuantity[t.AssetName] += t.AssetQuantity
			cash -= float64(t.AssetQuantity) * t.AssetPrice
		} else if t.TradeType == enum.SELL {
			assetsQuantity[t.AssetName] -= t.AssetQuantity
			cash += float64(t.AssetQuantity) * t.AssetPrice
		}
	}

	return cash
}
