package router

import (
	"github.com/NarHakobyan/go-postgres-api/core/constants/roles"
	. "github.com/NarHakobyan/go-postgres-api/http/controllers"
	"github.com/NarHakobyan/go-postgres-api/http/middlewares"
	"github.com/NarHakobyan/go-postgres-api/http/response"
	"github.com/gin-gonic/gin"
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
