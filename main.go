package main

import (
	"delta/pkg/app"
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
	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "DeltΔ",
		Width:  1200,
		Height: 800,

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

		WindowStartState: options.Minimised,
		// WindowStartState: options.Minimised,

		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "e3984e08-28dc-4e3d-b70a-45e961589cdc",
			OnSecondInstanceLaunch: app.OnSecondInstanceLaunch,
		},
		Bind: []interface{}{
			app,
		},
		Logger: nil,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
