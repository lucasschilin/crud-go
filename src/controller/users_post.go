package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/controller/model/request"
	"github.com/lucasschilin/crud-go/src/controller/model/response"
)

func PostUser(c *gin.Context) {
	var userRequest request.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Println(err.Error())
		restErr := rest_err.NewBadRequestError("There are some incorrect fields")
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
