package auth

import (
	"go-web-staging/internal/app"
	"go-web-staging/internal/entity"
	"go-web-staging/pkg/middlewares"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Reg(c *gin.Context, req Reg) (string, error)
	Auth(c *gin.Context, auth Auth) (string, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Reg(c *gin.Context, req Reg) (string, error) {
	user := &entity.User{
		Username: req.Username,
		Password: req.Password,
	}
	// err := s.repo.Create(user)
	token, err := s.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Auth(c *gin.Context, auth Auth) (string, error) {

	return "", nil
}

func (s *service) generateJWT(user *entity.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		middlewares.UserID: user.UserID,
		"username":         user.Username,
		"exp":              time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString([]byte(app.Conf.Secret))
}
