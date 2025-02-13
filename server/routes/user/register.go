package user

import (
	"cow_backend_mobile/models"
	"log"
	"math/rand/v2"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email                 string `example:"User321@gmail.com" `      // Почта
	NameSurnamePatronymic string `example:"Иванов Федор Петрович"  ` // Фио
	Password              string `example:"FedorsPassword15" `       // Пароль
	LicenseNumber         *string
}

type RegisterUserData struct {
	Code uint
	jwt.RegisteredClaims
	RegisterRequest
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

		userData := RegisterUserData{}
		userData.RegisterRequest = regReq

		jwtKey := os.Getenv("JWT_KEY")
		expTimeRegisterToken := time.Now().Add(1 * time.Hour)
		code := rand.IntN(8997) + 1001

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, RegisterUserData{
			Code:            uint(code),
			RegisterRequest: regReq,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTimeRegisterToken),
			},
		})

		encryptedUserData, err := accessToken.SignedString([]byte(jwtKey))

		if err != nil {
			c.JSON(500, gin.H{"error": "ошибка создания токена"})
			return
		}

		from := os.Getenv("EMAIL_FROM")
		password := os.Getenv("EMAIL_PASS")
		to := []string{userData.Email}
		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort := os.Getenv("SMTP_PORT")
		message := []byte("From: genmilk@aurusoft.ru\r\n" +
			"To: " + userData.Email + "\r\n" +
			"Subject: Подтвердите эл. почту\r\n" +
			"\r\n" +
			"Код подтверждения электронной почты:  " + strconv.FormatUint(uint64(code), 10) + " .\r\n")
		auth := smtp.PlainAuth("", from, password, smtpHost)
		log.Println("Email to: ", to)
		if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"userData": encryptedUserData})
	}
}

type ConfirmMailRequest struct {
	UserData string `json:"userData" example:"12308mjkfa01jkfa_!@#" ` // Зашифрованные данные пользователя из Register
	Code     string `example:"3295" `                                 // Код с эл. почты пользователя
}

// ConfirmMail
// @Summary      Register new user
// @Description  Подтверждение почты пользователя. После подтверждения пользователь становится зарегистрированным
// @Description  В случае успеха возвращает словарь с ключем "status" и значением "ok"
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки
// @Description  Код 406 - неверный код подтверждения
// @Description  Код 409 - пользователь с таким email уже существует
// @Param        confirmationData    body     ConfirmMailRequest true  "Данные для подтверждения"
// @Tags         Register
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Failure      406  {object}   map[string]string
// @Failure      409  {object}   map[string]string
// @Router       /user/confirmMail [post]
func (u User) ConfirmMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		confirmReq := ConfirmMailRequest{}
		if err := c.ShouldBind(&confirmReq); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}
		regData := RegisterUserData{}
		token, err := jwt.ParseWithClaims(confirmReq.UserData, &regData, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(422, gin.H{"error": "ошибка подтверждения:" + err.Error()})
			return
		}
		if confirmReq.Code != strconv.FormatUint(uint64(regData.Code), 10) {
			c.JSON(406, gin.H{"error": "Неверный код подтверждения"})
			return
		}
		db := models.GetDatabase()
		newUser := models.User{}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}

		searchRes := db.Where("email = ?", regData.Email).Limit(1).Find(&newUser)
		if searchRes.Error != nil {
			c.JSON(500, gin.H{"error": searchRes.Error.Error()})
			return
		}
		if searchRes.RowsAffected != 0 {
			c.JSON(409, gin.H{"error": "Пользователь с таким email уже существует"})
			return
		}

		newUser.Email = regData.Email
		newUser.NameSurnamePatronymic = regData.NameSurnamePatronymic
		newUser.Password = hashedPassword
		newUser.LicenceNumber = regData.LicenseNumber
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}
