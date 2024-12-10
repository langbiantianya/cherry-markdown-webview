package appmenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type AppMenuFunc interface {
	OpenFileEvent()
	SaveFileEvent()
	SaveAsFileEvent()
	ExportPdfEvent()
	ExportHtmlEvent()
}

var appMenu = menu.NewMenu()
var menuFunc AppMenuFunc

func NewAppMenu(appMenuFunc AppMenuFunc) *menu.Menu {
	menuFunc = appMenuFunc
	return appMenu
}
