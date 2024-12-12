package utils

import (
	"errors"
	"net/url"
	"strings"
)

func ParsedURI(uri string) (*url.URL, error) {
	// 去除可能存在的前后空格
	uri = strings.TrimSpace(uri)

	// 检查是否以 "file://" 开头
	if !strings.HasPrefix(uri, "file://") {
		return nil, errors.New("协议错误")
	}

	// 尝试解析URI
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// 检查路径是否为空
	if parsedURI.Path == "" {
		return nil, errors.New("文件路径为空")
	}

	// 可以添加更多的校验逻辑，例如检查路径是否合法等

	return parsedURI, nil
}
