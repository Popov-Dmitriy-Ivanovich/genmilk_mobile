package load

import (
	"cow_backend_mobile/models"
	"cow_backend_mobile/routes/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MeasuresInput struct {
	Cow            models.Cow            `validate:"required"` // Загружаемая корова
	Exterior       models.Exterior       `validate:"required"` // Экстерьер коровы
	Measures       models.Measures       `validate:"required"` // Замеры коровы
	DownSides      *models.DownSides     // Недостатки
	Ratings        models.Ratings        `validate:"required"` // Оценки экстерьера
	AdditionalInfo models.AdditionalInfo `validate:"required"` // Доп. информация к измерению
	Weights        models.Weights        `validate:"required"` // Веса использованные в расчете
}

func (mi MeasuresInput) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(mi.Cow); err != nil {
		return err
	}

	if err := validate.Struct(mi.Exterior); err != nil {
		return err
	}
	return nil
}

type Load struct{}

func (l Load) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/load")
	apiGroup.Use(middleware.AuthMiddleware(middleware.AuthConfig{}))
	apiGroup.POST("/measuredData", l.LoadMeasuredData())
}

// LoadMeasuredData
// @Summary      Рут загрузки данных измерения коровы
// @Description  Требует авторизации!
// @Description  Принимает данные оценки, парсит, создает запись в БД.
// @Description  В случае успеха возвращает словарь с ключем "status" и значением "ok".
// @Description  В случае ошибки возвращает словарь с ключем "error" и строкой ошибки.
// @Description  В этом проекте предлагается хранить 4 группы 100-бальных оценок (с/без недостатками + авто/пользовательские).
// @Description  В genmilk.ru хранится только одна группа 100-бальных признаков, поля MilkType, Body, Limbs ...
// @Description  в экстерьере означают те 100-бальные оценки, которые будут в базе genmilk.ru.
// @Description  Разделитель для строк, содержащих недостатки выбран из соображения о том, что URL работает с
// @Description  разделителем "/", других причин для выбора "/" нет, можно перевыбрать, это совершенно не принципиально
// @Param        measures    body     MeasuresInput true  "Данные измерения коровы"
// @Tags         Load
// @Produce      json
// @Success      200  {object}   map[string]string
// @Failure      500  {object}   map[string]string
// @Failure      422  {object}   map[string]string
// @Failure      401  {object}   map[string]string
// @Router       /load/measuredData [post]
func (l Load) LoadMeasuredData() gin.HandlerFunc {
	return func(c *gin.Context) {
		mInput := MeasuresInput{}
		if err := c.ShouldBindJSON(&mInput); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}
		if err := mInput.Validate(); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}
		db := models.GetDatabase()
		cow := models.Cow{}
		if err := db.FirstOrCreate(&cow, mInput.Cow).Error; err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}
		val, ok := c.Get("UserID")
		if !ok {
			c.JSON(401, gin.H{"error": "No User Found"})
			return
		}
		userId, ok := val.(uint)
		if !ok {
			c.JSON(500, gin.H{"error": "UserId has wrong type"})
			return
		}
		user := models.User{}
		if err := db.First(&user, userId).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		cowExterior := mInput.Exterior
		cowExterior.User = user
		cowExterior.Measures = mInput.Measures
		cowExterior.DownSides = mInput.DownSides
		cowExterior.Ratings = mInput.Ratings
		cowExterior.AdditionalInfo = mInput.AdditionalInfo
		cowExterior.Weights = mInput.Weights
		cow.Exteriors = append(cow.Exteriors, cowExterior)

		if err := db.Save(&cow).Error; err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200,
			gin.H{
				"loadedData": mInput,
				"status":     "ok",
			})
	}
}
