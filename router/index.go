package router

import "github.com/gin-gonic/gin"

var ApiRouter *gin.RouterGroup
var Router *gin.Engine

func init() {
	Router = gin.Default()

	ApiRouter = Router.Group("/api/v1")
}
