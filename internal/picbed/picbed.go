package picbed

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"net/url"
)

type Picbed interface {
	Upload(file.File) (*url.URL, error)
	Check(string) (bool, error)
}

func GetPicbed() Picbed {
	switch config.GetConfig().PicBed.Activated {
	case config.COS:
		if cosClient == nil {
			NewCOSClient()
		}
		return cosClient
	case config.OSS:
		if ossClient == nil {
			NewOSSClient()
		}
		return ossClient
	default:
		return nil
	}
}
