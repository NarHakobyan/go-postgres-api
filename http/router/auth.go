package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/core/constants/roles"
	. "github.com/narhakobyan/go-pg-api/http/controllers"
	"github.com/narhakobyan/go-pg-api/http/middlewares"
	"github.com/narhakobyan/go-pg-api/http/response"
)

var AuthRouter *gin.RouterGroup

func initAuthRoutes() {
	AuthRouter = ApiRouter.Group("/auth")

	AuthRouter.POST("/login", response.HandleFunc(AuthController.PostLogin))
	AuthRouter.POST("/register", response.HandleFunc(AuthController.PostRegister))
	AuthRouter.GET("/my-profile", middlewares.AuthMiddleware([]roles.RoleType{roles.AdminRole, roles.UserRole}), response.HandleFunc(AuthController.GetMyProfile))
}
