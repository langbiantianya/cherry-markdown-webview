package main

import (
	"cherry-markdown-webview/internal/appmenu"
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"cherry-markdown-webview/internal/quitext"
	"cherry-markdown-webview/internal/web"
	"context"

	"embed"
	"os"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	web.StartServer()
	app := NewApp()

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
	// Create application with options
	err := wails.Run(&options.App{
		Windows: &windows.Options{
			IsZoomControlEnabled: false,
			WebviewIsTransparent: true,
			BackdropType:         windows.Auto,
		},
		Title:     "cherry-markdown-webview",
		Width:     800,
		Height:    600,
		MinWidth:  768,
		MinHeight: 420,
		Menu:      appmenu.NewAppMenu(app),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "CherryMarkdowmOnWebView",
				Message: "© 2024 lbty",
			},
			// TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			OnFileOpen: func(filePaths string) {
				file.AsynLoadingToRam(filePaths)
			},
		},
		// Linux: &linux.Options{
		// WindowIsTranslucent: true,
		// },
		OnStartup: app.startup,
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			app.QuitEvent()
			//询问是否保存
			return !quitext.Saved
		},
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		logs.Logger.Fatal(err.Error())
	}
}
