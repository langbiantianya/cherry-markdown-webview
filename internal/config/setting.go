package config

import "sync"

var conf *Config = &Config{
	Web: &Web{},
	OSS: &OSS{},
}

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
	// mu  sync.Mutex
	Web *Web
	OSS *OSS
}

type Web struct {
	mu   sync.Mutex
	port int
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
