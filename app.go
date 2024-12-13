package main

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"cherry-markdown-webview/public/utils"
	"time"

	"context"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// ExportHtml implements appmenu.AppMenuFunc.
func (a *App) ExportHtmlEvent() {
	wailsRuntime.EventsEmit(a.ctx, "exportHtmlEvent")
}

// ExportPdf implements appmenu.AppMenuFunc.
func (a *App) ExportPdfEvent() {
	wailsRuntime.EventsEmit(a.ctx, "exportPdfEvent")
}

// SaveAsFile implements appmenu.AppMenuFunc.
func (a *App) SaveAsFileEvent() {
	wailsRuntime.EventsEmit(a.ctx, "saveAsFileEvent")
}

// SaveFile implements appmenu.AppMenuFunc.
func (a *App) SaveFileEvent() {
	wailsRuntime.EventsEmit(a.ctx, "saveFileEvent")
}

// OpenFileEvent implements appmenu.AppMenuFunc.
func (a *App) OpenFileEvent() {
	wailsRuntime.EventsEmit(a.ctx, "openFileEvent")
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AssociateOpen() file.File {
	return *file.GetFile()
}

func (a *App) OpenFile() {
	filePath, err := wailsRuntime.OpenFileDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title: "选择文件",
		Filters: []wailsRuntime.FileFilter{
			{
				DisplayName: "Markdown file",
				Pattern:     "*.md;*.mdx",
			},
		},
		CanCreateDirectories: true,
	})
	if err != nil {
		logs.Logger.Error(err.Error())
		wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    "error",
			Title:   "错误",
			Message: err.Error(),
		})
		return
	}
	file.AsynLoadingToRam(filePath)
	time.Sleep(1000 * time.Nanosecond)
	wailsRuntime.WindowReload(a.ctx)
}

func (a *App) SaveFile(doc file.File) {
	if doc.Path == "" {
		// 打开选项框
		displayName := "Markdown file"
		pattern := "*.md;*.mdx"
		if doc.Pattern != "" {
			pattern = doc.Pattern
		}
		if doc.DisplayName != "" {
			displayName = doc.DisplayName
		}
		filePath, err := wailsRuntime.SaveFileDialog(a.ctx, wailsRuntime.SaveDialogOptions{
			DefaultFilename: doc.Name,
			Title:           "保存",
			Filters: []wailsRuntime.FileFilter{
				{
					DisplayName: displayName,
					Pattern:     pattern,
				},
			},
		})
		if err != nil {
			logs.Logger.Error(err.Error())
			wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
				Type:    "error",
				Title:   "错误",
				Message: err.Error(),
			})
		}
		doc.Path = filePath
	}
	// 保存
	err := file.WriteFile(doc)
	if err != nil {
		logs.Logger.Error(err.Error())
		wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    "error",
			Title:   "错误",
			Message: err.Error(),
		})
	}
	wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
		Type:    "info",
		Title:   "成功",
		Message: "保存成功",
	})

}

func (a *App) ReadLocalFile(uri string) file.File {
	parsedURI, err := utils.ParsedURI(uri)
	if err != nil {
		logs.Logger.Error(err.Error())
		return file.File{}
	}
	localFile, err := file.ReadFile(parsedURI.Path[1:])
	if err != nil {
		logs.Logger.Error(err.Error())
		return file.File{}
	}
	return *localFile

}

func (a *App) GetWebServerPort() int {
	return config.GetConfig().Web.GetPort()
}
