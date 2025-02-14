package user

import (
	"cow_backend_mobile/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Email    string `example:"user@mail.ru" `      // Почта пользователя
	Password string `example:"samplePassword123" ` // Пароль пользователя
}

// Login
// Get new access and refresh tokens
// @Summary      Login
// @Description  Рут логина
// @Description  В случае успеха возвращает словарь с ключем "access" и значением access токена и ключем "refresh" и значение refresh токена
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Description  Также в случае успеха по ключу "user" доступны пользовательские данные
// @Description  В случае если пользователь не найден 404
// @Param        loginData    body     LoginData true  "Почта и пароль"
// @Tags         Login
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Failure		 404  {object}   map[string]string
// @Router       /user/login [post]
func (u User) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginData := LoginData{}
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		db := models.GetDatabase()
		user := models.User{}

		if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
			c.JSON(404, gin.H{"error": "Пользователь с таким email не найден"})
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(loginData.Password)); err != nil {
			c.JSON(401, gin.H{"error": "Пароли не совпадают", "message": err.Error()})
			return
		}

		access, refresh, err := GenerateTokens(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"access":  access,
			"refresh": refresh,
			"user":    user,
		})
	}
}
