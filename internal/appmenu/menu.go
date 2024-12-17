package appmenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type AppMenuFunc interface {
	FileMenu
	ExportMenu
	SettingsMenu
	HelpMenu
}

type FileMenu interface {
	OpenFileEvent()
	SaveFileEvent()
	SaveAsFileEvent()
	Quit()
}

type ExportMenu interface {
	ExportPdfEvent()
	ExportHtmlEvent()
}

type SettingsMenu interface {
	OptionsEvent()
	PersonalizaEvent()
}

type HelpMenu interface {
	AboutEvent()
	IssuesEvent()
}

var appMenu = menu.NewMenu()
var appMenuFunc AppMenuFunc

func NewAppMenu(_appMenuFunc AppMenuFunc) *menu.Menu {
	appMenuFunc = _appMenuFunc
	initFile(appMenuFunc)
	initExport(appMenuFunc)
	initSettings(appMenuFunc)
	initHelp(appMenuFunc)
	return appMenu
}
