package main

import (
	widget "go-stream-widget/api"
	"net/http"
	"os"
)

func main() {

	// これで静的ファイルにアクセスできるとおもったのになあ
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/widget/", http.StripPrefix("/widget/", fs))

	http.HandleFunc("/ppppCounter", widget.TskCounter)
	http.HandleFunc("/view", widget.Handler)
	http.HandleFunc("/", widget.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}

	// 起動する
	http.ListenAndServe(":"+port, nil/)
	// http.ListenAndServe("localhost:"+port, nil)
}
