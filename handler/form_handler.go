package handler

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// ServeForm serves the HTML form.
func ServeForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

// HandleFormSubmission processes form submissions and redirects to the /og endpoint.
func HandleFormSubmission(c *gin.Context) {
	links := c.PostForm("links")
	filename := c.PostForm("filename")

	// Clean up links and filename
	links = strings.TrimSpace(links)
	filename = strings.TrimSpace(filename)

	if links == "" || filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Links and filename cannot be empty"})
		return
	}

	// Construct the URL for the /og endpoint
	ogEndpoint := "http://localhost:8080/og?url=" + url.QueryEscape(links) + "&filename=" + url.QueryEscape(filename)

	// Redirect to the /og endpoint to handle ZIP file creation
	c.Redirect(http.StatusSeeOther, ogEndpoint)
}