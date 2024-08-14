package main

import (
	"go_og_service/handler"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    gin.SetMode(gin.ReleaseMode) // production mode
     if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }
   
    port := os.Getenv("PORT")
    if port == "" {
        port = ":3030" // Default port
    }
    r := gin.Default()
	r.LoadHTMLGlob("*.html")
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"GET", "POST"}
    config.AllowHeaders = []string{"Origin", "Content-Type"}
    config.ExposeHeaders = []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods"}
    config.MaxAge = 2 * time.Hour

    r.Use(cors.New(config))

    r.SetTrustedProxies([]string{"*"})

    r.GET("/", handler.ServeForm)          // Serve the form
    r.POST("/generate", handler.HandleFormSubmission) // Handle form submission
    r.GET("/og", handler.GetOgData)
    r.Run(port)
   
}