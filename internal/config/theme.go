package config

import (
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ThemeItemConf struct {
	Name                string   `json:"name"`
	Colors              []string `json:"colors"`
	ScssPath            string   `json:"scss"`
	ToolBarPath         string   `json:"toolBar"`
	BackgroundImagePath string   `json:"backgroundImage"`
}

type ThemeItem struct {
	Name           ExtTheme   `json:"name"`
	Colors         []string   `json:"colors"`
	Scss           string     `json:"scss"`
	ToolBar        string     `json:"toolBar"`
	BaclgroudImage *file.File `json:"backgroundImage"`
}

func getBasehemePath() string {
	return filepath.Join(configFolderPath(), "theme")
}

func LoadThemeConf(name string) (ThemeItemConf, error) {
	baseThemePath := filepath.Join(getBasehemePath(), name, "theme.json")
	themeData, err := file.ReadFileToByteArray(baseThemePath)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed read config.json: %s", err.Error()))
		return ThemeItemConf{}, err
	}

	themeConf := ThemeItemConf{}

	err = json.Unmarshal(themeData, &themeConf)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed unmarshal JSON: %s", err.Error()))
		return ThemeItemConf{}, err
	}
	return themeConf, nil
}

func LoadTheme(name string) (*ThemeItem, error) {
	themeConf, err := LoadThemeConf(name)
	if err != nil {
		return nil, err
	}
	themeItem := ThemeItem{
		Name: ExtTheme{
			ClassName: themeConf.Name,
			Label:     name,
		},
		Colors: themeConf.Colors,
	}
	// 加载scss
	scssByteArray, err := file.ReadFileToByteArray(filepath.Join(getBasehemePath(), name, themeConf.ScssPath))
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed load scss: %s", err.Error()))
		return nil, err
	}
	themeItem.Scss = string(scssByteArray)

	toolBarBytearry, err := file.ReadFileToByteArray(filepath.Join(getBasehemePath(), name, themeConf.ToolBarPath))
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed load scss: %s", err.Error()))
		return nil, err
	}

	themeItem.ToolBar = string(toolBarBytearry)
	// 加载背景图片
	if themeConf.BackgroundImagePath != "" {
		bgImg, err := file.ReadFile(filepath.Join(getBasehemePath(), name, themeConf.BackgroundImagePath))
		if err != nil {
			logs.Logger.Error(fmt.Sprintf("Failed load background image: %s", err.Error()))
			return nil, err
		}
		themeItem.BaclgroudImage = bgImg
	}
	return &themeItem, nil
}

func UpsertThemeItem(theme ThemeItem) error {
	// 判断主题是否存在
	themePath := filepath.Join(getBasehemePath(), theme.Name.Label)
	_, err := os.Stat(themePath)
	var themeItemConf ThemeItemConf
	if os.IsNotExist(err) {
		os.MkdirAll(themePath, 0755)
		// 新建配置文件
		themeItemConf = ThemeItemConf{
			Name:        theme.Name.ClassName,
			Colors:      theme.Colors,
			ScssPath:    "theme.scss",
			ToolBarPath: "toolbar.css",
		}
		if theme.BaclgroudImage != nil {
			themeItemConf.BackgroundImagePath = theme.BaclgroudImage.Name
		}
	} else if err != nil {
		logs.Logger.Error(err.Error())
		return err
	} else {
		// 加载旧配置
		themeItemConf, err = LoadThemeConf(theme.Name.Label)
		if err != nil {
			logs.Logger.Error(fmt.Sprintf("Failed load theme: %s", err.Error()))
			return err
		}
	}
	// 保存scss
	err = file.WireByteArray(filepath.Join(themePath, themeItemConf.ScssPath), []byte(theme.Scss))
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed save scss: %s", err.Error()))
		return err

	}
	// 保存toolbar

	err = file.WireByteArray(filepath.Join(themePath, themeItemConf.ToolBarPath), []byte(theme.ToolBar))
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed save css: %s", err.Error()))
		return err

	}

	// 保存背景图片
	if theme.BaclgroudImage != nil {
		err = file.WireByteArray(filepath.Join(themePath, themeItemConf.BackgroundImagePath), theme.BaclgroudImage.Bytes)
		if err != nil {
			logs.Logger.Error(fmt.Sprintf("Failed save background image: %s", err.Error()))
			return err
		}
	}
	// 更新json
	themeItemConfJson, err := json.Marshal(themeItemConf)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed marshal theme: %s", err.Error()))
		return err
	}
	err = file.WireByteArray(filepath.Join(themePath, "theme.json"), themeItemConfJson)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("Failed save theme: %s", err.Error()))
		return err
	}
	return nil
}

func scanerExtTheme() []ExtTheme {
	theme := make([]ExtTheme, 0)
	err := os.MkdirAll(getBasehemePath(), 0755)

	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed makdir theme conf path: %s", err.Error()))
		return theme
	}

	err = filepath.Walk(getBasehemePath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只处理目录
		if info.IsDir() {
			themeConf, err := LoadThemeConf(info.Name())
			if err != nil {
				logs.Logger.Error(fmt.Sprintf("Failed load theme: %s", err.Error()))
				return nil
			}
			theme = append(theme, ExtTheme{
				ClassName: themeConf.Name,
				Label:     info.Name(),
			})
		}
		return nil
	})
	if err != nil {
		logs.Logger.Error(err.Error())
		return theme
	}
	return theme
}
