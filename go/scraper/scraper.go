package scraper

import (
	"bytes"
	"fmt"
	"go_og_service/models"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func isValidURL(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func HandleURL(urlStr string) models.OgData {
	// urlStr = strings.TrimSpace(urlStr)

	var ogData models.OgData

	if isValidURL(urlStr) {
		ogData.OriginalURL = urlStr
		} else {
			log.Printf("Invalid URL %s", urlStr)
			ogData.Title = "not valid url"
			ogData.Description = "not valid url"
			ogData.Image = "not valid url"
			ogData.Icon = "not valid url"
			ogData.OriginalURL = urlStr
		return ogData
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		log.Printf("Invalid URL: %s, Error: %v", urlStr, err)
		ogData.Title = "not found"
		ogData.Description = "not found"
		ogData.Image = "not found"
		ogData.Icon = "not found"
		return ogData
	}

	c := colly.NewCollector(
		colly.UserAgent(
			"facebookexternalhit/1.1"))
		// Define the OnHTML callback for handling the HTML content
	c.OnHTML("head", func(e *colly.HTMLElement) {
		// Convert []byte to io.Reader
		bodyReader := bytes.NewReader(e.Response.Body)

		// Create a GoQuery document from the HTML content
		doc, err := goquery.NewDocumentFromReader(bodyReader)
		if err != nil {
			log.Fatal(err)
		}

		// Iterate over meta tags and print their attributes
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			name, existsName := s.Attr("name")
			property, existsProperty := s.Attr("property")
			content, _ := s.Attr("content")

			if existsName {
				fmt.Printf("by goquery: Meta Name: %s, Content: %s\n", name, content)
			}
			if existsProperty {
				fmt.Printf("by goquery: Meta Property: %s, Content: %s\n", property, content)
			}
		})
	})
		c.OnHTML("meta", func(e *colly.HTMLElement) {
		name := e.Attr("name")
		property := e.Attr("property")
		content := e.Attr("content")
		if name != "" {
			fmt.Printf("by coly: Meta Name: %s, Content: %s\n", name, content)
		}
		if property != "" {
			fmt.Printf("by coly: Meta Property: %s, Content: %s\n", property, content)
		}
	})
	// Set up the collector to only parse the <head> tag
	c.OnHTML("head", func(e *colly.HTMLElement) {
		
		// Extract Open Graph meta tags
		e.ForEach("meta[property='og:title']", func(_ int, el *colly.HTMLElement) {
			if content := el.Attr("content"); content != "" {
				ogData.Title = content
			}
		})
		e.ForEach("meta[property='og:description']", func(_ int, el *colly.HTMLElement) {
			if content := el.Attr("content"); content != "" {
				ogData.Description = content
			}
		})
		e.ForEach("meta[property='og:image']", func(_ int, el *colly.HTMLElement) {
			if content := el.Attr("content"); content != "" {
				ogData.Image = content
			}
		})
		e.ForEach("meta[property='twitter:image']", func(_ int, el *colly.HTMLElement) {
			if ogData.Image == "" {
				if content := el.Attr("content"); content != "" {
					ogData.Image = content
				}
			}
		})
		e.ForEach("meta[property='twitter:description']", func(_ int, el *colly.HTMLElement) {
			if ogData.Description == "" {
				if content := el.Attr("content"); content != "" {
					ogData.Description = content
				}
			}
		})
		e.ForEach("meta[property='twitter:title']", func(_ int, el *colly.HTMLElement) {
			if ogData.Title == "" {
				ogData.Title = el.Text
			}
		})

		// Extract favicon links
		e.ForEach("link[rel='icon']", func(_ int, el *colly.HTMLElement) {
			icon := el.Attr("href")
			if icon != "" {
				if !strings.HasPrefix(icon, "http") {
					if baseURL, err := e.Request.URL.Parse(icon); err == nil {
						icon = baseURL.String()
					}
				}
				ogData.Icon = icon
			}
		})
		e.ForEach("link[rel='shortcut icon']", func(_ int, el *colly.HTMLElement) {
			icon := el.Attr("href")
			if ogData.Icon == "" && icon != "" {
				if !strings.HasPrefix(icon, "http") {
					if baseURL, err := e.Request.URL.Parse(icon); err == nil {
						icon = baseURL.String()
					}
				}
				ogData.Icon = icon
			}
		})
		e.ForEach("link[href*='favicon']", func(_ int, el *colly.HTMLElement) {
			icon := el.Attr("href")
			if ogData.Icon == "" && icon != "" {
				if !strings.HasPrefix(icon, "http") {
					if baseURL, err := e.Request.URL.Parse(icon); err == nil {
						icon = baseURL.String()
					}
				}
				ogData.Icon = icon
			}
		})
	})

	// Setup error handling and visit URL
	err = c.Visit(urlStr)
	if err != nil {
		log.Printf("Error visiting URL %s: %v", urlStr, err)
		ogData.Title = "not found"
		ogData.Description = "not found"
		ogData.Image = "not found"
		ogData.Icon = "not found"
	}

	log.Printf("Scraped Open Graph data: %+v", ogData)

	return ogData
}

