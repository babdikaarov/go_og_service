package main

import (
	"go_og_service/handler"
	"log"
	"os"
	"time"

	_ "go_og_service/docs" // Import the generated docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Open graph data generator
// @version 1.0
// @Scrap any Open Graph data .
// @contact.name Beksultan Abdikaarov
// @contact.email babdikaarov@gmail.com
// @contact.url https://github.com/babdikaarov
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @host babdikaarov.home.kg
// @BasePath /api/v1
func main() {


    gin.SetMode(gin.ReleaseMode) // production mode
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading .env file setting Default port :3030")
    }
   
    port := os.Getenv("PORT")
    if port == "" {
        port = ":3030" // Default port
    }
    r := gin.Default()
    r.Static("/static", "./static")
    r.StaticFile("/favicon.ico", "./static/favicon.ico")
    r.LoadHTMLGlob("templates/*")
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"GET", "POST"}
    config.AllowHeaders = []string{"Origin", "Content-Type"}
    config.ExposeHeaders = []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods"}
    config.MaxAge = 2 * time.Hour

    r.Use(cors.New(config))

    r.SetTrustedProxies([]string{"*"})
    r.GET("/", handler.ServeForm)          // Serve the form
    // Swagger endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.POST("/generate", handler.HandleFormSubmission) // Handle form submission
    r.GET("/api/v1/zip", handler.GenerateOgData)
    r.GET("/api/v1/json", handler.GetOgData)
    r.Run(port)
   
}