package domain

import (
	"app/core/domain/enum"
	"fmt"
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

func (trades Trades) FilterInInterval(start, end time.Time) Trades {
	filteredTrades := make(Trades, 0)

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

func (trades Trades) CalculateCashBalancePerInterval(assetsQuantity AssetsQuantity) float64 {
	cash := 0.0

	for _, t := range trades {
		if t.TradeType == enum.BUY {
			assetsQuantity[t.AssetName] += t.AssetQuantity
			cash -= float64(t.AssetQuantity) * t.AssetPrice
			fmt.Printf("[%s] Bought %d units of %s, spent $%.2f\n", t.Date.Format("2006-01-02 15:04:05"), t.AssetQuantity, t.AssetName, float64(t.AssetQuantity)*t.AssetPrice)
		} else if t.TradeType == enum.SELL {
			assetsQuantity[t.AssetName] -= t.AssetQuantity
			cash += float64(t.AssetQuantity) * t.AssetPrice
			fmt.Printf("[%s] Sold %d units of %s, recovered $%.2f\n", t.Date.Format("2006-01-02 15:04:05"), t.AssetQuantity, t.AssetName, float64(t.AssetQuantity)*t.AssetPrice)
		}
	}

	return cash
}
