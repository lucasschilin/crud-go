package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/model/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	PostUser(c *gin.Context)
	GetUser(c *gin.Context)
	PutUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
