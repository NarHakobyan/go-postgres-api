package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/core/constants/roles"
	. "github.com/narhakobyan/go-pg-api/http/controllers"
	"github.com/narhakobyan/go-pg-api/http/middlewares"
)

var UserRouter *gin.RouterGroup

func initUserRoutes() {
	UserRouter = ApiRouter.Group("/users")

	UserRouter.Use(middlewares.AuthMiddleware([]roles.RoleType{roles.AdminRole, roles.UserRole}))

	UserRouter.GET("/", UserController.GetUsers)

	UserRouter.GET("/:id", UserController.GetUser)
	UserRouter.PUT("/:id", UserController.UpdateUser)

	UserRouter.POST("/", UserController.PostUser)
}
