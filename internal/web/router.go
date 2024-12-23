package web

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"cherry-markdown-webview/public/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func CorsInterceptor(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//resolve the cross origin[解决预请求]
		//w3c规范要求，当浏览器判定请求为复杂请求时，会在真实携带数据发送请求前，多一个预处理请求：
		//1. 请求方法不是get head post
		//2. post 的content-type不是application/x-www-form-urlencode,multipart/form-data,text/plain [也就是把content-type设置成"application/json"]
		//3. 请求设置了自定义的header字段: 比如业务需求，传一个字段，方面后端获取，不需要每个接口都传
		if r.Method == "OPTIONS" {
			//handle the preflight request
			w.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("127.0.0.1:%d", config.GetConfig().Web.GetPort()))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept,yi-token")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,yi-token")
		http.HandlerFunc(next).ServeHTTP(w, r)
	}
}

func registerRouter(mux *http.ServeMux) {
	defer func() {
		if err := recover(); err != nil {
			logs.Logger.Error(fmt.Sprintf("%w", err))
			logs.Logger.Error(fmt.Sprintf("%v", os.Stderr))
		}
	}()

	mux.HandleFunc("/file", CorsInterceptor(func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()

		uriValue := queryParams.Get("uri")
		var fileUri string
		if uriValue != "" {
			parsedURI, err := utils.ParsedURI(uriValue)
			if err != nil {
				logs.Logger.Error(err.Error())
				w.Write(make([]byte, 0))
				return
			}
			if len(parsedURI.Path) > 1 && parsedURI.Path[1:] != "" {
				fileUri = parsedURI.Path[1:]
			} else {
				w.Write(make([]byte, 0))
				return
			}

		}
		absPath := queryParams.Get("absPath")
		docPath := queryParams.Get("docPath")
		if absPath != "" && docPath != "" {
			fileUri = file.JoinAbsPath(docPath, absPath)
		}
		fileUri = filepath.Clean(fileUri)

		localFile, err := file.ReadFile(fileUri)
		if err != nil {
			logs.Logger.Error(err.Error())
			w.Write(make([]byte, 0))
			return
		}

		w.Header().Set("Content-Type", localFile.Mime)
		w.Write(localFile.Bytes)

	}))
}
