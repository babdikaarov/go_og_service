package scraper

import (
	"go_og_service/models"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

func HandleURL(urlStr string) models.OgData {
	urlStr = strings.TrimSpace(urlStr)

	var ogData models.OgData
	ogData.OriginalURL = urlStr

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		log.Printf("Invalid URL: %s, Error: %v", urlStr, err)
		ogData.Title = "not found"
		ogData.Description = "not found"
		ogData.Image = "not found"
		ogData.Icon = "not found"
		return ogData
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Requesting URL: %s", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Printf("Received response with status: %d for URL: %s", r.StatusCode, r.Request.URL)
	})

	c.OnHTML("link[rel='icon']", func(e *colly.HTMLElement) {
		icon := e.Attr("href")
		log.Printf("Found icon href (rel='icon'): %s", icon)
		if icon != "" {
			if !strings.HasPrefix(icon, "http") {
				if baseURL, err := e.Request.URL.Parse(icon); err == nil {
					icon = baseURL.String()
				}
			}
			ogData.Icon = icon
		}
	})

	c.OnHTML("link[rel='shortcut icon']", func(e *colly.HTMLElement) {
		icon := e.Attr("href")
		if ogData.Icon == "" && icon != "" {
			log.Printf("Found icon href (rel='shortcut icon'): %s", icon)
			if !strings.HasPrefix(icon, "http") {
				if baseURL, err := e.Request.URL.Parse(icon); err == nil {
					icon = baseURL.String()
				}
			}
			ogData.Icon = icon
		}
	})

	c.OnHTML("link[href*='favicon']", func(e *colly.HTMLElement) {
		icon := e.Attr("href")
		if ogData.Icon == "" && icon != "" {
			log.Printf("Found favicon href: %s", icon)
			if !strings.HasPrefix(icon, "http") {
				if baseURL, err := e.Request.URL.Parse(icon); err == nil {
					icon = baseURL.String()
				}
			}
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

	c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
        if ogData.Description == "" { // Only set if not already set
            if content := e.Attr("content"); content != "" {
                ogData.Description = content
            }
        }
    })

    c.OnHTML("title", func(e *colly.HTMLElement) {
        if ogData.Title == "" { // Only set if not already set
            ogData.Title = e.Text
        }
    })

	err = c.Visit(urlStr)
	if err != nil {
		log.Printf("Error visiting URL %s: %v", urlStr, err)
		ogData.Title = "not found"
		ogData.Description = "not found"
		ogData.Image = "not found"
	}

	log.Printf("Scraped Open Graph data: %+v", ogData)

	return ogData
}