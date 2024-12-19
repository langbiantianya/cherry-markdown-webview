package config

import (
	"cherry-markdown-webview/internal/logs"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func ConfigFolderPath() string {
	currentUser, err := user.Current()
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed to get current user: %s", err.Error()))
		return ""
	}

	homeDir := currentUser.HomeDir
	logs.Logger.Info("Home directory: " + homeDir)

	return filepath.Join(homeDir, ".config", "cherry-markdown")
}

func ConfigFilePath() string {
	return filepath.Join(ConfigFolderPath(), "config.json")
}

func init() {
	folderPath := ConfigFolderPath()
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed to get current user: %s", err.Error()))
		return
	}
	logs.Logger.Info("Folder created successfully: " + folderPath)

	_, err = os.Stat(ConfigFilePath())
	if err != nil && os.IsNotExist(err) {
		// 生成默认配置文件
		conf = &Config{
			Web: Web{},
			OSS: OSS{},
		}
		conf.UpsertConfigFile()
	} else if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed to get config file: %s", err.Error()))
		return
	} else {
		// 解析配置文件
		confFile, err := os.Open(ConfigFilePath())
		if err != nil {
			logs.Logger.Fatal(fmt.Sprintf("Failed open File %s", err.Error()))
			return
		}
		defer confFile.Close()

		confJsondata, err := io.ReadAll(confFile)
		if err != nil {
			logs.Logger.Fatal(fmt.Sprintf("Failed read config.json: %s", err.Error()))
			return
		}

		config := &Config{}
		err = json.Unmarshal(confJsondata, config)
		if err != nil {
			logs.Logger.Fatal(fmt.Sprintf("Failed unmarshal JSON: %s", err.Error()))
			return
		}
		conf = config
	}

}
