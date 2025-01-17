package main

import (
	"context"
	"delta/pkg/persistence"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/lmittmann/tint"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App struct
type App struct {
	ctx         context.Context
	Persistence *persistence.Persistence
	Alpaca      *alpaca.Client
}

func NewApp() *App {
	p, err := persistence.New("delta")
	if err != nil {
		slog.Info("error initializing persistence", err.Error())
	}

	appData, err := p.Load()
	if err != nil {
		slog.Error("error loading app data", err.Error())
	}

	client := alpaca.NewClient(alpaca.ClientOpts{
		// Alternatively you can set your key and secret using the
		// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables
		APIKey:    appData.Keys.ApiKey,
		APISecret: appData.Keys.SecretKey,
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	slog.Info("Starting DeltΔ...")

	return &App{
		Persistence: p,
		Alpaca:      client,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
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
		slog.Info("Login successful")
		a.Alpaca = client
		return true
	}
}

func (a *App) GetAssets() (tickers []alpaca.Asset, err error) {
	if a.Alpaca == nil {
		return nil, errors.New("not logged in")
	}

	assets, err := a.Alpaca.GetAssets(alpaca.GetAssetsRequest{
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
	a.Alpaca = nil
}

func (a *App) GetAccount() (*alpaca.Account, error) {
	if a.Alpaca == nil {
		slog.Error("not logged in")
		return nil, errors.New("not logged in")
	}
	acct, err := a.Alpaca.GetAccount()
	if err != nil {
		return nil, err
	}
	return acct, nil
}

func (a *App) GetAppData() (persistence.AppData, error) {
	return a.Persistence.Load()
}

func (a *App) SaveAppData(data persistence.AppData) error {
	return a.Persistence.Save(data)
}
