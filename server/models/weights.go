package models

// Weights
// Хранит в себе веса, использовавшиеся в оценке
type Weights struct {
	ID         uint `gorm:"primary_key" json:"-"`
	ExteriorID uint `gorm:"index" json:"-"`

	UserDefinedMilkType float64 // Молочный тип. Вес, определенный пользователем
	UserDefinedBody     float64 // Туловище. Вес, определенный пользователем
	UserDefinedLimbs    float64 // Конечности. Вес, определенный пользователем
	UserDefinedUdder    float64 // Вымя. Вес, определенный пользователем
	UserDefinedSacrum   float64 // Крестец. Вес, определенный пользователем

	AutomaticMilkType float64 // Молочный тип. Вес по методике
	AutomaticBody     float64 // Туловище. Вес по методике
	AutomaticLimbs    float64 // Конечности. Вес по методике
	AutomaticUdder    float64 // Вымя. Вес по методике
	AutomaticSacrum   float64 // Крестец. Вес по методике

	ChestWidth   float64 // Ширина груди
	BodyDepth    float64 // Глубина туловища
	SacrumHeight float64 // высота в крестце
	SacrumWidth  float64 // ширина в крестце
	BodyType     float64 // Тип телосложения
	Fatness      float64 // Упитанность
	RibsAngle    float64 // Угол наклона ребер
	SacrumAngle  float64 // Угол наклона крестца

	HarmonyOfMovement   float64 // Гармоничность движения
	ForelegWalk         float64 // Поступь передних ног
	HindLegWalkSideView float64 // Поступь задних ног вид сбоку
	HindLegWalkBackView float64 // Поступь задних ног вид cзади
	HoofAngle           float64 // Угол копыта

	UdderDepth                      float64 // Глубина вымени
	FrontNippleLength               float64 // Длинна передних сосков
	FrontNippleLocationBackView     float64 // Расположение передних сосков вид сзади
	BackNippleLocationBackView      float64 // Расположение задних сосков вид сзади
	FrontUdderSegmentsLocation      float64 // Прикрепление передних долей вымени
	BackUdderSegmentsLocationHeight float64 // Высота прикрепления задних долей вымени
	BackUdderSegmentsWidth          float64 // Ширина задних долей вымени
	CentralLigamentDepth            float64 // Глубина центральной связки

	BackBoneQuality     float64 // качество костяка
	Deception           float64 // Обмускульность
	SacrumLength        float64 // Длина крестца
	TopLine             float64 // Линия верха
	UdderBalance        float64 // Баланс вымени
	FrontNippleDiameter float64 // Диаметр передних сосков (9 балов)
	BackNippleDiameter  float64 // Диаметр задних сосков
	UdderVeins          float64 // Выраженность вен вымени
}
