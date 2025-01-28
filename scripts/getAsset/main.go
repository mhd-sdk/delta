package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/kr/pretty"
	"github.com/lmittmann/tint"
)

func main() {
	// get first parameter
	if len(os.Args) < 2 {
		slog.Error("Please provide a ticker")
		return
	}
	ticker := os.Args[1]
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			AddSource:  true,
		}),
	))
	apiKey := "PKXRB7R43VGCYC9CX34R"
	secretKey := "DHhpzJBYejT00v913aDpbJKViSWcI7QaJOAy1zp3"

	tradingClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	marketData := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   "https://data.alpaca.markets",
	})

	asset, err := tradingClient.GetAsset(ticker)
	if err != nil {
		slog.Error("Error getting asset", err.Error())
		return
	}
	pretty.Println(asset)
	startDate := time.Now().AddDate(0, 0, -20)
	bars, err := marketData.GetBars(ticker, marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneDay,
		Feed:      marketdata.IEX,
		Start:     startDate,
	})
	if err != nil {
		slog.Error("Error getting bars", err.Error())
		return
	}

	var (
		AvgVolume  int
		currentVol int
		AvgChange  float64
		DayChange  float64
		LastPrice  float64
	)
	for _, bar := range bars {
		AvgVolume += int(bar.Volume)
		if bar.Close > bar.Open {
			AvgChange += (bar.Close - bar.Open) / bar.Open
		} else {
			AvgChange += (bar.Open - bar.Close) / bar.Open
		}
	}

	AvgVolume = AvgVolume / len(bars)
	if AvgVolume == 0 {
		slog.Error("No volume data available")
		return
	}

	snapshot, err := marketData.GetSnapshot(ticker, marketdata.GetSnapshotRequest{
		Feed: marketdata.IEX,
	})
	if err != nil {
		slog.Error("Error getting bars", err.Error())
	}

	AvgChange = AvgChange / float64(len(bars))
	currentVol = int(snapshot.DailyBar.Volume)
	LastPrice = snapshot.DailyBar.Close
	DayChange = (snapshot.DailyBar.Close - snapshot.DailyBar.Open) / snapshot.DailyBar.Open
	RelativeVolume := currentVol / AvgVolume

	slog.Info("Asset stats", "ticker", ticker, "name", asset.Name, "AvgVolume", AvgVolume, "currentVol", currentVol, "AvgChange", AvgChange, "DayChange", DayChange, "LastPrice", LastPrice, "RelativeVolume", RelativeVolume)

}
