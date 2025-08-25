package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"github.com/lucasschilin/crud-go/src/configuration/validation"
	"github.com/lucasschilin/crud-go/src/controller/model/request"
	"github.com/lucasschilin/crud-go/src/controller/model/response"
	"go.uber.org/zap"
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

	logger.Info("user 000-000-000 created successfully", journeyTag)
	c.JSON(http.StatusOK, &response.User{
		ID:    "000-000-000",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	})
}
