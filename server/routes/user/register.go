package user

import (
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email                 string `example:"User321@gmail.com"`     // Почта
	NameSurnamePatronymic string `example:"Иванов Федор Петрович"` // Фио
	Password              string `example:"FedorsPassword15"`      // Пароль
}

// Register
// @Summary      Register new user
// @Description  Регистрация пользователя
// @Description  В случае успеха возвращает словарь с ключем "userData" и значением зашифрованных данных пользователя, которые нужно будет передать в confirmMail
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Param        registerData    body     RegisterRequest true  "Данные регистрации пользователя"
// @Tags         Register
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Router       /user/register [post]
func (u User) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		regReq := RegisterRequest{}
		if err := c.ShouldBind(&regReq); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"userData": "932`jfdq78j9804312jf89_--f?0"})
	}
}

type ConfirmMailRequest struct {
	UserData string `json:"userData" example:"12308mjkfa01jkfa_!@#"` // Зашифрованные данные пользователя из Register
	Code     string `example:"3295"`                                 // Код с эл. почты пользователя
}

// ConfirmMail
// @Summary      Register new user
// @Description  Подтверждение почты пользователя. После подтверждения пользователь становится зарегистрированным
// @Description  В случае успеха возвращает словарь с ключем "status" и значением "ok"
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Param        confirmationData    body     ConfirmMailRequest true  "Данные для подтверждения"
// @Tags         Register
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Router       /user/confirmMail [post]
func (u User) ConfirmMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		confirmReq := ConfirmMailRequest{}
		if err := c.ShouldBind(&confirmReq); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}
