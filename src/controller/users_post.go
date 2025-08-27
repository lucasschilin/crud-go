package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/validation"
	"github.com/lucasschilin/crud-go/src/controller/model/request"
	"github.com/lucasschilin/crud-go/src/model"
	"github.com/lucasschilin/crud-go/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func PostUser(c *gin.Context) {
	journeyTag := zap.String("journey", "create-user")

	logger.Info("init PostUser() controller", journeyTag)
	var userRequest request.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate POST /users request", err, journeyTag)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	err := service.CreateUser(domain)
	if err != nil {
		logger.Error("error trying to validate POST /users request", err, journeyTag)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user 000-000-000 created successfully", journeyTag)
	c.JSON(http.StatusOK, domain)
}
