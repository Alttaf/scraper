package handler

import (
	"fmt"
	"net/http"
)

func scrapeSite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/health called")
	w.Write([]byte(`{"scrape":"ok"}`))
}
