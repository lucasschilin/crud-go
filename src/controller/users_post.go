package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/validation"
	"github.com/lucasschilin/crud-go/src/controller/model/request"
	"github.com/lucasschilin/crud-go/src/model"
	"github.com/lucasschilin/crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) PostUser(c *gin.Context) {
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

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("error trying to call CreateUser service", err, journeyTag)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userID", domain.GetID()),
		journeyTag,
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
