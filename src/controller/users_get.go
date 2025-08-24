package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/controller/model/response"
)

func GetUser(c *gin.Context) {

	c.JSON(http.StatusOK, response.User{
		ID:    "6d5s4",
		Email: "lucasschilin@gmail.com",
		Name:  "Lucas Schilin",
		Age:   24,
	})
}
