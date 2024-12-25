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
		return NewCOSClient()
	case config.OSS:
		return NewOSSClient()
	default:
		return nil
	}
}
