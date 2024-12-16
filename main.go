package main

import (
	"embed"

	"github.com/leaanthony/u"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails",
		Width:  700,
		Height: 400,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			IsZoomControlEnabled: false,
			WebviewGpuIsDisabled: false,
		},

		Mac: &mac.Options{

			DisableZoom: false,
			Preferences: &mac.Preferences{
				TabFocusesLinks:        u.True,
				TextInteractionEnabled: u.True,
				FullscreenEnabled:      u.True,
			},
		},

		// WindowStartState: options.Maximised,
		WindowStartState: options.Minimised,

		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Logger: nil,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// package main

// import (
// 	"delta/pkg/generated/rti"
// 	"delta/pkg/rithmic"
// 	"flag"
// 	"log/slog"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/kr/pretty"
// 	"github.com/lmittmann/tint"
// )

// var (
// 	addr = flag.String("addr", " wss://rituz00100.rithmic.com:443", "address of the server")
// )

// func main() {
// 	slog.SetDefault(slog.New(
// 		tint.NewHandler(os.Stderr, &tint.Options{
// 			Level:      slog.LevelDebug,
// 			TimeFormat: time.Kitchen,
// 		}),
// 	))
// 	slog.Info("Starting DeltΔ...")

// 	url := "wss://rituz00100.rithmic.com:443"

// 	// usr := "mhdi.seddik@gmail.com"
// 	// pwd := "lDIKLQCX"
// 	usr := "xmhd"
// 	pwd := "TST563"

// 	rithmicWS := rithmic.New(rithmic.ConnectionArgs{
// 		Url:      url,
// 		User:     usr,
// 		Password: pwd,
// 	})

// 	rithmicWS.SubscribeMarketDataLastTrade("ESZ4", "CME", func(l rti.LastTrade) error {
// 		pretty.Println(l.String())
// 		return nil
// 	})

// 	stopChan := make(chan os.Signal, 1)
// 	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

// 	<-stopChan
// 	slog.Info("Shutting down DeltΔ...")

// }
