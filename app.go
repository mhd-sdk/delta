package main

import (
	"context"
	"delta/pkg/persistence"
	"log/slog"
)

// App struct
type App struct {
	ctx         context.Context
	Persistence *persistence.Persistence
}

// NewApp creates a new App application struct
func NewApp() *App {
	p, err := persistence.New("delta")
	if err != nil {
		slog.Info("erreur")
	}
	return &App{
		Persistence: p,
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
