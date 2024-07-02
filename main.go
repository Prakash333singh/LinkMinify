package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Prakash333singh/url_shotner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() ///for exracting with env
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	setupRoutes(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))

}

func setupRoutes(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortId", routes.GetByShortID)
	router.DELETE("/api/v1/:shortId", routes.DeleteURL)
	router.PUT("/api/v1/:shortId", routes.EditURl)
	router.POST("/api/v1/addTag", routes.AddTag)
}
