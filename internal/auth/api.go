package auth

import (
	"go-web-staging/internal/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type resource struct {
}

func ServeResouce(rg *gin.RouterGroup) {
	r := &resource{}

	rg.POST("/signup", r.signup)
}

func (r *resource) signup(c *gin.Context) {
	// 1. 获取参数
	signup := &requests.Signup{}
	if err := c.ShouldBind(signup); err != nil {
		zap.L().Error("signup should bind", zap.Error(err))
		c.JSON(401, gin.H{"msg": "参数错误"})
		return
	}

	// 2. 校验参数
	// 3. 调用业务
	// 4. 返回结果
	c.JSON(201, gin.H{
		"msg":      "success",
		"username": signup.Username,
		"password": signup.Password,
	})
}
