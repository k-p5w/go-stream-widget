package main

import (
	widget "go-stream-widget/api"
	"net/http"
	"os"
)

func main() {

	// これで静的ファイルにアクセスできるとおもったのになあ
	fs := http.FileServer(http.Dir("public"))
	// http.Handle("/tool/", http.StripPrefix("/tool/", fs))

	// `/tool/`以下のパスにアクセスされた場合、`fs`ハンドラーで処理
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/ppppCounter", widget.TskCounter)
	http.HandleFunc("/view", widget.Handler)
	http.HandleFunc("/", widget.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}

	// 起動する
	http.ListenAndServe(":"+port, nil)
	// http.ListenAndServe("localhost:"+port, nil)
}
