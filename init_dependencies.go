package main

import (
	"github.com/lucasschilin/crud-go/src/controller"
	"github.com/lucasschilin/crud-go/src/model/repository"
	"github.com/lucasschilin/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepository)
	userController := controller.NewUserControllerInterface(userService)

	return userController
}
