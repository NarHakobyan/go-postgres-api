package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/core/constants/roles"
	. "github.com/narhakobyan/go-pg-api/http/controllers"
	"github.com/narhakobyan/go-pg-api/http/middlewares"
)

var AuthRouter *gin.RouterGroup

func initAuthRoutes() {
	AuthRouter = ApiRouter.Group("/auth")

	AuthRouter.POST("/login", AuthController.PostLogin)
	AuthRouter.POST("/register", AuthController.PostRegister)
	AuthRouter.GET("/my-profile", middlewares.AuthMiddleware([]roles.RoleType{roles.AdminRole, roles.UserRole}), AuthController.GetMyProfile)
}
