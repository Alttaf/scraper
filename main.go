package main

import (
	"github.com/Alttaf/scraper/internal/handler"
)

func main() {
	// TODO add config through this config obj
	c := handler.Config{}
	server := handler.New(&c)
	handler.Run(server)
}
