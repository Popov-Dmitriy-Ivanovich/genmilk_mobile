package cows

import (
	"cow_backend_mobile/models"
	"cow_backend_mobile/routes/middleware"

	"github.com/gin-gonic/gin"
)

type Cows struct{}

func (c Cows) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/cows")
	apiGroup.Use(middleware.AuthMiddleware(middleware.AuthConfig{}))
	apiGroup.GET("/cows", c.Get())
	apiGroup.GET("/cows/:id")
}

// Get
// @Summary      Рут получения списка id коров
// @Description  Возвращает массив id всех коров, внесенных в базу данных
// @Tags         Cows
// @Produce      json
// @Success      200  {object}   []uint
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Failure      401  {object}   map[string]string
// @Router       /cows [get]
func (c Cows) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		ids := []uint{}
		db := models.GetDatabase()
		if err := db.Model(models.Cow{}).Pluck("id", &ids).Error; err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, ids)
	}
}

// Get
// @Summary      Рут получения списка id коров
// @Description  Возвращает массив id всех коров, внесенных в базу данных
// @Param        id    path     uint true  "Данные измерения коровы"
// @Tags         Cows
// @Produce      json
// @Success      200  {object}   models.Cow
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Failure      401  {object}   map[string]string
// @Router       /cows/{id} [get]
func (c Cows) GetConcrete() gin.HandlerFunc {
	return func(c *gin.Context) {
		cow_id := c.Param("id")
		cow := models.Cow{}
		db := models.GetDatabase()
		if err := db.
			Preload("Exteriors").
			Preload("Exteriors.User").
			Preload("Exteriors.Measures").
			Preload("Exteriors.DownSides").
			Preload("Exteriors.Ratings").
			Preload("Exteriors.AdditionalInfo").
			Preload("Exteriors.Weights").
			First(&cow, cow_id).Error; err != nil {
			c.AbortWithError(404, err)
			return
		}
		c.JSON(200, cow)
	}
}
