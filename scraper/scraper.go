package scraper

import (
	"go_og_service/models"
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
        ogData.Title = "not found"
        ogData.Description = "not found"
        ogData.Image = "not found"
        ogData.Icon = "not found"
        return ogData
    }

    // Initialize a new Colly collector
    c := colly.NewCollector()

    // Set up Colly callbacks to extract Open Graph data
    c.OnHTML("link[rel='icon']", func(e *colly.HTMLElement) {
        if icon := e.Attr("href"); icon != "" {
            ogData.Icon = icon
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
        ogData.Title = "not found"
        ogData.Description = "not found"
        ogData.Image = "not found"
    }

    return ogData
}