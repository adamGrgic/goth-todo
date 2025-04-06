package auth

import (
	"fmt"
	"net/http"
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
	token, err := c.Cookie("jwt_token")
	if err != nil {
		return nil
	}
	return &token
}

func SetUserJWT(c *gin.Context, username string) error {
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return err
	}

	c.Set("jwt_token", signedToken)
	fmt.Println("JWT set:", signedToken)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt_token",
		Value:    signedToken,
		MaxAge:   10000,
		Path:     "/",
		Domain:   "192.168.0.26",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

// func GetUser(c *gin.Context) {
// 	c.Get("jwt")
// }
