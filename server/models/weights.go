package models

// Weights
// Хранит в себе веса, использовавшиеся в оценке
type Weights struct {
	ID         uint `gorm:"primary_key" json:"-"`
	ExteriorID uint `gorm:"index" json:"-"`

	UserDefinedMilkType *float64 // Молочный тип. Вес, определенный пользователем
	UserDefinedBody     *float64 // Туловище. Вес, определенный пользователем
	UserDefinedLimbs    *float64 // Конечности. Вес, определенный пользователем
	UserDefinedUdder    *float64 // Вымя. Вес, определенный пользователем
	UserDefinedSacrum   *float64 // Крестец. Вес, определенный пользователем

	AutomaticMilkType float64 `binding:"required"` // Молочный тип. Вес по методике
	AutomaticBody     float64 `binding:"required"` // Туловище. Вес по методике
	AutomaticLimbs    float64 `binding:"required"` // Конечности. Вес по методике
	AutomaticUdder    float64 `binding:"required"` // Вымя. Вес по методике
	AutomaticSacrum   float64 `binding:"required"` // Крестец. Вес по методике

	ChestWidth   float64 `binding:"required"` // Ширина груди
	BodyDepth    float64 `binding:"required"` // Глубина туловища
	SacrumHeight float64 `binding:"required"` // высота в крестце
	SacrumWidth  float64 `binding:"required"` // ширина в крестце
	BodyType     float64 `binding:"required"` // Тип телосложения
	Fatness      float64 `binding:"required"` // Упитанность
	RibsAngle    float64 `binding:"required"` // Угол наклона ребер
	SacrumAngle  float64 `binding:"required"` // Угол наклона крестца

	HarmonyOfMovement   float64 `binding:"required"` // Гармоничность движения
	ForelegWalk         float64 `binding:"required"` // Поступь передних ног
	HindLegWalkSideView float64 `binding:"required"` // Поступь задних ног вид сбоку
	HindLegWalkBackView float64 `binding:"required"` // Поступь задних ног вид cзади
	HoofAngle           float64 `binding:"required"` // Угол копыта

	UdderDepth                      float64 `binding:"required"` // Глубина вымени
	FrontNippleLength               float64 `binding:"required"` // Длинна передних сосков
	FrontNippleLocationBackView     float64 `binding:"required"` // Расположение передних сосков вид сзади
	BackNippleLocationBackView      float64 `binding:"required"` // Расположение задних сосков вид сзади
	FrontUdderSegmentsLocation      float64 `binding:"required"` // Прикрепление передних долей вымени
	BackUdderSegmentsLocationHeight float64 `binding:"required"` // Высота прикрепления задних долей вымени
	BackUdderSegmentsWidth          float64 `binding:"required"` // Ширина задних долей вымени
	CentralLigamentDepth            float64 `binding:"required"` // Глубина центральной связки

	BackBoneQuality     float64 `binding:"required"` // качество костяка
	Deception           float64 `binding:"required"` // Обмускульность
	SacrumLength        float64 `binding:"required"` // Длина крестца
	TopLine             float64 `binding:"required"` // Линия верха
	UdderBalance        float64 `binding:"required"` // Баланс вымени
	FrontNippleDiameter float64 `binding:"required"` // Диаметр передних сосков (9 балов)
	BackNippleDiameter  float64 `binding:"required"` // Диаметр задних сосков
	UdderVeins          float64 `binding:"required"` // Выраженность вен вымени
}
