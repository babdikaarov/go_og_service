package scraper

import (
	"go_og_service/models"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

// HandleURL scrapes Open Graph data from a URL
func HandleURL(urlStr string) models.OgData {
	// Trim whitespace
	urlStr = strings.TrimSpace(urlStr)

	// Define a variable to hold Open Graph data for the current URL
	var ogData models.OgData
	ogData.OriginalURL = urlStr

	// Validate the URL
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		log.Printf("Invalid URL: %s, Error: %v", urlStr, err)
		ogData.Title = "not found"
		ogData.Description = "not found"
		ogData.Image = "not found"
		ogData.Icon = "not found"
		return ogData
	}

	// Initialize a new Colly collector
	c := colly.NewCollector()

	// Log requests and responses
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Requesting URL: %s", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Printf("Received response with status: %d for URL: %s", r.StatusCode, r.Request.URL)
	})

	c.OnHTML("link[rel='icon']", func(e *colly.HTMLElement) {
		icon := e.Attr("href")

		// Log icon details
		log.Printf("Found icon href: %s", icon)

		// Check if the href is a full URL
		if icon != "" && (len(icon) >= 4 && icon[:4] == "http") {
			ogData.Icon = icon
		} else {
			// Handle relative paths by prepending the base URL
			if baseURL, err := e.Request.URL.Parse(icon); err == nil {
				ogData.Icon = baseURL.String()
			}
		}
	})

	c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
		if content := e.Attr("content"); content != "" {
			ogData.Title = content
		}
	})

	c.OnHTML("meta[property='og:description']", func(e *colly.HTMLElement) {
		if content := e.Attr("content"); content != "" {
			ogData.Description = content
		}
	})

	c.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
		if content := e.Attr("content"); content != "" {
			ogData.Image = content
		}
	})

	// Scrape the URL
	err = c.Visit(urlStr)
	if err != nil {
		log.Printf("Error visiting URL %s: %v", urlStr, err)
		ogData.Title = "not found"
		ogData.Description = "not found"
		ogData.Image = "not found"
	}

	// Log the final Open Graph data
	log.Printf("Scraped Open Graph data: %+v", ogData)

	return ogData
}