package main

import (
	_ "cow_backend_mobile/docs"
	"cow_backend_mobile/models"
	"cow_backend_mobile/routes"
	"cow_backend_mobile/routes/load"
	"cow_backend_mobile/routes/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GenMilk mobile API
// @version         1.0
// @description     Сваггер сгенерирован из кода, поэтому может содержать неточности. По мере возможности они будут описаны далее
// @description     Все даты передаются как строки
// @description     Большая часть рутов не возвращает вложенные объекты
// @description 	Авторизация по JWT, стандартная
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      77.221.145.143:8080
// @BasePath  /api/mobile

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	apiGroup := r.Group("/api/mobile")
	routes.WriteRoutes(apiGroup, &load.Load{}, &user.User{})
	models.GetDatabase()
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//apiGroup.Static("/static", "static")
	r.Run()
}
