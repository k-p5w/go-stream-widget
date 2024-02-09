package main

import (
	widget "go-stream-widget/api"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/view", widget.Handler)
	http.HandleFunc("/", widget.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	// 起動する
	http.ListenAndServe(":"+port, nil)
}
