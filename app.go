package main

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"cherry-markdown-webview/internal/picbed"
	"cherry-markdown-webview/internal/quitext"
	"cherry-markdown-webview/public/utils"
	"path/filepath"
	"time"

	"context"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
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

// ExportHtmlEvent implements appmenu.ExportMenu.
func (a *App) ExportHtmlEvent() {
	wailsRuntime.EventsEmit(a.ctx, "exportHtmlEvent")
}

// ExportPdfEvent implements appmenu.ExportMenu.
func (a *App) ExportPdfEvent() {
	wailsRuntime.EventsEmit(a.ctx, "exportPdfEvent")
}

// SaveAsFileEvent implements appmenu.FileMenu.
func (a *App) SaveAsFileEvent() {
	wailsRuntime.EventsEmit(a.ctx, "saveAsFileEvent")
}

// SaveFileEvent implements appmenu.FileMenu.
func (a *App) SaveFileEvent() {
	wailsRuntime.EventsEmit(a.ctx, "saveFileEvent")
}

// OpenFileEvent implements appmenu.FileMenu.
func (a *App) OpenFileEvent() {
	wailsRuntime.EventsEmit(a.ctx, "openFileEvent")
}

// Quit implements appmenu.FileMenu.
func (a *App) Quit() {
	wailsRuntime.Quit(a.ctx)
}

// OptionsEvent implements appmenu.SettingsMenu.
func (a *App) OptionsEvent() {
	wailsRuntime.EventsEmit(a.ctx, "optionsEvent")
}

// PersonalizaEvent implements appmenu.SettingsMenu.
func (a *App) PersonalizaEvent() {
	wailsRuntime.EventsEmit(a.ctx, "personalizaEvent")
}

// AboutEvent implements appmenu.HelpMenu.
func (a *App) AboutEvent() {
	wailsRuntime.EventsEmit(a.ctx, "aboutEvent")
}

// IssuesEvent implements appmenu.HelpMenu.
func (a *App) IssuesEvent() {
	wailsRuntime.BrowserOpenURL(a.ctx, "https://github.com/langbiantianya/cherry-markdown-webview/issues")
}

func (a *App) QuitEvent() {
	wailsRuntime.EventsEmit(a.ctx, "quitEvent")
}

func (a *App) SetSaved(save bool) {
	quitext.Saved = save
}

func (a *App) GetSaved() bool {
	return quitext.Saved
}

func (a *App) AssociateOpen() file.File {
	return *file.GetFile()
}

func (a *App) OpenFile() {
	quitext.Saved = false
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
	if filePath != "" {
		file.AsynLoadingToRam(filePath)
		time.Sleep(1000 * time.Nanosecond)
		wailsRuntime.WindowReload(a.ctx)
	}
}

func (a *App) SaveFile(doc file.File) *file.File {
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
			return nil
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
		return nil
	}
	wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
		Type:    "info",
		Title:   "成功",
		Message: "保存成功",
	})
	quitext.Saved = true
	if doc.Name == "" {
		doc.Name = filepath.Base(doc.Path)
	}
	return &doc
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

func (a *App) SelectLocalFile(doc file.File) file.File {
	displayName := "Markdown file"
	pattern := "*.md;*.mdx"
	if doc.Pattern != "" {
		pattern = doc.Pattern
	}
	if doc.DisplayName != "" {
		displayName = doc.DisplayName
	}

	filePath, err := wailsRuntime.OpenFileDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title: "选择文件",
		Filters: []wailsRuntime.FileFilter{
			{
				DisplayName: displayName,
				Pattern:     pattern,
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
		return file.File{}
	}
	localFile, err := file.ReadFile(filePath)
	if err != nil {
		logs.Logger.Error(err.Error())
		return file.File{}
	}
	return *localFile
}

func (a *App) UpsertPicBed(picBed config.PicBed) {
	config.GetConfig().UpsertPicBed(picBed)
}
func (a *App) GetPicBed() config.PicBed {
	return config.GetConfig().PicBed
}

func (a *App) UploadPicbed(sourceFile file.File) string {
	url, err := picbed.GetPicbed().Upload(sourceFile)
	if err != nil {
		wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    "error",
			Title:   "错误",
			Message: err.Error(),
		})
		return ""
	}
	if url != nil {
		return url.String()
	}
	return ""
}

func (a *App) GetActivatedTheme() string {
	return config.GetConfig().GetActivatedTheme()
}

func (a *App) LoadTheme(name string) *config.ThemeItem {
	themeItem, err := config.LoadTheme(name)
	if err != nil {
		wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    "error",
			Title:   "错误",
			Message: err.Error(),
		})
		return nil
	}
	return themeItem
}

func (a *App) UpsertThemeItem(theme config.ThemeItem) {
	err := config.UpsertThemeItem(theme)
	if err != nil {
		wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    "error",
			Title:   "错误",
			Message: err.Error(),
		})
	}
}

func (a *App) GetThemeList() []config.ExtTheme {
	return config.GetConfig().Theme.ExtThemes
}

func (a *App) SetBackground(img file.File) {
	config.GetConfig().SetBackgroundImage(img)
}

func (a *App) SetActivatedTheme(theme string) {
	config.GetConfig().SetActivatedTheme(theme)
}
