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

	AutomaticWithoutDownsidesMilkType *float64 // Молочный тип (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesBody     *float64 // Туловище (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesLimbs    *float64 // Конечности (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesUdder    *float64 // Вымя (100 баллов) рассчитанный автоматически без учета недостатков
	AutomaticWithoutDownsidesSacrum   *float64 // Крестец (100 баллов) рассчитанный автоматически без учета недостатков

	AutomaticWithDownsidesMilkType *float64 // Молочный тип (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesBody     *float64 // Туловище (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesLimbs    *float64 // Конечности (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesUdder    *float64 // Вымя (100 баллов) рассчитанный автоматически с учетом недостатков
	AutomaticWithDownsidesSacrum   *float64 // Крестец (100 баллов) рассчитанный автоматически с учетом недостатков
}

type Exterior struct {
	ID    uint `gorm:"primaryKey" json:"-"`
	CowID uint `json:"-"`

	UserID uint `json:"-" gorm:"index"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	AssessmentDate DateOnly       `json:",string" example:"2001-03-23" swaggertype:"string"` // Дата проведения оценочных мероприятий ГГГГ-ММ-ДД
	Measures       Measures       `json:"-"`
	DownSides      DownSides      `json:"-"`
	Ratings        Ratings        `json:"-"`
	AdditionalInfo AdditionalInfo `json:"-"`
	Weights        Weights        `json:"-"`

	Rating   float64 // Суммарный рейтинг (100 баллов)
	Category string  // Категория (хорошо, плохо и т.д.) ХЗ зачем это надо, но мало ли в разных хозяйствах категории разные

	MilkType float64 // Молочный тип (100 баллов)
	Body     float64 // Туловище (100 баллов)
	Limbs    float64 // Конечности (100 баллов)
	Udder    float64 // Вымя (100 баллов)
	Sacrum   float64 // Крестец (100 баллов)

	ChestWidth   float64 // Ширина груди (9 баллов)
	BodyDepth    float64 // Глубина туловища (9 баллов)
	SacrumHeight float64 // высота в крестце (9 баллов)
	SacrumWidth  float64 // ширина в крестце (9 баллов)
	BodyType     float64 // Тип телосложения (9 баллов)
	Fatness      float64 // Упитанность (9 баллов)
	RibsAngle    float64 // Угол наклона ребер (9 баллов)
	SacrumAngle  float64 // Угол наклона крестца (9 баллов)

	HarmonyOfMovement   float64 // Гармоничность движения (9 баллов)
	ForelegWalk         float64 // Поступь передних ног (9 баллов)
	HindLegWalkSideView float64 // Поступь задних ног вид сбоку (9 баллов)
	HindLegWalkBackView float64 // Поступь задних ног вид cзади (9 баллов)
	HoofAngle           float64 // Угол копыта (9 баллов)

	UdderDepth                      float64 // Глубина вымени (9 баллов)
	FrontNippleLength               float64 // Длинна передних сосков (9 баллов)
	FrontNippleLocationBackView     float64 // Расположение передних сосков вид сзади (9 баллов)
	BackNippleLocationBackView      float64 // Расположение задних сосков вид сзади (9 баллов)
	FrontUdderSegmentsLocation      float64 // Прикрепление передних долей вымени (9 баллов)
	BackUdderSegmentsLocationHeight float64 // Высота прикрепления задних долей вымени (9 баллов)
	BackUdderSegmentsWidth          float64 // Ширина задних долей вымени (9 баллов)
	CentralLigamentDepth            float64 // Глубина центральной связки (9 баллов)

	BackBoneQuality     float64 // качество костяка (9 баллов)
	Deception           float64 // Обмускульность (9 баллов)
	SacrumLength        float64 // Длина крестца (9 баллов)
	TopLine             float64 // Линия верха (9 баллов)
	UdderBalance        float64 // Баланс вымени (9 баллов)
	FrontNippleDiameter float64 // Диаметр передних сосков (9 балов)
	BackNippleDiameter  float64 // Диаметр задних сосков (9 баллов)
	UdderVeins          float64 // Выраженность вен вымени (9 баллов)
}
