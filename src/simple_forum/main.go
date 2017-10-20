package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))


	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	server := &http.Server {
		Addr:            "0.0.0.0:8080",
		Handler:         mux,
		ReadTimeout:     time.Duration(10 * int64(time.Second)),
		WriteTimeout:    time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes:  1 << 20,
	}

	server.ListenAndServe()
}
