package middleware

import (
	"cow_backend_mobile/routes/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

type AuthConfig struct {
	Placeholder string
}

func AuthMiddleware(config AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		tokenString = strings.TrimPrefix(tokenString, "bearer ")
		authClaims := &user.AccessTokenClaims{}
		jwtToken, err := jwt.ParseWithClaims(tokenString, authClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil || !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("UserID", authClaims.UserID)
		c.Next()
	}
}
