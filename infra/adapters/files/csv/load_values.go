package csv

import (
	"app/core/use-case/dto"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
)

func (adp cvsHandler) LoadValuesInInterval(startDate, endDate time.Time) (dto.TradeDtos, dto.PricesDto, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chTrades := make(chan dto.TradeDto, 100)
	chPrices := make(chan dto.PricesDto, len(adp.assetsFiles))
	chErr := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := loadTrades(ctx, adp.tradesFile, startDate, endDate, chTrades); err != nil {
			select {
			case chErr <- err:
				cancel()
			default:
			}
		}
	}()

	for assetName, assetFile := range adp.assetsFiles {
		wg.Add(1)
		go func(asset string, file io.Reader) {
			defer wg.Done()
			if err := loadPrices(ctx, file, asset, startDate, endDate, chPrices); err != nil {
				select {
				case chErr <- err:
					cancel()
				default:
				}
			}
		}(assetName, assetFile)
	}

	go func() {
		wg.Wait()
		close(chTrades)
		close(chPrices)
		close(chErr)
	}()

	var trades dto.TradeDtos
	for trade := range chTrades {
		trades = append(trades, trade)
	}

	prices := make(map[time.Time]map[string]float64)
	for assetsPriceAtInstant := range chPrices {
		for instant, assetsPrice := range assetsPriceAtInstant {
			if _, ok := prices[instant]; !ok {
				prices[instant] = make(map[string]float64)
			}
			for assets, price := range assetsPrice {
				prices[instant][assets] = price
			}
		}
	}

	if err, ok := <-chErr; ok {
		return nil, nil, err
	}

	return trades, prices, nil
}
func loadTrades(ctx context.Context, file io.Reader, start, end time.Time, ch chan<- dto.TradeDto) error {
	defer close(ch)

	reader := csv.NewReader(file)
	layout := "2006-01-02 15:04:05"

	_, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading trades file headers: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, err := reader.Read()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("error reading trades file line: %v", err)
			}

			date, err := time.Parse(layout, line[0])
			if err != nil {
				return fmt.Errorf("error parsing date: %v", err)
			}
			if date.Before(start) {
				continue
			}
			if date.After(end) {
				return nil
			}

			tradeDto, err := dto.NewTradeDtoFromCSV(line, date)
			if err != nil {
				return err
			}

			select {
			case ch <- tradeDto:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func loadPrices(ctx context.Context, file io.Reader, asset string, start, end time.Time, ch chan<- dto.PricesDto) error {
	defer close(ch)
	reader := csv.NewReader(file)

	prices := make(dto.PricesDto)
	layout := "2006-01-02 15:04:05"

	_, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading prices file headers: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, err := reader.Read()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("error reading prices file line: %v", err)
			}

			date, err := time.Parse(layout, line[0])
			if err != nil {
				return fmt.Errorf("error parsing date: %v", err)
			}

			if date.Before(start) {
				continue
			}
			if date.After(end) {
				return nil
			}

			price, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				return fmt.Errorf("error parsing price: %v", err)
			}

			if _, ok := prices[date]; !ok {
				prices[date] = make(map[string]float64)
			}
			prices[date][asset] = price

			select {
			case ch <- prices:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}
