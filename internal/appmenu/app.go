package appmenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type AppMenuFunc interface {
	FileMenu
	ExportMenu
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

var appMenu = menu.NewMenu()
var menuFunc AppMenuFunc

func NewAppMenu(appMenuFunc AppMenuFunc) *menu.Menu {
	menuFunc = appMenuFunc
	return appMenu
}
