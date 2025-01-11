package main

import (
	"context"
	"delta/pkg/persistence"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App struct
type App struct {
	ctx         context.Context
	Persistence *persistence.Persistence
}

func NewApp() *App {
	p, err := persistence.New("delta")
	if err != nil {
		slog.Info("erreur")
	}

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	slog.Info("Starting DeltΔ...")

	return &App{
		Persistence: p,
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

func (a *App) Ping() string {
	return "pong"
}

func (a *App) GetAppData() (persistence.AppData, error) {
	return a.Persistence.Load()
}

func (a *App) SaveAppData(data persistence.AppData) error {
	return a.Persistence.Save(data)
}
