package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/common/constants/roles"
	. "github.com/narhakobyan/go-pg-api/http/controllers"
	"github.com/narhakobyan/go-pg-api/http/middlewares"
)

var UserRouter *gin.RouterGroup

func init() {
	UserRouter = ApiRouter.Group("/users")

	UserRouter.Use(middlewares.AuthMiddleware([]roles.RoleType{roles.AdminRole}))

	UserRouter.GET("/", UserController.GetUsers)

	UserRouter.GET("/:id", UserController.GetUser)
	UserRouter.PUT("/:id", UserController.UpdateUser)

	UserRouter.POST("/", UserController.PostUser)
}
