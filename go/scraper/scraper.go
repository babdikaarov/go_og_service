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

	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"))

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

// package scraper

// import (
// 	"go_og_service/models"
// 	"log"
// 	"net/url"
// 	"strings"

// 	"github.com/gocolly/colly/v2"
// )

// func HandleURL(urlStr string) models.OgData {
// 	urlStr = strings.TrimSpace(urlStr)

// 	var ogData models.OgData
// 	ogData.OriginalURL = urlStr

// 	parsedURL, err := url.ParseRequestURI(urlStr)
// 	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
// 		log.Printf("Invalid URL: %s, Error: %v", urlStr, err)
// 		ogData.Title = "not found"
// 		ogData.Description = "not found"
// 		ogData.Image = "not found"
// 		ogData.Icon = "not found"
// 		return ogData
// 	}

// 	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"))
	
// 	c.OnRequest(func(r *colly.Request) {
// 		log.Printf("Requesting URL: %s", r.URL.String())
// 		log.Println("Request Headers:", r.Headers)
// 		log.Println("Request User-Agent:", r.Headers.Get("User-Agent"))
// 		for key, value := range *r.Headers {
// 			log.Printf("Request Header: %s: %s", key, value)
// 		}
// 	})


// 	c.OnResponse(func(r *colly.Response) {
// 		log.Printf("Received response with status: %d for URL: %s", r.StatusCode, r.Request.URL)
// 		for key, value := range *r.Headers {
// 			log.Printf("Response Header: %s: %s", key, value)
// 		}
// 	})

// 	c.OnHTML("link[rel='icon']", func(e *colly.HTMLElement) {
// 		icon := e.Attr("href")
// 		log.Printf("Found icon href (rel='icon'): %s", icon)
// 		if icon != "" {
// 			if !strings.HasPrefix(icon, "http") {
// 				if baseURL, err := e.Request.URL.Parse(icon); err == nil {
// 					icon = baseURL.String()
// 				}
// 			}
// 			ogData.Icon = icon
// 		}
// 	})

// 	c.OnHTML("link[rel='shortcut icon']", func(e *colly.HTMLElement) {
// 		icon := e.Attr("href")
// 		if ogData.Icon == "" && icon != "" {
// 			log.Printf("Found icon href (rel='shortcut icon'): %s", icon)
// 			if !strings.HasPrefix(icon, "http") {
// 				if baseURL, err := e.Request.URL.Parse(icon); err == nil {
// 					icon = baseURL.String()
// 				}
// 			}
// 			ogData.Icon = icon
// 		}
// 	})

// 	c.OnHTML("link[href*='favicon']", func(e *colly.HTMLElement) {
// 		icon := e.Attr("href")
// 		if ogData.Icon == "" && icon != "" {
// 			log.Printf("Found favicon href: %s", icon)
// 			if !strings.HasPrefix(icon, "http") {
// 				if baseURL, err := e.Request.URL.Parse(icon); err == nil {
// 					icon = baseURL.String()
// 				}
// 			}
// 			ogData.Icon = icon
// 		}
// 	})

// 	c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
// 		if content := e.Attr("content"); content != "" {
// 			ogData.Title = content
// 		}
// 	})

// 	c.OnHTML("meta[property='og:description']", func(e *colly.HTMLElement) {
// 		if content := e.Attr("content"); content != "" {
// 			ogData.Description = content
// 		}
// 	})

// 	c.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
// 		if content := e.Attr("content"); content != "" {
// 			ogData.Image = content
// 		}
// 	})
// 	c.OnHTML("meta[property='twitter:image']", func(e *colly.HTMLElement) {
// 		 if ogData.Image == "" { 
// 			if content := e.Attr("content"); content != "" {
// 				ogData.Image = content
// 			}
// 		}
// 	})

// 	c.OnHTML("meta[property='twitter:description']", func(e *colly.HTMLElement) {
//         if ogData.Description == "" { // Only set if not already set
//             if content := e.Attr("content"); content != "" {
//                 ogData.Description = content
//             }
//         }
//     })

//     c.OnHTML("meta[property='twitter:title']", func(e *colly.HTMLElement) {
//         if ogData.Title == "" { // Only set if not already set
//             ogData.Title = e.Text
//         }
//     })

// 	err = c.Visit(urlStr)
// 	if err != nil {
// 		log.Printf("Error visiting URL %s: %v", urlStr, err)
// 		ogData.Title = "not found"
// 		ogData.Description = "not found"
// 		ogData.Image = "not found"
// 	}

// 	log.Printf("Scraped Open Graph data: %+v", ogData)

// 	return ogData
// }