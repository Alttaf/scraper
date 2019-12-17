package main

import (
	"fmt"

	"github.com/Alttaf/scraper/internal/handler"
	"github.com/Alttaf/scraper/internal/scraper"
)

func main() {
	fmt.Println("hello world")
	handler.Init()
	scraper.Scrape()
	handler.Run()
}
