package web

import (
	"cherry-markdown-webview/internal/file"
	"cherry-markdown-webview/internal/logs"
	"cherry-markdown-webview/public/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func registerRouter(mux *http.ServeMux) {
	defer func() {
		if err := recover(); err != nil {
			logs.Logger.Error(fmt.Sprintf("%w", err))
			logs.Logger.Error(fmt.Sprintf("%v", os.Stderr))
		}
	}()

	mux.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()

		uriValue := queryParams.Get("uri")
		var fileUri string
		if uriValue != "" {
			parsedURI, err := utils.ParsedURI(uriValue)
			if err != nil {
				logs.Logger.Error(err.Error())
				w.Write(make([]byte, 0))
			}
			fileUri = parsedURI.Path[1:]

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
		}

		w.Header().Set("Content-Type", localFile.Mime)
		w.Write(localFile.Bytes)

	})
}
