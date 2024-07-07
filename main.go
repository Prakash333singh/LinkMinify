package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Prakash333singh/url_shotner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	// Setup routes for the URL shortener application
	setupRoutes(router)

	// Get the port from environment variable or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

// CORSMiddleware sets up the CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// setupRoutes sets up the routes for the application
func setupRoutes(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.DELETE("/api/v1/:shortID", routes.DeleteURL)
	router.PUT("/api/v1/:shortID", routes.EditURl)
	router.POST("/api/v1/addTag", routes.AddTag)
	router.GET("/hello", myGetFunction)
}

// simpleMessage struct for testing endpoint
type simpleMessage struct {
	Hello   string `json:"hello"`
	Message string `json:"message"`
}

// myGetFunction is a test endpoint to verify the server is running
func myGetFunction(c *gin.Context) {
	simpleMessage := simpleMessage{
		Hello:   "World!",
		Message: "how are you doingg!!!",
	}

	c.IndentedJSON(http.StatusOK, simpleMessage)
}
