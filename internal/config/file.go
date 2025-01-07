package config

import (
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func configFolderPath() string {
	currentUser, err := user.Current()
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed to get current user: %s", err.Error()))
		return ""
	}

	homeDir := currentUser.HomeDir
	logs.Logger.Info("Home directory: " + homeDir)

	return filepath.Join(homeDir, ".config", "cherry-markdown")
}

func configFilePath() string {
	return filepath.Join(configFolderPath(), "config.json")
}

func init() {
	folderPath := configFolderPath()
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed to get current user: %s", err.Error()))
		return
	}
	logs.Logger.Info("Folder created successfully: " + folderPath)

	_, err = os.Stat(configFilePath())
	if err != nil && os.IsNotExist(err) {
		// 生成默认配置文件
		conf = &Config{
			Web: Web{},
			PicBed: PicBed{
				BasePath:  "/cherrymarkdown",
				Activated: Base64,
			},
			Theme: Theme{
				Activated: "default",
			},
		}
		conf.UpsertConfigFile()
	} else if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed to get config file: %s", err.Error()))
		return
	} else {
		// 解析配置文件
		confJsondata, err := file.ReadFileToByteArray(configFilePath())
		if err != nil {
			logs.Logger.Fatal(fmt.Sprintf("Failed read config.json: %s", err.Error()))
			return
		}

		config := Config{}
		err = json.Unmarshal(confJsondata, &config)
		if err != nil {
			logs.Logger.Fatal(fmt.Sprintf("Failed unmarshal JSON: %s", err.Error()))
			return
		}
		// 遍历主题列表
		config.Theme.ExtThemes = scanerExtTheme()
		conf = &config
	}

}
