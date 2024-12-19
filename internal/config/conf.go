package config

import (
	"cherry-markdown-webview/internal/logs"
	"encoding/json"
	"fmt"
	"os"
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
	Web Web `json:"-"`
	OSS OSS `json:"oss"`
}

func (conf *Config) UpsertConfigFile() {
	confJsonData, err := json.Marshal(conf)
	if err != nil {
		logs.Logger.Fatal(fmt.Sprintf("Failed marshaling JSON: %s", err.Error()))
		return
	}
	confFile, err := os.OpenFile(ConfigFilePath(), os.O_CREATE|os.O_WRONLY, 0644)
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
	logs.Logger.Info("Config.json upsert successfully: " + ConfigFilePath())
}

func (conf *Config) UpsertOss(oss OSS) {
	conf.OSS = oss
	conf.UpsertConfigFile()
}

type Web struct {
	mu   sync.Mutex `json:"-"`
	port int        `json:"-"`
}

type OSS struct {
	BasePath     string        `json:"basePath"`
	Aliyun       *Aliyun       `json:"aliyun"`
	TencentCloud *TencentCloud `json:"tencentCloud"`
}

type Aliyun struct {
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	BucketName      string `json:"bucketName"`
}

type TencentCloud struct {
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
	Region     string `json:"region"`
	BucketName string `json:"bucketName"`
}
