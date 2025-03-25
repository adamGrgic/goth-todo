package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your-secret-key")

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetUserToken(c *gin.Context) *string {
	token, exists := c.Get("jwt_token")
	if !exists {
		return nil
	}

	strToken, ok := token.(string)
	if !ok {
		return nil
	}

	return &strToken
}

func GenerateJWT(username string) (string, error) {
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// func GetUser(c *gin.Context) {
// 	c.Get("jwt")
// }
