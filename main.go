// package main

// import (
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gocolly/colly/v2"
// )

// // OgData represents the structure of the Open Graph data
// type OgData struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Image       string `json:"image"`
// 	OriginalURL string `json:"original_url"`
// }

// // getOgData handles the request and scrapes the Open Graph data
// func getOgData(context *gin.Context) {
// 	// Ensure that only GET requests are allowed
// 	if context.Request.Method != http.MethodGet {
// 		context.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
// 		return
// 	}

// 	// Get URL parameter
// 	urlParam := context.Query("url")

// 	// Check if URL parameter is empty
// 	if urlParam == "" {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "No URL parameter found"})
// 		return
// 	}

// 	// Split the URL parameter into multiple URLs if it's a comma-separated list
// 	urls := strings.Split(urlParam, ",")

// 	var ogDataList []OgData

// 	// Initialize a new Colly collector
// 	c := colly.NewCollector()

// 	// Define a function to handle the extraction and scraping logic for each URL
// 	handleURL := func(urlStr string) OgData {
// 		// Trim whitespace
// 		urlStr = strings.TrimSpace(urlStr)

// 		// Define a variable to hold Open Graph data for the current URL
// 		var ogData OgData
// 		ogData.OriginalURL = urlStr

// 		// Validate the URL
// 		parsedURL, err := url.ParseRequestURI(urlStr)
// 		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
// 			ogData.Title = "not found"
// 			ogData.Description = "not found"
// 			ogData.Image = "null"
// 			return ogData
// 		}

// 		// Initialize default values
// 		ogData.Title = "not found"
// 		ogData.Description = "not found"
// 		ogData.Image = "not found"

// 		// Set up Colly callbacks to extract Open Graph data
// 		c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
// 			if content := e.Attr("content"); content != "" {
// 				ogData.Title = content
// 			}
// 		})

// 		c.OnHTML("meta[property='og:description']", func(e *colly.HTMLElement) {
// 			if content := e.Attr("content"); content != "" {
// 				ogData.Description = content
// 			}
// 		})

// 		c.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
// 			if content := e.Attr("content"); content != "" {
// 				ogData.Image = content
// 			}
// 		})

// 		// Scrape the URL
// 		err = c.Visit(urlStr)
// 		if err != nil {
// 			ogData.Title = "not found"
// 			ogData.Description = "not found"
// 			ogData.Image = "null"
// 		}

// 		return ogData
// 	}

// 	// Process each URL
// 	for _, urlStr := range urls {
// 		ogData := handleURL(urlStr)
// 		ogDataList = append(ogDataList, ogData)
// 	}

// 	// Return the Open Graph data as a JSON response
// 	context.JSON(http.StatusOK, ogDataList)
// }

// func main() {
// 	r := gin.Default()
// 	// Allow all proxies (not recommended for production)
// 	r.SetTrustedProxies([]string{"*"}) // Allow all proxies (not secure)
// 	// r.SetTrustedProxies([]string{"127.0.0.1"}) // Adjust this as needed

// 	r.GET("/og", getOgData)
// 	r.Run(":8080")
// }

package main

import (
	"go_og_service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    // Allow all proxies (not recommended for production)
    r.SetTrustedProxies([]string{"*"}) // Allow all proxies (not secure)
    // r.SetTrustedProxies([]string{"127.0.0.1"}) // Adjust this as needed

    // Define routes
    r.GET("/og", handler.GetOgData)

    // Start the server
    r.Run(":8080")
}