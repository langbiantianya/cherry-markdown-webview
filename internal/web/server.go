package web

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"cherry-markdown-webview/public/utils"
	"fmt"
	"net"

	"net/http"
)

func StartServer() {

	port := 8100
	// 尝试监听端口
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	for countTry := 50; err != nil && countTry > 0; countTry-- {
		port++
		ln, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	}

	if err != nil {
		logs.Logger.Fatal(err.Error())
	}

	err = ln.Close()

	if err != nil {
		logs.Logger.Fatal(err.Error())
	}

	conf := config.GetConfig()
	conf.Web.SetPort(port)

	mux := http.NewServeMux()

	handlerFile := func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()

		uriValue := queryParams.Get("uri")

		parsedURI, err := utils.ParsedURI(uriValue)
		if err != nil {
			logs.Logger.Error(err.Error())
			w.Write(make([]byte, 0))
		}
		localFile, err := file.ReadFile(parsedURI.Path[1:])
		if err != nil {
			logs.Logger.Error(err.Error())
			w.Write(make([]byte, 0))
		}

		w.Header().Set("Content-Type", localFile.Mime)
		w.Write(localFile.Bytes)

	}

	mux.HandleFunc("/file", handlerFile)

	logs.Logger.Info(fmt.Sprintf("127.0.0.1:%d", port))
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux)
		if err != nil {
			logs.Logger.Fatal("Error starting server:" + err.Error())
		}
	}()

}
