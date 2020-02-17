package handler

import (
	"net/http"
)

func scrapeSite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"scrape":"ok"}`))
}
