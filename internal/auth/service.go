package auth

import (
	"go-web-staging/internal/requests"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Signup(c *gin.Context, req requests.Signup) (string, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Signup(c *gin.Context, req requests.Signup) (string, error) {
	if err := req.Validate(); err != nil {
		return "err", err
	}

	return "ok", nil
}
