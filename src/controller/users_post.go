package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/configuration/validation"
	"github.com/lucasschilin/crud-go/src/controller/model/request"
	"github.com/lucasschilin/crud-go/src/controller/model/response"
)

func PostUser(c *gin.Context) {
	var userRequest request.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Print(err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, &response.User{
		ID:    "000-000-000",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	})
}
