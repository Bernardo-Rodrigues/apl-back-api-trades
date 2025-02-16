package dto

import (
	"app/core/domain"
	"app/core/domain/enum"
	"fmt"
	"strconv"
	"time"
)

type TradeDto struct {
	date          time.Time
	assetName     string
	assetQuantity int
	assetPrice    float64
	tradeType     string
}

type TradeDtos []TradeDto

func (tds TradeDtos) ToDomain() domain.Trades {
	var domainTrades []domain.Trade
	for _, t := range tds {
		domainTrade := domain.Trade{
			Date:          t.date,
			AssetName:     t.assetName,
			AssetQuantity: t.assetQuantity,
			AssetPrice:    t.assetPrice,
			TradeType:     enum.TradeType(t.tradeType),
		}
		domainTrades = append(domainTrades, domainTrade)
	}
	return domainTrades
}

func NewTradeDtoFromCSV(line []string, date time.Time) (TradeDto, error) {
	quantity, err := strconv.Atoi(line[2])
	if err != nil {
		return TradeDto{}, fmt.Errorf("error converting quantity: %v", err)
	}

	price, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		return TradeDto{}, fmt.Errorf("error converting price: %v", err)
	}

	trade := TradeDto{
		date:          date,
		assetName:     line[1],
		assetQuantity: quantity,
		assetPrice:    price,
		tradeType:     line[4],
	}

	return trade, nil
}
