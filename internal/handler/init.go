package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Config for server
type Config struct {
	dataStore string
}

func initializeRoutes(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/health called")
		w.Write([]byte(`{"health":"ok"}`))
	})
	r.HandleFunc("/scrape", scrapeSite)

}

//New - create a new insance of the handler
func New(c *Config) *http.Server {
	// you can use config to overrid any values
	r := mux.NewRouter()
	initializeRoutes(r)
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", "127.0.0.1", 9999),
		Handler:           r,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		ReadTimeout:       60 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
