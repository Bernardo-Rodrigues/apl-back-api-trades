package controller

import (
	"app/core/use-case/dto"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
)

func loadValues(startDate, endDate time.Time, tradesFile io.Reader, assetsFiles map[string]io.Reader) (dto.TradeDtos, dto.PricesDto) {
	chTrades := make(chan dto.TradeDto, 100)
	chPrices := make(chan dto.PricesDto, len(assetsFiles))
	var wg sync.WaitGroup

	wg.Add(1)
	go loadTrades(tradesFile, startDate, endDate, chTrades, &wg)

	for assetName, assetFile := range assetsFiles {
		wg.Add(1)
		go loadPrices(assetFile, assetName, startDate, endDate, chPrices, &wg)
	}

	go func() {
		wg.Wait()
		close(chTrades)
		close(chPrices)
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

	return trades, prices
}

func loadTrades(file io.Reader, start, end time.Time, ch chan<- dto.TradeDto, wg *sync.WaitGroup) {
	defer wg.Done()

	reader := csv.NewReader(file)
	layout := "2006-01-02 15:04:05"

	_, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading trades file headers:", err)
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading trades file line:", err)
			break
		}

		date, err := time.Parse(layout, line[0])
		if err != nil {
			fmt.Println(fmt.Errorf("error parsing date: %v", err))
			continue
		}
		if date.Before(start) {
			continue
		}
		if date.After(end) {
			break
		}

		tradeDto, err := dto.NewTradeDtoFromCSV(line, date)

		ch <- tradeDto
	}
}

func loadPrices(file io.Reader, asset string, start, end time.Time, ch chan<- dto.PricesDto, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := csv.NewReader(file)

	prices := make(dto.PricesDto)
	layout := "2006-01-02 15:04:05"

	_, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading prices file headers:", err)
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading prices file line:", err)
			break
		}

		date, err := time.Parse(layout, line[0])
		if err != nil {
			fmt.Println(fmt.Errorf("error parsing date: %v", err))
			continue
		}

		if date.Before(start) {
			continue
		}
		if date.After(end) {
			break
		}

		price, _ := strconv.ParseFloat(line[1], 64)

		if _, ok := prices[date]; !ok {
			prices[date] = make(map[string]float64)
		}
		prices[date][asset] = price
	}

	ch <- prices
}
