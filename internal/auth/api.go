package auth

import (
	"fmt"
	res "go-web-staging/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type resource struct {
	service Service
}

func ServeResouce(rg *gin.RouterGroup, service Service) {
	r := &resource{service}

	rg.POST("/reg", r.reg)

	rg.POST("/auth", r.auth)
}

func (r *resource) reg(c *gin.Context) {
	// 1. 获取参数
	reg := &Reg{}
	if err := c.ShouldBind(reg); err != nil {
		zap.L().Error("reg should bind", zap.Error(err))
		res.ResponseErrorWithMsg(c, res.CodeInvalidParam, err.Error())
		return
	}

	// 2. 校验参数
	// if err := reg.Validate(); err != nil {
	// 	c.JSON(401, gin.H{"msg": err.Error()})
	// 	return
	// }

	// 3. 调用业务
	token, err := r.service.Reg(c, *reg)
	fmt.Println("-----------------token", token)
	// 4. 返回结果
	res.Response(c, token, err)
}

func (r *resource) auth(c *gin.Context) {

	var input Auth
	if err := c.ShouldBind(&input); err != nil {
		zap.L().Error("auth should bind", zap.Error(err))
		res.ResponseErrorWithMsg(c, res.CodeInvalidParam, err.Error())
		return
	}

	r.service.Auth(c, input)
}
