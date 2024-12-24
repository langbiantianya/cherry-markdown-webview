package picbed

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"bytes"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/wailsapp/mimetype"
)

type AliOSS struct {
	*oss.Client
}

var ossClient *AliOSS

func NewOSSClient() {
	var (
		// 以从环境变量中获取访问凭证为例
		provider credentials.CredentialsProvider = credentials.NewStaticCredentialsProvider(
			config.GetConfig().PicBed.OSS.AccessKeyID,
			config.GetConfig().PicBed.OSS.AccessKeySecret,
		)
	)

	// 创建OSS客户端配置
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(config.GetConfig().PicBed.OSS.Region)

	// 创建OSS客户端
	client := oss.NewClient(cfg)
	ossClient = &AliOSS{client}
}

func (aliOSS AliOSS) Upload(sourceFile file.File) (*url.URL, error) {
	var reader *bytes.Buffer
	contentType := ""
	if sourceFile.StrData != "" && len(sourceFile.Bytes) == 0 {
		reader = bytes.NewBufferString(sourceFile.StrData)
		mtype, err := mimetype.DetectReader(reader)
		if err == nil {
			contentType = mtype.String()
		}
		reader.Reset()
		reader.WriteString(sourceFile.StrData)
	} else {
		reader = bytes.NewBuffer(sourceFile.Bytes)
		mtype, err := mimetype.DetectReader(reader)
		if err == nil {
			contentType = mtype.String()
		}
		reader.Reset()
		reader.Write(sourceFile.Bytes)
	}

	fileKey := strings.ReplaceAll(filepath.Clean(filepath.Join(config.GetConfig().PicBed.BasePath, sourceFile.Name)), "\\", "/")
	request := &oss.PutObjectRequest{
		Bucket:      oss.Ptr(config.GetConfig().PicBed.OSS.BucketName), // 存储空间名称
		Key:         oss.Ptr(fileKey[1:]),                              // 对象名称
		Body:        reader,
		ContentType: &contentType,
	}
	// 发送上传对象的请求
	result, err := aliOSS.PutObject(context.Background(), request)
	if err != nil {
		return nil, err
	}
	// 打印上传结果
	logs.Logger.Info(fmt.Sprintf("Upload successfully: %v", result))

	u, err := url.Parse(fmt.Sprintf("https://%s.%s.aliyuncs.com/%s", config.GetConfig().PicBed.OSS.BucketName, config.GetConfig().PicBed.OSS.Region, fileKey[1:]))
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (aliOSS AliOSS) Check(sourceFileName string) (bool, error) {
	fileKey := strings.ReplaceAll(filepath.Clean(filepath.Join(config.GetConfig().PicBed.BasePath, sourceFileName)), "\\", "/")

	return aliOSS.IsObjectExist(context.Background(), config.GetConfig().PicBed.OSS.BucketName, fileKey[1:])
}
