package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
)

func loadValues(startDate, endDate time.Time, tradesFile io.Reader, assetsFiles map[string]io.Reader) ([]Trade, map[time.Time]map[string]float64) {
	chTrades := make(chan Trade, 100)
	chPrices := make(chan map[time.Time]map[string]float64, len(assetsFiles))
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

	trades := []Trade{}
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

func loadTrades(file io.Reader, start, end time.Time, ch chan<- Trade, wg *sync.WaitGroup) {
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

		date, _ := time.Parse(layout, line[0])
		if date.Before(start) {
			continue
		}
		if date.After(end) {
			break
		}

		quantity, _ := strconv.Atoi(line[2])
		price, _ := strconv.ParseFloat(line[3], 64)

		trade := Trade{
			Date:          date,
			AssetName:     line[1],
			AssetQuantity: quantity,
			AssetPrice:    price,
			TradesType:    line[4],
		}

		ch <- trade
	}
}

func loadPrices(file io.Reader, asset string, start, end time.Time, ch chan<- map[time.Time]map[string]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := csv.NewReader(file)

	prices := make(map[time.Time]map[string]float64)
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

		date, _ := time.Parse(layout, line[0])

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
