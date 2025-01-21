package load

import (
	"cow_backend_mobile/models"
	"github.com/gin-gonic/gin"
)

type MeasuresInput struct {
	Cow            models.Cow            // Загружаемая корова
	Exterior       models.Exterior       // Экстерьер коровы
	Measures       models.Measures       // Замеры коровы
	DownSides      models.DownSides      // Недостатки
	Ratings        models.Ratings        // Оценки экстерьера
	AdditionalInfo models.AdditionalInfo // Доп. информация к измерению
	Weights        models.Weights        // Веса использованные в расчете
}

type Load struct{}

func (l Load) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/load")
	apiGroup.POST("/measuredData", l.LoadMeasuredData())
}

// LoadMeasuredData
// @Summary      Рут загрузки данных измерения коровы
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
		c.JSON(200,
			gin.H{
				"loadedData": mInput,
				"status":     "ok",
			})
	}
}
