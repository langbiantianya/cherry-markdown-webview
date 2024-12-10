package main

import (
	"cherry-markdown-webview/internal/appmenu"
	"cherry-markdown-webview/internal/file"
	"embed"
	"os"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	switch runtime.GOOS {
	case "linux":
		argsWithoutProg := os.Args[1:]
		if len(argsWithoutProg) != 0 {
			file.AsynLoadingToRam(argsWithoutProg[0])
		}

	case "windows":
		argsWithoutProg := os.Args[1:]
		if len(argsWithoutProg) != 0 {
			file.AsynLoadingToRam(argsWithoutProg[0])
		}
	}
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Windows: &windows.Options{
			IsZoomControlEnabled: false,
		},

		Title:  "cherry-markdown-webview",
		Width:  800,
		Height: 600,
		Menu:   appmenu.NewAppMenu(app),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Mac: &mac.Options{
			OnFileOpen: func(filePaths string) {
				file.AsynLoadingToRam(filePaths)
			},
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
