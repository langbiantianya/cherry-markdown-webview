package web

import (
	"cherry-markdown-webview/internal/config"
	"cherry-markdown-webview/internal/logs"
	"fmt"
	"net"

	"net/http"
)

func StartServer() {

	port := 8100
	// 尝试监听端口
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))

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
	
	registerRouter(mux)

	logs.Logger.Info(fmt.Sprintf("127.0.0.1:%d", port))
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux)
		if err != nil {
			logs.Logger.Fatal("Error starting server:" + err.Error())
		}
	}()

}
