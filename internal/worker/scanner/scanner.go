package scanner

import (
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

type ScanResult struct {
	Asset          *alpaca.Asset
	AvgVolume      int
	CurrentVolume  int
	RelativeVolume int
	AvgChange      float64
	DayChange      float64
	LastPrice      float64
}

type ScanResults []ScanResult

type ScannerConfig struct {
	Exchange   string
	AssetClass string
}

type Scanner struct {
	config        ScannerConfig
	dataClient    *marketdata.Client
	tradingClient *alpaca.Client
	subscribers   []Observer
	scanResults   *ScanResults
}

func New(config ScannerConfig) *Scanner {
	if config.Exchange == "" {
		config.Exchange = "NASDAQ"
	}
	if config.AssetClass == "" {
		config.AssetClass = string(alpaca.USEquity)
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	marketDataUrl := os.Getenv("MARKET_DATA_URL")
	routingUrl := os.Getenv("ROUTING_URL")

	marketData := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   marketDataUrl,
	})

	tradingClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   routingUrl,
	})

	return &Scanner{
		config:        config,
		dataClient:    marketData,
		tradingClient: tradingClient,
		scanResults:   &ScanResults{},
		subscribers:   make([]Observer, 0),
	}
}

func (s *Scanner) Start() {
	assets, err := s.tradingClient.GetAssets(alpaca.GetAssetsRequest{
		AssetClass: s.config.AssetClass,
		Status:     string(alpaca.AssetActive),
		Exchange:   s.config.Exchange,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	assetNames := make([]string, 0)
	for _, asset := range assets {
		assetNames = append(assetNames, asset.Symbol)
		*s.scanResults = append(*s.scanResults, ScanResult{
			Asset: &asset,
		})
	}

	startDate := time.Now().AddDate(0, 0, -20)
	barsData, err := s.dataClient.GetMultiBars(assetNames, marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneDay,
		Feed:      marketdata.IEX,
		Start:     startDate,
	})

	if err != nil {
		slog.Error("Error getting bars", "err", err.Error())
	}
	for {
		snapshots, err := s.dataClient.GetSnapshots(assetNames, marketdata.GetSnapshotRequest{
			Feed: marketdata.IEX,
		})
		if err != nil {
			slog.Error("Error getting bars", "err", err.Error())
		}

		var wg sync.WaitGroup
		var mu sync.Mutex

		for symbol, bars := range barsData {
			if len(bars) == 0 {
				slog.Warn("No bars for asset", "ticker", symbol)
				continue
			}

			var asset alpaca.Asset
			for _, a := range assets {
				if a.Symbol == symbol {
					asset = a
					break
				}
			}
			wg.Add(1)
			go s.computeStats(&wg, &mu, s.scanResults, &asset, bars, snapshots[symbol])
		}

		wg.Wait()
		s.Notify()

		time.Sleep(5 * time.Second)
	}
}

func (s *Scanner) computeStats(wg *sync.WaitGroup, mu *sync.Mutex, scanResults *ScanResults, asset *alpaca.Asset, bars []marketdata.Bar, snapshot *marketdata.Snapshot) {
	// || snapshot.DailyBar.Timestamp.Day() != time.Now().Day()
	if snapshot == nil {
		wg.Done()
		return
	}

	var (
		AvgVolume int
		AvgChange float64
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
	currentVolume := int(snapshot.DailyBar.Volume)
	LastPrice := snapshot.DailyBar.Close
	DayChange := (snapshot.DailyBar.Close - snapshot.DailyBar.Open) / snapshot.DailyBar.Open
	RelativeVolume := currentVolume / AvgVolume

	ScanResult := ScanResult{
		Asset:          asset,
		AvgVolume:      AvgVolume,
		CurrentVolume:  currentVolume,
		RelativeVolume: RelativeVolume,
		AvgChange:      AvgChange,
		DayChange:      DayChange,
		LastPrice:      LastPrice,
	}

	mu.Lock()
	// *scanResults = append(*scanResults, ScanResult)
	// au lieu d'ajouter, on remplace l'élément du tableau qui correspond à l'asset
	for i, result := range *scanResults {
		if result.Asset.Symbol == asset.Symbol {
			(*scanResults)[i] = ScanResult
			break
		}
	}

	mu.Unlock()

	defer wg.Done()
}

func (s *Scanner) Notify() {
	for _, subscriber := range s.subscribers {
		subscriber.NotifyScan(*s.scanResults)
	}
}

func (s *Scanner) Subscribe(observer Observer) {
	s.subscribers = append(s.subscribers, observer)
}

func (s *Scanner) Unsubscribe(observer Observer) {
	for i, subscriber := range s.subscribers {
		if subscriber == observer {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			break
		}
	}
}

func (s *Scanner) GetResults() ScanResults {
	return *s.scanResults
}
