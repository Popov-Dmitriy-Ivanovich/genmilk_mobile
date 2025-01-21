package user

import (
	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	Email    string `example:"FedorsMail@gmail.com"` // Почта
	Password string `example:"FedorsPassword15"`     // Пароль
}

// ChangePassword
// @Summary      Change password for user
// @Description  Смена пользовательского пароля
// @Description  В случае успеха возвращает словарь с ключем "userData" и значением зашифрованных данных пользователя, которые нужно будет передать в confirmPassword
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Param        registerData    body     ChangePasswordRequest true  "Данные смены пароля пользователя"
// @Tags         ChangePassword
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Router       /user/changePassword [post]
func (u User) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &ChangePasswordRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"userData": "932`jfdq78j9804312jf89_--f?0"})
	}
}

type ConfirmPasswordRequest struct {
	UserData string `json:"userData" example:"12308mjkfa01jkfa_!@#"` // Зашифрованные данные пользователя из Register
	Code     string `example:"3295"`                                 // Код с эл. почты пользователя
}

// ConfirmPassword
// @Summary      Confirm password change
// @Description  Подтверждение смены пароля пользователя. После подтверждения у пользователя меняется пароль
// @Description  В случае успеха возвращает словарь с ключем "status" и значением "ok"
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Param        confirmationData    body     ConfirmPasswordRequest true  "Данные для подтверждения"
// @Tags         ChangePassword
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Router       /user/confirmPassword [post]
func (u User) ConfirmPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		confirmReq := ConfirmPasswordRequest{}
		if err := c.ShouldBind(&confirmReq); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}
