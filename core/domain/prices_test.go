package domain_test

import (
	domain2 "app/core/domain"
	"testing"
	"time"
)

func TestCalculateAssetsValueAtIntervalEnd(t *testing.T) {
	prices := domain2.PricesPerInstant{
		time.Date(2025, 2, 16, 10, 0, 0, 0, time.UTC): domain2.AssetsPrice{
			"BTC": 50000.0,
			"ETH": 3000.0,
		},
		time.Date(2025, 2, 16, 11, 0, 0, 0, time.UTC): domain2.AssetsPrice{
			"BTC": 51000.0,
			"ETH": 3100.0,
		},
	}

	assetsQuantity := domain2.AssetsQuantity{
		"BTC": 2,
		"ETH": 5,
	}

	end := time.Date(2025, 2, 16, 11, 0, 0, 0, time.UTC)

	totalValue := prices.CalculateAssetsValueAtIntervalEnd(end, assetsQuantity)

	expectedValue := (2 * 51000.0) + (5 * 3100.0)
	if totalValue != expectedValue {
		t.Errorf("Expected %f but got %f", expectedValue, totalValue)
	}
}

func TestGetInstantPriceTruncatingSeconds(t *testing.T) {
	prices := domain2.PricesPerInstant{
		time.Date(2025, 2, 16, 10, 15, 0, 0, time.UTC): domain2.AssetsPrice{
			"BTC": 50000.0,
		},
	}

	price := prices.GetInstantPrice("BTC", time.Date(2025, 2, 16, 10, 15, 10, 0, time.UTC))

	expectedPrice := 50000.0
	if price != expectedPrice {
		t.Errorf("Expected %f but got %f", expectedPrice, price)
	}
}

func TestGetInstantPrice_NoPriceFound(t *testing.T) {
	prices := domain2.PricesPerInstant{
		time.Date(2025, 2, 16, 10, 0, 0, 0, time.UTC): domain2.AssetsPrice{
			"BTC": 50000.0,
		},
	}

	price := prices.GetInstantPrice("ETH", time.Date(2025, 2, 16, 10, 15, 0, 0, time.UTC))

	if price != 0.0 {
		t.Errorf("Expected 0.0 but got %f", price)
	}
}
