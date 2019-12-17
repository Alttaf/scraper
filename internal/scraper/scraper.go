package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Scrape - begin scraping
func Scrape() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("https://blog.golang.org/"),

		colly.Async(true),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link Found : %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// On every a element which has href attribute call callback
	c.OnHTML(".blog-post-single-title", func(e *colly.HTMLElement) {
		// Print link
		fmt.Printf("post found: %v ->\n", e)
		fmt.Printf("post found: %q ->\n", e.Text)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://blog.golang.org/")
	fmt.Println("scrape")
}
