package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

type RefreshRequest struct {
	RefreshToken string `example:"queuefjad1908)fd_?1" binding:"required"` // refresh токен
}

// Refresh
// Get new access token
// @Summary      Refresh access token
// @Description  Рут обновления access токена
// @Description  В случае успеха возвращает словарь с ключем "access" и значением access токена и ключем "refresh" и значение refresh токена (переданное)
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Param        refreshToken    body     RefreshRequest true  "refresh токен"
// @Tags         Login
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Router       /user/refresh [post]
func (u User) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshReq := &RefreshRequest{}
		if err := c.ShouldBind(refreshReq); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}

		tokenString := strings.TrimPrefix(refreshReq.RefreshToken, "Bearer ")
		tokenString = strings.TrimPrefix(tokenString, "bearer ")
		authClaims := &RefreshTokenClaims{}
		jwtToken, err := jwt.ParseWithClaims(tokenString, authClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil || !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		access, refresh, err := GenerateTokens(authClaims.UserID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"access":  access,
			"refresh": refresh,
		})
	}
}
