package user

import "github.com/gin-gonic/gin"

type User struct{}

func (u User) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/user")
	apiGroup.POST("/login", u.Login())
	apiGroup.POST("/refresh", u.Refresh())
	apiGroup.POST("/register", u.Register())
	apiGroup.POST("/confirmMail", u.ConfirmMail())
	apiGroup.POST("/changePassword", u.ChangePassword())
	apiGroup.POST("/confirmPassword", u.ConfirmPassword())
}
