package handler

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"go_og_service/models"
	"go_og_service/scraper"
	"net/http"
	"strings"

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
    folderParam := context.Query("name")

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

    // Convert the Open Graph data to JSON
    jsonData, err := json.Marshal(ogDataList)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JSON"})
        return
    }

    // Create a buffer to write the ZIP file
    var buf bytes.Buffer
    zipWriter := zip.NewWriter(&buf)

    // Create a file in the ZIP archive
    zipFileName := "data.json" // Default file name
    if folderParam != "" {
        zipFileName = folderParam + ".json" // Use folderParam for file name
    }
    zipFile, err := zipWriter.Create(zipFileName)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ZIP file"})
        return
    }

    // Write JSON data to the file in the ZIP archive
    _, err = zipFile.Write(jsonData)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write JSON data to ZIP file"})
        return
    }

    // Close the ZIP writer
    err = zipWriter.Close()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close ZIP file"})
        return
    }

    // Set the content type to ZIP and trigger download
    context.Header("Content-Type", "application/zip")
    context.Header("Content-Disposition", "attachment; filename=ogdata.zip")

    // Write the ZIP file to the response
    context.Data(http.StatusOK, "application/zip", buf.Bytes())
}