package handler

import (
	"fmt"
	"net/http"
	"time"
)

var (
	server *http.Server
)

//Init - main init function to create server
func Init() {

	r := http.NewServeMux()
	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/health called")
		w.Write([]byte(`{"health":"ok"}`))
	})
	r.HandleFunc("/scrape", scrapeSite)

	server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", "127.0.0.1", 9999),
		Handler:           r,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		ReadTimeout:       60 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
