package main

import (
	"context"
	"delta/pkg/generated/rti"
	"delta/pkg/persistence"
	"delta/pkg/rithmic"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// App struct
type App struct {
	ctx         context.Context
	Persistence *persistence.Persistence
	RithmicWs   *rithmic.RithmicWS
}

// NewApp creates a new App application struct
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

	url := "wss://rituz00100.rithmic.com:443"

	usr := "mhdi.seddik@gmail.com"
	pwd := "lDIKLQCX"

	r := rithmic.New(rithmic.ConnectionArgs{
		Url:      url,
		User:     usr,
		Password: pwd,
	})

	return &App{
		Persistence: p,
		RithmicWs:   r,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Ping() string {
	return "pong"
}

func (a *App) Load() (persistence.AppData, error) {
	return a.Persistence.Load()
}

func (a *App) Save(data persistence.AppData) error {
	return a.Persistence.Save(data)
}

func (a *App) GetProducts() ([]*rti.ResponseProductCodes, error) {
	return a.RithmicWs.ListProducts()
}
