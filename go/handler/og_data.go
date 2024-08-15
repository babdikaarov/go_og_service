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

// GenerateOgData handles the request and scrapes the Open Graph data
func GenerateOgData(context *gin.Context) {
	// Ensure that only GET requests are allowed
	if context.Request.Method != http.MethodGet {
		context.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Get URL parameter
	urlParam := context.Query("url")
	folderParam := context.Query("filename")

	// Check if URL parameter is empty
	if urlParam == "" {
		context.JSON(http.StatusNotFound, gin.H{"error": "No URL parameter found"})
		return
	}

	// Format the folderParam: convert to lowercase and replace spaces with underscores
	formattedFolderParam := strings.ToLower(folderParam)
	formattedFolderParam = strings.ReplaceAll(formattedFolderParam, " ", "_")

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

	// Create a file in the ZIP archive with formatted folderParam
	zipFileName := "data.json" // Default file name
	if formattedFolderParam != "" {
		zipFileName = formattedFolderParam + ".json" // Use formatted folderParam for file name
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

	// Set the default ZIP file name
	filename := "ogdata.zip"
	if formattedFolderParam != "" {
		filename = formattedFolderParam + ".zip"
	}

	// Set the Content-Disposition header and content type
	context.Header("Content-Type", "application/zip")
	context.Header("Content-Disposition", "attachment; filename="+filename)

	// Write the ZIP file to the response
	context.Data(http.StatusOK, "application/zip", buf.Bytes())
}
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

	// Return the Open Graph data as JSON
	context.JSON(http.StatusOK, gin.H{
		"data": ogDataList,
	})
}