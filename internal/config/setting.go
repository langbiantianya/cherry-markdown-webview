package config

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
