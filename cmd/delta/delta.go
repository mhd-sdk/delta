package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/delta/internal/env"
	"github.com/delta/internal/server"
	"github.com/lmittmann/tint"
)

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	slog.Info("Starting Delta service")

	err := env.LoadEnv()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	server := server.New()

	err = server.TestCredentials()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	server.StartScanner()

	server.ServeAPI()

	fmt.Println("Stop")
}
