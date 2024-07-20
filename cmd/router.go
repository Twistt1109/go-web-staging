package main

import (
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	r.GET("/")
}

func buildRoute() *gin.Engine {
	r := gin.New()
	r.Use(GinLogger(), GinRecovery(true))

	routes(r)
	return r
}
