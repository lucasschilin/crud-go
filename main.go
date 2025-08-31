package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasschilin/crud-go/src/configuration/database/mongodb"
	"github.com/lucasschilin/crud-go/src/controller/routes"
)

const (
	GIN_MODE = gin.EnvGinMode
	API_PORT = "API_PORT"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("error on loading .env file: %v", err))
	}

	database, err := mongodb.NewConnection(context.Background())
	if err != nil {
		log.Fatalf("error to connect to database: %v", err)
	}

	userController := initDependencies(database)

	gin.SetMode(os.Getenv(GIN_MODE))
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":" + os.Getenv(API_PORT)); err != nil {
		log.Fatal(err)
	}
}
