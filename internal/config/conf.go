package config

import (
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var conf *Config

func GetConfig() *Config {
	return conf
}

func (w *Web) GetPort() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.port
}

func (w *Web) SetPort(port int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.port = port
}

type Config struct {
	Web    Web    `json:"-"`
	PicBed PicBed `json:"picBed"`
	Theme  Theme  `json:"theme"`
}

func (conf *Config) SetActivatedTheme(theme string) {
	conf.Theme.Activated = theme
	conf.UpsertConfigFile()
}

func (conf *Config) SetBackgroundImage(img file.File) {
	if img.Path != "" {
		conf.Theme.BackgroundImage = img.Path
	} else {
		bgFolderPath := filepath.Join(configFolderPath(), "background")
		err := os.MkdirAll(bgFolderPath, 0755)
		if err != nil {
			logs.Logger.Error(fmt.Sprintf("Failed to mkdir path: %s", err.Error()))
			return
		}
		bgPath := filepath.Join(bgFolderPath, img.Name)
		err = file.WireByteArray(bgPath, img.Bytes)
		if err != nil {
			logs.Logger.Error(fmt.Sprintf("Failed to save background: %s", err.Error()))
			return
		}
		conf.Theme.BackgroundImage = bgPath
	}
	conf.UpsertConfigFile()
}

func (conf *Config) GetActivatedTheme() string {
	return conf.Theme.Activated
}

func (conf *Config) UpsertConfigFile() {
	confJsonData, err := json.Marshal(conf)
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed marshaling JSON: %s", err.Error()))
		return
	}
	confFile, err := os.OpenFile(configFilePath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed open File %s", err.Error()))
		return
	}
	defer confFile.Close()
	_, err = confFile.Write(confJsonData)
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed write File %s", err.Error()))
		return
	}
	logs.Logger.Info("Config.json upsert successfully: " + configFilePath())
}

func (conf *Config) UpsertPicBed(picBed PicBed) {
	conf.PicBed = picBed
	conf.UpsertConfigFile()
}

type Theme struct {
	Activated       string     `json:"activated"`
	BackgroundImage string     `json:"backgroundImage"`
	ExtThemes       []ExtTheme `json:"-"`
}

type ExtTheme struct {
	ClassName string `json:"className"`
	Label     string `json:"label"`
}

type Web struct {
	mu   sync.Mutex `json:"-"`
	port int        `json:"-"`
}

type PicBedType string

const (
	Base64 PicBedType = "Base64"
	OSS    PicBedType = "OSS"
	COS    PicBedType = "COS"
)

type PicBed struct {
	BasePath  string           `json:"basePath"`
	Activated PicBedType       `json:"activated"`
	OSS       *AliyunOSS       `json:"oss"`
	COS       *TencentCloudCOS `json:"cos"`
}

type AliyunOSS struct {
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	Region          string `json:"region"`
	BucketName      string `json:"bucketName"`
}

type TencentCloudCOS struct {
	SecretID  string `json:"secretID"`
	SecretKey string `json:"secretKey"`
	BucketURL string `json:"bucketURL"`
}
