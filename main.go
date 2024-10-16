package main

import (
	"delta/pkg/generated/rti"
	"delta/pkg/rithmic"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kr/pretty"
	"github.com/lmittmann/tint"
)

var (
	addr = flag.String("addr", " wss://rituz00100.rithmic.com:443", "address of the server")
)

func main() {
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

	rithmicWS := rithmic.New(rithmic.ConnectionArgs{
		Url:      url,
		User:     usr,
		Password: pwd,
	})

	rithmicWS.SubscribeMarketDataLastTrade("ESZ4", "CME", func(l rti.LastTrade) error {
		pretty.Println(l.String())
		return nil
	})

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan
	slog.Info("Shutting down DeltΔ...")

}
