package appmenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func initExport(menuFunc ExportMenu) {

	exportMenu := appMenu.AddSubmenu("导出")

	exportMenu.AddText("导出pdf", nil, func(cd *menu.CallbackData) {
		menuFunc.ExportPdfEvent()
	})
	exportMenu.AddText("导出html", nil, func(cd *menu.CallbackData) {
		menuFunc.ExportHtmlEvent()
	})
}
