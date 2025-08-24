package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/configuration/rest_err"
	"github.com/lucasschilin/crud-go/src/controller/model/request"
)

func PostUser(c *gin.Context) {
	var userRequest request.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect fields, error: %s", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
