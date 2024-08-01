package handler

import (
    "net/http"
    "strings"
    "go_og_service/scraper"
    "go_og_service/models"
    "github.com/gin-gonic/gin"
)

// GetOgData handles the request and scrapes the Open Graph data
func GetOgData(context *gin.Context) {
    // Ensure that only GET requests are allowed
    if context.Request.Method != http.MethodGet {
        context.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
        return
    }

    // Get URL parameter
    urlParam := context.Query("url")

    // Check if URL parameter is empty
    if urlParam == "" {
        context.JSON(http.StatusNotFound, gin.H{"error": "No URL parameter found"})
        return
    }

    // Split the URL parameter into multiple URLs if it's a comma-separated list
    urls := strings.Split(urlParam, ",")

    var ogDataList []models.OgData

    // Process each URL
    for _, urlStr := range urls {
        ogData := scraper.HandleURL(urlStr)
        ogDataList = append(ogDataList, ogData)
    }

    // Return the Open Graph data as a JSON response
    context.JSON(http.StatusOK, ogDataList)
}