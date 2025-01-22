package user

import (
	"cow_backend_mobile/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand/v2"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

type ChangePasswordRequest struct {
	Email    string `example:"FedorsMail@gmail.com"` // Почта
	Password string `example:"FedorsPassword15"`     // Пароль
}

type ChangePasswordUserData struct {
	Code uint
	jwt.RegisteredClaims
	ChangePasswordRequest
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
		req := ChangePasswordRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}

		code := rand.IntN(8997) + 1001
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ChangePasswordUserData{
			Code:                  uint(code),
			ChangePasswordRequest: req,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		})

		encryptedUserData, err := accessToken.SignedString([]byte(os.Getenv("JWT_KEY")))

		if err != nil {
			c.JSON(500, gin.H{"error": "ошибка создания токена"})
			return
		}

		from := os.Getenv("EMAIL_FROM")
		password := os.Getenv("EMAIL_PASS")
		to := []string{req.Email}
		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort := os.Getenv("SMTP_PORT")
		message := []byte("From: genmilk@aurusoft.ru\r\n" +
			"To: " + req.Email + "\r\n" +
			"Subject: Подтвердите эл. почту\r\n" +
			"\r\n" +
			"Код подтверждения смены пароля:  " + strconv.FormatUint(uint64(code), 10) + " .\r\n")
		auth := smtp.PlainAuth("", from, password, smtpHost)
		log.Println("Email to: ", to)
		if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"userData": encryptedUserData})
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

		pasChangeData := ChangePasswordUserData{}
		token, err := jwt.ParseWithClaims(confirmReq.UserData, &pasChangeData, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(422, gin.H{"error": "ошибка подтверждения:" + err.Error()})
			return
		}

		if confirmReq.Code != strconv.FormatUint(uint64(pasChangeData.Code), 10) {
			c.JSON(406, gin.H{"error": "Неверный код подтверждения"})
			return
		}

		db := models.GetDatabase()
		user := models.User{}

		if err := db.First(&user, "email = ?", pasChangeData.Email).Error; err != nil {
			c.JSON(404, gin.H{"error": "Не найден пользователь с таким email", "message": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pasChangeData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		user.Password = hashedPassword
		if err := db.Save(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}
