package main

import (
	"go_og_service/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods"}
	config.MaxAge = 2 * time.Hour

	r.Use(cors.New(config))

	r.SetTrustedProxies([]string{"*"})

	r.GET("/og", handler.GetOgData)
	r.Run(":8080")
}