package appmenu

import (
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func init() {
	fileMenu := appMenu.AddSubmenu("文件")
	fileMenu.AddText("打开", keys.CmdOrCtrl("o"), func(cd *menu.CallbackData) {
		menuFunc.OpenFileEvent()
	})

	fileMenu.AddText("保存", keys.CmdOrCtrl("s"), func(cd *menu.CallbackData) {
		menuFunc.SaveFileEvent()
	})

	fileMenu.AddText("另存为", keys.Combo("s", keys.ControlKey, keys.ShiftKey), func(cd *menu.CallbackData) {
		menuFunc.SaveAsFileEvent()
	})

	fileMenu.AddText("退出", keys.CmdOrCtrl("q"), func(cd *menu.CallbackData) {
		menuFunc.Quit()
	})

	exportMenu := appMenu.AddSubmenu("导出")

	exportMenu.AddText("导出pdf", nil, func(cd *menu.CallbackData) {
		menuFunc.ExportPdfEvent()
	})
	exportMenu.AddText("导出html", nil, func(cd *menu.CallbackData) {
		menuFunc.ExportHtmlEvent()
	})

	if runtime.GOOS == "darwin" {
		appMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}
}
