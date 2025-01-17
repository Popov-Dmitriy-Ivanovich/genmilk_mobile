package user

import (
	"github.com/gin-gonic/gin"
)

type RefreshRequest struct {
	RefreshToken string `example:"queuefjad1908)fd_?1"` // refresh токен
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
		if err := c.ShouldBind(&RefreshRequest{}); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{
			"access":  "FKLADFJJl98jfkald",
			"refresh": "aklfdjadfkslj*(&aj*f8342",
		})
	}
}
