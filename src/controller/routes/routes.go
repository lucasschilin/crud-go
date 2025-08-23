package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.PostUser)
	r.PUT("/users/:id", controller.PutUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
