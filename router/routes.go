package router

import (
	"go-web-staging/logger"

	"github.com/gin-gonic/gin"
)

func buildRoutes(r *gin.Engine) {
	r.GET("/")
}

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	buildRoutes(r)
	return r
}
