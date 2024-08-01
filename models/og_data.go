package models

// OgData represents the structure of the Open Graph data
type OgData struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Image       string `json:"image"`
    OriginalURL string `json:"original_url"`
}