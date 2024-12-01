package middlewares

import (
	"errors"
	"go-web-staging/internal/app"
	res "go-web-staging/pkg/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	UserID = "userID"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if len(header) == 0 {
			res.ResponseError(c, res.CodeAuthFail)
			c.Abort()
		}

		arr := strings.SplitN(header, " ", 2)
		if !(len(arr) == 2 && arr[0] == "Bearer") {
			res.ResponseError(c, res.CodeAuthFail)
			c.Abort()
		}

		token := arr[1]
		jt, err := ParseToken(token)
		if err != nil {
			res.ResponseError(c, res.CodeTokenInvalid)
			c.Abort()
		}

		c.Set(UserID, jt["userID"])
		c.Next()
	}
}

// 解析JWT
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(app.Conf.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
