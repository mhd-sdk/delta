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
	// url := "wss://rprotocol-de.rithmic.com:443"

	r, err := rithmic.New(url)
	if err != nil {
		slog.Error("Error creating RithmicWS", "error", err)
	}

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

func (a *App) GetAppData() (persistence.AppData, error) {
	return a.Persistence.Load()
}

func (a *App) SaveAppData(data persistence.AppData) error {
	return a.Persistence.Save(data)
}

func (a *App) GetProducts() ([]*rti.ResponseProductCodes, error) {
	return a.RithmicWs.ListProducts()
}

func (a *App) GetSystems() (*rti.ResponseRithmicSystemInfo, error) {
	return a.RithmicWs.ListSystems()
}

type loginArgs struct {
	Username string
	Password string
	System   string
}

func (a *App) Login(args loginArgs) error {

	// usr := "mhdi.seddik@gmail.com"
	// usr := "xmhd"
	// pwd := "lDIKLQCX"
	// pwd := "TST563"
	// system := "TopstepTrader"

	connArgs := rithmic.ConnectionArgs{
		User:       args.Username,
		Password:   args.Password,
		SystemName: args.System,
	}

	return a.RithmicWs.Login(connArgs)
}
