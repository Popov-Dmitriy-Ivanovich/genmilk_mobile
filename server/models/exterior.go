package models

// Ratings
// Все поля оценок являются опциональными. Критично, чтобы в экстерьере были хоть какие-нибудь 100-бальные оценки,
// этих оценок нет в БД genmilk.ru.
type Ratings struct {
	ID         uint `gorm:"primary_key" json:"-"`
	ExteriorID uint `json:"-" gorm:"index"`

	UserDefinedWithoutDownsidesMilkType *float64 // Молочный тип (100 баллов) без учета недостатков, оценка пользователя
	UserDefinedWithoutDownsidesBody     *float64 // Туловище (100 баллов) без учета недостатков, оценка пользователя
	UserDefinedWithoutDownsidesLimbs    *float64 // Конечности (100 баллов) без учета недостатков, оценка пользователя
	UserDefinedWithoutDownsidesUdder    *float64 // Вымя (100 баллов) без учета недостатков, оценка пользователя
	UserDefinedWithoutDownsidesSacrum   *float64 // Крестец (100 баллов) без учета недостатков, оценка пользователя

	UserDefinedWithDownsidesMilkType *float64 // Молочный тип (100 баллов) с учетом недостатков, оценка пользователя
	UserDefinedWithDownsidesBody     *float64 // Туловище (100 баллов) с учетом недостатков, оценка пользователя
	UserDefinedWithDownsidesLimbs    *float64 // Конечности (100 баллов) с учетом недостатков, оценка пользователя
	UserDefinedWithDownsidesUdder    *float64 // Вымя (100 баллов) с учетом недостатков, оценка пользователя
	UserDefinedWithDownsidesSacrum   *float64 // Крестец (100 баллов) с учетом недостатков, оценка пользователя

	AutomaticWithoutDownsidesMilkType float64 `` // Молочный тип (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesBody     float64 `` // Туловище (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesLimbs    float64 `` // Конечности (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesUdder    float64 `` // Вымя (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesSacrum   float64 `` // Крестец (100 баллов) рассчитанный автоматически без учета недостатков

	AutomaticWithDownsidesMilkType float64 `` // Молочный тип (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesBody     float64 `` // Туловище (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesLimbs    float64 `` // Конечности (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesUdder    float64 `` // Вымя (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesSacrum   float64 `` // Крестец (100 баллов) рассчитанный автоматически с учетом недостатков
}

type Exterior struct {
	ID    uint `gorm:"primaryKey" json:"-"`
	CowID uint `json:"-"`

	UserID uint `json:"-" gorm:"index"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	AssessmentDate DateOnly `json:",string" example:"2001-03-23" swaggertype:"string" ` // Дата проведения оценочных мероприятий ГГГГ-ММ-ДД
	Measures       Measures
	DownSides      *DownSides
	Ratings        Ratings
	AdditionalInfo AdditionalInfo
	Weights        Weights

	Rating   float64 `validate:"lte=100,gte=-100"` // Суммарный рейтинг (100 баллов)
	Category string  ``                            // Категория (хорошо, плохо и т.д.) ХЗ зачем это надо, но мало ли в разных хозяйствах категории разные

	MilkType float64 `validate:"lte=100,gte=-100"` // Молочный тип (100 баллов)
	Body     float64 `validate:"lte=100,gte=-100"` // Туловище (100 баллов)
	Limbs    float64 `validate:"lte=100,gte=-100"` // Конечности (100 баллов)
	Udder    float64 `validate:"lte=100,gte=-100"` // Вымя (100 баллов)
	Sacrum   float64 `validate:"lte=100,gte=-100"` // Крестец (100 баллов)

	ChestWidth   float64 `validate:"lte=9,gte=-9"` // Ширина груди (9 баллов)
	BodyDepth    float64 `validate:"lte=9,gte=-9"` // Глубина туловища (9 баллов)
	SacrumHeight float64 `validate:"lte=9,gte=-9"` // Высота в крестце (9 баллов)
	SacrumWidth  float64 `validate:"lte=9,gte=-9"` // Ширина в крестце (9 баллов)
	BodyType     float64 `validate:"lte=9,gte=-9"` // Тип телосложения (9 баллов)
	Fatness      float64 `validate:"lte=9,gte=-9"` // Упитанность (9 баллов)
	RibsAngle    float64 `validate:"lte=9,gte=-9"` // Угол наклона ребер (9 баллов)
	SacrumAngle  float64 `validate:"lte=9,gte=-9"` // Угол наклона крестца (9 баллов)

	HarmonyOfMovement   float64 `validate:"lte=9,gte=-9"` // Гармоничность движения (9 баллов)
	ForelegWalk         float64 `validate:"lte=9,gte=-9"` // Поступь передних ног (9 баллов)
	HindLegWalkSideView float64 `validate:"lte=9,gte=-9"` // Поступь задних ног вид сбоку (9 баллов)
	HindLegWalkBackView float64 `validate:"lte=9,gte=-9"` // Поступь задних ног вид cзади (9 баллов)
	HoofAngle           float64 `validate:"lte=9,gte=-9"` // Угол копыта (9 баллов)

	UdderDepth                      float64 `validate:"lte=9,gte=-9"` // Глубина вымени (9 баллов)
	FrontNippleLength               float64 `validate:"lte=9,gte=-9"` // Длинна передних сосков (9 баллов)
	FrontNippleLocationBackView     float64 `validate:"lte=9,gte=-9"` // Расположение передних сосков вид сзади (9 баллов)
	BackNippleLocationBackView      float64 `validate:"lte=9,gte=-9"` // Расположение задних сосков вид сзади (9 баллов)
	FrontUdderSegmentsLocation      float64 `validate:"lte=9,gte=-9"` // Прикрепление передних долей вымени (9 баллов)
	BackUdderSegmentsLocationHeight float64 `validate:"lte=9,gte=-9"` // Высота прикрепления задних долей вымени (9 баллов)
	BackUdderSegmentsWidth          float64 `validate:"lte=9,gte=-9"` // Ширина задних долей вымени (9 баллов)
	CentralLigamentDepth            float64 `validate:"lte=9,gte=-9"` // Глубина центральной связки (9 баллов)

	BackBoneQuality     float64 `validate:"lte=9,gte=-9"` // Качество костяка (9 баллов)
	Deception           float64 `validate:"lte=9,gte=-9"` // Обмускульность (9 баллов)
	SacrumLength        float64 `validate:"lte=9,gte=-9"` // Длина крестца (9 баллов)
	TopLine             float64 `validate:"lte=9,gte=-9"` // Линия верха (9 баллов)
	UdderBalance        float64 `validate:"lte=9,gte=-9"` // Баланс вымени (9 баллов)
	FrontNippleDiameter float64 `validate:"lte=9,gte=-9"` // Диаметр передних сосков (9 балов)
	BackNippleDiameter  float64 `validate:"lte=9,gte=-9"` // Диаметр задних сосков (9 баллов)
	UdderVeins          float64 `validate:"lte=9,gte=-9"` // Выраженность вен вымени (9 баллов)
}
