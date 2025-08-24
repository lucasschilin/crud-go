package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasschilin/crud-go/src/controller/routes"
)

const (
	GIN_MODE = gin.EnvGinMode
	API_PORT = "API_PORT"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(os.Getenv(GIN_MODE))
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":" + os.Getenv(API_PORT)); err != nil {
		log.Fatal(err)
	}
}
