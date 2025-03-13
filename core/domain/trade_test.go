package domain_test

import (
	domain2 "app/core/domain"
	"app/core/domain/enum"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilterInInterval(t *testing.T) {
	trades := domain2.Trades{
		{Date: time.Date(2025, 2, 15, 10, 0, 0, 0, time.UTC), AssetName: "BTC", AssetQuantity: 2, AssetPrice: 30000, TradeType: enum.BUY},
		{Date: time.Date(2025, 2, 16, 10, 0, 0, 0, time.UTC), AssetName: "BTC", AssetQuantity: 1, AssetPrice: 35000, TradeType: enum.SELL},
		{Date: time.Date(2025, 2, 17, 10, 0, 0, 0, time.UTC), AssetName: "ETH", AssetQuantity: 5, AssetPrice: 2000, TradeType: enum.BUY},
	}

	start := time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 2, 16, 23, 59, 59, 0, time.UTC)
	final := time.Date(2025, 2, 18, 23, 59, 59, 0, time.UTC)

	filteredTrades := trades.FilterInInterval(start, end, final)

	assert.Len(t, filteredTrades, 2, "Expected 2 trades in the interval")
	assert.Equal(t, "BTC", filteredTrades[0].AssetName, "First trade should be BTC")
	assert.Equal(t, "BTC", filteredTrades[1].AssetName, "Second trade should be BTC")
}

func TestCalculateCashBalancePerInterval(t *testing.T) {
	trades := domain2.Trades{
		{Date: time.Date(2025, 2, 15, 10, 0, 0, 0, time.UTC), AssetName: "BTC", AssetQuantity: 2, AssetPrice: 30000, TradeType: enum.BUY},
		{Date: time.Date(2025, 2, 16, 10, 0, 0, 0, time.UTC), AssetName: "BTC", AssetQuantity: 1, AssetPrice: 35000, TradeType: enum.SELL},
	}

	assetsQuantity := domain2.AssetsQuantity{
		"BTC": 0,
	}

	cash := trades.CalculateCashBalancePerInterval(assetsQuantity)

	assert.Equal(t, -25000.0, cash, "Expected cash balance to be -25000.0 after the trades")
	assert.Equal(t, 1, assetsQuantity["BTC"], "Expected 1 BTC remaining after trades")
}
