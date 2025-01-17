package user

import (
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Email    string `example:"user@mail.ru"`      // Почта пользователя
	Password string `example:"samplePassword123"` // Пароль пользователя
}

// Login
// Get new access and refresh tokens
// @Summary      Login
// @Description  Рут логина
// @Description  В случае успеха возвращает словарь с ключем "access" и значением access токена и ключем "refresh" и значение refresh токена
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Param        loginData    body     LoginData true  "Почта и пароль"
// @Tags         Login
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Router       /user/login [post]
func (u User) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginData := LoginData{}
		if err := c.ShouldBind(&loginData); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{
			"access":  "8321jfda893241jJLDSAJ0a",
			"refresh": "aklfdjadfkslj*(&aj*f8342",
		})
	}
}
