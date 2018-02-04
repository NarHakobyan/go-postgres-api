package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/http/controllers"
)

var UserRouter *gin.RouterGroup

func init() {
	UserRouter = ApiRouter.Group("/users")

	UserRouter.GET("/", controllers.UserController.GetUsers)

	UserRouter.GET("/:id", controllers.UserController.GetUser)

	UserRouter.POST("/", controllers.UserController.PostUser)
}
