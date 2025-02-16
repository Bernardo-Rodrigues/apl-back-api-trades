package dto

import (
	"app/core/domain"
	"app/core/domain/enum"
	"fmt"
	"strconv"
	"time"
)

type TradeDto struct {
	Date          time.Time
	AssetName     string
	AssetQuantity int
	AssetPrice    float64
	TradeType     string
}

type TradeDtos []TradeDto

func (tds TradeDtos) ToDomain() domain.Trades {
	var domainTrades []domain.Trade
	for _, t := range tds {
		domainTrade := domain.Trade{
			Date:          t.Date,
			AssetName:     t.AssetName,
			AssetQuantity: t.AssetQuantity,
			AssetPrice:    t.AssetPrice,
			TradeType:     enum.TradeType(t.TradeType),
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
		Date:          date,
		AssetName:     line[1],
		AssetQuantity: quantity,
		AssetPrice:    price,
		TradeType:     line[4],
	}

	return trade, nil
}
