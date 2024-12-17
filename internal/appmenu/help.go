package appmenu

import "github.com/wailsapp/wails/v2/pkg/menu"

func initHelp(menuFunc HelpMenu) {
	helpMenu := appMenu.AddSubmenu("帮助")
	helpMenu.AddText("报告错误", nil, func(cd *menu.CallbackData) {
		menuFunc.IssuesEvent()
	})
	helpMenu.AddText("关于", nil, func(cd *menu.CallbackData) {
		menuFunc.AboutEvent()
	})
}
