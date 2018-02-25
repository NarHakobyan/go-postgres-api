package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/core/constants/roles"
	. "github.com/narhakobyan/go-pg-api/http/controllers"
	"github.com/narhakobyan/go-pg-api/http/middlewares"
	"github.com/narhakobyan/go-pg-api/http/response"
)

var UserRouter *gin.RouterGroup

func initUserRoutes() {
	UserRouter = ApiRouter.Group("/users")

	UserRouter.Use(middlewares.AuthMiddleware([]roles.RoleType{roles.AdminRole, roles.UserRole}))

	UserRouter.GET("/", response.HandleFunc(UserController.GetUsers))

	UserRouter.GET("/:id", response.HandleFunc(UserController.GetUser))
	UserRouter.PUT("/:id", response.HandleFunc(UserController.UpdateUser))

	UserRouter.POST("/", response.HandleFunc(UserController.PostUser))
}
