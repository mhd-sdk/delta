package main

import (
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/kr/pretty"
	"github.com/lmittmann/tint"
)

type ScanResult struct {
	Ticker         string
	Name           string
	AvgVolume      int
	CurrentVolume  int
	RelativeVolume int
	AvgChange      float64
	DayChange      float64
	LastPrice      float64
}

type ScanResults []ScanResult

type Event struct {
	Type string
	Time time.Time
	ScanResult
}

type Events []Event

type WatcherConfig struct {
	PollInterval time.Duration
	ApiKey       string
	SecretKey    string
}

type Watcher struct {
	config        WatcherConfig
	events        chan Event
	storedEvents  *Events
	dataClient    *marketdata.Client
	tradingClient *alpaca.Client
}

func (evts *Events) ContainsTicker(ticker string) bool {
	for _, event := range *evts {
		if event.Ticker == ticker {
			return true
		}
	}
	return false
}

func NewWatcher(config WatcherConfig) *Watcher {
	marketData := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    config.ApiKey,
		APISecret: config.SecretKey,
		BaseURL:   "https://data.alpaca.markets",
	})

	tradingClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    config.ApiKey,
		APISecret: config.SecretKey,
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	return &Watcher{
		config:        config,
		events:        make(chan Event, 100),
		storedEvents:  &Events{},
		dataClient:    marketData,
		tradingClient: tradingClient,
	}
}

func (w *Watcher) Start() {
	assets, err := w.tradingClient.GetAssets(alpaca.GetAssetsRequest{
		AssetClass: string(alpaca.USEquity),
		Exchange:   "NASDAQ",
		Status:     "active",
	})
	if err != nil {
		slog.Error("Error getting asset", err.Error())
		return
	}

	for {
		assetNames := make([]string, 0)
		for _, asset := range assets {
			assetNames = append(assetNames, asset.Symbol)
		}
		slog.Info("Getting bars for assets", "assets count", len(assetNames))
		startDate := time.Now().AddDate(0, 0, -20)
		barsData, err := w.dataClient.GetMultiBars(assetNames, marketdata.GetBarsRequest{
			TimeFrame: marketdata.OneDay,
			Feed:      marketdata.IEX,
			Start:     startDate,
		})
		if err != nil {
			slog.Error("Error getting bars", err.Error())
		}

		snapshots, err := w.dataClient.GetSnapshots(assetNames, marketdata.GetSnapshotRequest{
			Feed: marketdata.IEX,
		})
		if err != nil {
			slog.Error("Error getting bars", err.Error())
		}

		slog.Info("Data received", "data count", len(barsData))

		var wg sync.WaitGroup
		var mu sync.Mutex
		for ticker, bars := range barsData {
			if len(bars) == 0 {
				slog.Warn("No bars for asset", "ticker", ticker)
				continue
			}

			// find name of the asset
			var name string
			for _, asset := range assets {
				if asset.Symbol == ticker {
					name = asset.Name
					break
				}
			}

			wg.Add(1)
			go computeStats(&wg, &mu, w.events, ticker, name, bars, snapshots[ticker], w.storedEvents)
		}
		wg.Wait()
	}
}

func computeStats(wg *sync.WaitGroup, mu *sync.Mutex, eventChan chan Event, ticker string, name string, bars []marketdata.Bar, snapshot *marketdata.Snapshot, storedEvents *Events) {
	// if daily bar is not today
	if snapshot.DailyBar.Timestamp.Day() != time.Now().Day() {
		wg.Done()
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
		wg.Done()
		return
	}
	AvgChange = AvgChange / float64(len(bars))
	currentVol = int(snapshot.DailyBar.Volume)
	LastPrice = snapshot.DailyBar.Close
	DayChange = (snapshot.DailyBar.Close - snapshot.DailyBar.Open) / snapshot.DailyBar.Open
	RelativeVolume := currentVol / AvgVolume
	if strings.Contains(name, "ETF") || strings.Contains(name, "Warrant") {
		wg.Done()
		return
	}
	if RelativeVolume >= 3 && !storedEvents.ContainsTicker(ticker) && DayChange > 0.07 {
		mu.Lock()
		newEvent := Event{
			Type: "",
			Time: time.Now(),
			ScanResult: ScanResult{
				Name:           name,
				Ticker:         ticker,
				AvgVolume:      AvgVolume,
				AvgChange:      AvgChange,
				DayChange:      DayChange,
				LastPrice:      LastPrice,
				CurrentVolume:  currentVol,
				RelativeVolume: RelativeVolume,
			},
		}

		*storedEvents = append(*storedEvents, newEvent)

		eventChan <- newEvent

		mu.Unlock()
	}

	defer wg.Done()
}

func (w *Watcher) Events() <-chan Event {
	return w.events
}

func (w *Watcher) LogEvents() {
	pretty.Println(w.storedEvents)
}

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			AddSource:  true,
		}),
	))
	watcher := NewWatcher(WatcherConfig{
		PollInterval: 5 * time.Minute,
		ApiKey:       "PKXRB7R43VGCYC9CX34R",
		SecretKey:    "DHhpzJBYejT00v913aDpbJKViSWcI7QaJOAy1zp3",
	})
	go watcher.Start()

	for event := range watcher.Events() {
		slog.Info("New event", "ticker", event.Ticker, "name", event.Name, "DayChange", event.DayChange, "RelativeVolume", event.RelativeVolume)
	}

	defer watcher.LogEvents()
}
