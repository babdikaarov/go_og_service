package handler

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// ServeForm serves the HTML form.
func ServeForm(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func HandleFormSubmission(c *gin.Context) {
	links := c.PostForm("links")
	filename := c.PostForm("filename")
	outputType := c.PostForm("outputType")

	// Clean up links and filename
	links = strings.TrimSpace(links)
	filename = strings.TrimSpace(filename)

	if links == "" || (outputType == "zip" && filename == "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Links and filename cannot be empty"})
		return
	}

	var ogEndpoint string
	if outputType == "zip" {
		// Construct the URL for the /zip endpoint
		ogEndpoint = "/api/v1/zip?url=" + url.QueryEscape(links) + "&filename=" + url.QueryEscape(filename)
	} else if outputType == "json" {
		// Construct the URL for the /json endpoint
		ogEndpoint = "/api/v1/json?url=" + url.QueryEscape(links)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid output type"})
		return
	}

	// Redirect to the appropriate endpoint
	c.Redirect(http.StatusSeeOther, ogEndpoint)
}