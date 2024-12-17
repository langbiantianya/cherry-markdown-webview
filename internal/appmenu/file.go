package appmenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func initFile(menuFunc FileMenu) {
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

}
