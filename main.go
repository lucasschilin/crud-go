package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasschilin/crud-go/src/controller"
	"github.com/lucasschilin/crud-go/src/controller/routes"
	"github.com/lucasschilin/crud-go/src/model/service"
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

	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	gin.SetMode(os.Getenv(GIN_MODE))
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":" + os.Getenv(API_PORT)); err != nil {
		log.Fatal(err)
	}
}
