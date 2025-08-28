package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasschilin/crud-go/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {

	r.GET("/users/:id", userController.GetUser)
	r.POST("/users", userController.PostUser)
	r.PUT("/users/:id", userController.PutUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
