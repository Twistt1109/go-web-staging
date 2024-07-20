package auth

import "github.com/gin-gonic/gin"

type resource struct {
}

func ServeResouce(rg *gin.RouterGroup) {
	rg.GET("/signup", func(c *gin.Context) {
	})
}
