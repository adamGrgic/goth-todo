package middleware

import (
	// "net/http"
	// "strings"

	"goth-todo/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
	// "goth-todo/internal/auth"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := auth.GetUserToken(c)
		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		// 	return
		// }

		// tokenStr := parts[1]

		// token, err := jwt.ParseWithClaims(tokenStr, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 	return auth.JwtKey, nil
		// })

		// if err != nil || !token.Valid {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	return
		// }

		// if claims, ok := token.Claims.(*auth.CustomClaims); ok {
		// 	c.Set("username", claims.Username)
		// }

		c.Next()
	}
}
