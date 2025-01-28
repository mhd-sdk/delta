package app

import (
	"context"
	"delta/pkg/models"
	"delta/pkg/persistence"
	"errors"
	"log/slog"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App struct
type App struct {
	ctx              context.Context
	Persistence      *persistence.Persistence
	TradingClient    *alpaca.Client
	MarketDataClient *marketdata.Client
}

func NewApp() *App {
	p, err := persistence.New("delta")
	if err != nil {
		slog.Error("error initializing persistence", err.Error())
	}

	appData, err := p.Load()
	if err != nil {
		slog.Error("error loading app data", err.Error())
	}

	tradingClient := alpaca.NewClient(alpaca.ClientOpts{
		// Alternatively you can set your key and secret using the
		// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables
		APIKey:    appData.Keys.ApiKey,
		APISecret: appData.Keys.SecretKey,
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	marketDataClient := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    appData.Keys.ApiKey,
		APISecret: appData.Keys.SecretKey,
		BaseURL:   "https://data.alpaca.markets",
	})

	return &App{
		Persistence:      p,
		TradingClient:    tradingClient,
		MarketDataClient: marketDataClient,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OnSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	slog.Info("Second instance launched")
}

func (a *App) TestCredentials(key string, secret string) bool {
	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    key,
		APISecret: secret,
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	_, err := client.GetAccount()

	if err != nil {
		slog.Error("login failed", "error", err.Error())
		return false
	} else {
		a.TradingClient = client
		return true
	}
}

func (a *App) GetAssets() (tickers []alpaca.Asset, err error) {
	if a.TradingClient == nil {
		return nil, errors.New("not logged in")
	}

	assets, err := a.TradingClient.GetAssets(alpaca.GetAssetsRequest{
		Status:     "active",
		AssetClass: "us_equity",
	})
	if err != nil {
		slog.Error("error fetching assets", "error", err.Error())
		return nil, err
	}
	return assets, nil
}

func (a *App) Logout() {
	a.TradingClient = nil
}

func (a *App) GetAccount() (*alpaca.Account, error) {
	if a.TradingClient == nil {
		slog.Error("not logged in")
		return nil, errors.New("not logged in")
	}
	acct, err := a.TradingClient.GetAccount()
	if err != nil {
		return nil, err
	}
	return acct, nil
}

func (a *App) GetAppData() (models.AppData, error) {
	return a.Persistence.Load()
}

func (a *App) SaveAppData(data models.AppData) error {
	return a.Persistence.Save(data)
}

func (a *App) ResetPreferences() error {
	return a.Persistence.ResetPreferences()
}

type GetCandlesticksConfig struct {
	Ticker    string
	Start     time.Time
	End       time.Time
	Timeframe models.Timeframe `json:"timeframe"`
}

func (a *App) GetCandlesticks(config GetCandlesticksConfig) (data []marketdata.Bar, err error) {
	if a.MarketDataClient == nil {
		return nil, errors.New("not logged in")
	}
	data, err = a.MarketDataClient.GetBars(config.Ticker, marketdata.GetBarsRequest{
		TimeFrame: marketdata.TimeFrame{
			N:    config.Timeframe.N,
			Unit: marketdata.TimeFrameUnit(config.Timeframe.Unit),
		},
		Start: config.Start,
		// End:   config.End,
		Feed: marketdata.IEX,
	})
	if err != nil {
		slog.Error("error fetching assets", "error", err.Error())
		return nil, err
	}

	return data, nil
}
