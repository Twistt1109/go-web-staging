package main

import (
	"go-web-staging/internal/auth"
	"go-web-staging/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	rg := r.Group("v1")

	auth.ServeResouce(rg, auth.NewService(auth.NewRepo()))

	rg.Use(middlewares.JWTAuth())
}

func buildRouter() *gin.Engine {
	r := gin.New()
	r.Use(GinLogger(), GinRecovery(true))

	routes(r)
	return r
}
