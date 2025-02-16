package main

import "time"

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
