package main

import (
	"go-web-staging/internal/auth"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	rg := r.Group("v1")

	auth.ServeResouce(rg, auth.NewService())
}

func buildRouter() *gin.Engine {
	r := gin.New()
	r.Use(GinLogger(), GinRecovery(true))

	routes(r)
	return r
}
