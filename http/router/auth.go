package router

import (
	"github.com/NarHakobyan/go-postgres-api/core/constants/roles"
	. "github.com/NarHakobyan/go-postgres-api/http/controllers"
	"github.com/NarHakobyan/go-postgres-api/http/middlewares"
	"github.com/NarHakobyan/go-postgres-api/http/response"
	"github.com/gin-gonic/gin"
)

var AuthRouter *gin.RouterGroup

func initAuthRoutes() {
	AuthRouter = ApiRouter.Group("/auth")

	AuthRouter.POST("/login", response.HandleFunc(AuthController.PostLogin))
	AuthRouter.POST("/register", response.HandleFunc(AuthController.PostRegister))
	AuthRouter.GET("/my-profile", middlewares.AuthMiddleware([]roles.RoleType{roles.AdminRole, roles.UserRole}), response.HandleFunc(AuthController.GetMyProfile))
}
