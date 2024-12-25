package picbed

import (
	"bytes"
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"context"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/wailsapp/mimetype"
)

type TencentCOS struct {
	*cos.Client
}

func NewCOSClient() *TencentCOS {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(config.GetConfig().PicBed.COS.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.GetConfig().PicBed.COS.SecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: config.GetConfig().PicBed.COS.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return &TencentCOS{client}
}

func (tentcentCos TencentCOS) Upload(sourceFile file.File) (*url.URL, error) {
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

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
		},
	}
	fileKey := strings.ReplaceAll(filepath.Clean(filepath.Join(config.GetConfig().PicBed.BasePath, sourceFile.Name)), "\\", "/")
	_, err := tentcentCos.Object.Put(context.Background(), fileKey[1:], reader, opt)
	if err != nil {
		return nil, err
	}
	ourl := tentcentCos.Object.GetObjectURL(fileKey[1:])
	return ourl, nil
}

func (tentcentCos TencentCOS) Check(sourceFileName string) (bool, error) {
	fileKey := strings.ReplaceAll(filepath.Clean(filepath.Join(config.GetConfig().PicBed.BasePath, sourceFileName)), "\\", "/")
	return tentcentCos.Object.IsExist(context.Background(), fileKey[1:])
}
