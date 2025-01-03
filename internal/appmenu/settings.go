package appmenu

import "github.com/wailsapp/wails/v2/pkg/menu"

func initSettings(menuFunc SettingsMenu) {
	settingsMenu := appMenu.AddSubmenu("设置")
	settingsMenu.AddText("首选项", nil, func(cd *menu.CallbackData) {
		menuFunc.OptionsEvent()
	})
	settingsMenu.AddText("个性化", nil, func(cd *menu.CallbackData) {
		menuFunc.PersonalizaEvent()
	})
}
