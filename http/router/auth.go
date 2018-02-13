package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/narhakobyan/go-pg-api/http/controllers"
)

var AuthRouter *gin.RouterGroup

func initAuthRoutes() {
	AuthRouter = ApiRouter.Group("/auth")

	AuthRouter.POST("/login", AuthController.PostLogin)

	AuthRouter.POST("/register", AuthController.PostRegister)
}
