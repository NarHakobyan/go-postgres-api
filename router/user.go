package router

import "github.com/gin-gonic/gin"

var UserRouter *gin.RouterGroup

func init() {
	UserRouter = ApiRouter.Group("/user")

	UserRouter.GET("/", func(context *gin.Context) {
		context.String(200, "working")
	})

	UserRouter.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
}
