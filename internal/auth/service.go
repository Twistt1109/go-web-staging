package auth

import (
	"fmt"
	"go-web-staging/internal/entity"
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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI4MTIwMDYsInVzZXJJZCI6MCwidXNlcm5hbWUiOiJ0d2lzdCJ9.AdNWqwIkR1MMcUhQwCjPdg00o_D8q0nSD0-MAury2EE"
	jt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("1111111111", token)
		return []byte("4564564564"), nil
	})

	fmt.Println("222222222222222222", jt.Claims.(jwt.MapClaims)["username"], err)
	return "", nil
}

func (s *service) generateJWT(user *entity.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.UserID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString([]byte("4564564564"))
}
