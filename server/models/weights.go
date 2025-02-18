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

	AutomaticMilkType float64 `` // Молочный тип. Вес по методике
	AutomaticBody     float64 `` // Туловище. Вес по методике
	AutomaticLimbs    float64 `` // Конечности. Вес по методике
	AutomaticUdder    float64 `` // Вымя. Вес по методике
	AutomaticSacrum   float64 `` // Крестец. Вес по методике

	MilkTypeChestWidth      float64 `` // МТ Ширина груди
	MilkTypeBodyDepth       float64 `` // МТ Глубина туловища
	MilkTypeSacrumHeight    float64 `` // МТ высота в крестце
	MilkTypeBodyType        float64 `` // МТ Тип телосложения
	MilkTypeFatness         float64 `` // МТ Упитанность
	MilkTypeRibsAngle       float64 `` // МТ Угол наклона ребер
	MilkTypeBackBoneQuality float64 `` // МТ качество костяка
	MilkTypeTopLine         float64 `` // МТ Линия верха

	BodySacrumHeight float64 `` // Т высота в крестце
	BodyChestWidth   float64 `` // Т Ширина груди
	BodyBodyDepth    float64 `` // Т Глубина туловища
	BodyDeception    float64 `` // Т Обмускульность

	SacrumAngle  float64 `` // Кр Угол наклона крестца
	SacrumWidth  float64 `` // Кр ширина в крестце
	SacrumLength float64 `` // Кр Длина крестца

	LimbsForelegWalk         float64 `` // К Поступь передних ног
	LimbsHindLegWalkSideView float64 `` // К Поступь задних ног вид сбоку
	LimbsHindLegWalkBackView float64 `` // К Поступь задних ног вид cзади
	LimbsDeception           float64 `` // К Обмускульность
	LimbsHoofAngle           float64 `` // К Угол копыта
	LimbsHarmonyOfMovement   float64 `` // К Гармоничность движения

	UdderDepth                      float64 `` // В Глубина вымени
	FrontNippleLength               float64 `` // В Длинна передних сосков
	FrontNippleLocationBackView     float64 `` // В Расположение передних сосков вид сзади
	BackNippleLocationBackView      float64 `` // В Расположение задних сосков вид сзади
	FrontUdderSegmentsLocation      float64 `` // В Прикрепление передних долей вымени
	BackUdderSegmentsLocationHeight float64 `` // В Высота прикрепления задних долей вымени
	BackUdderSegmentsWidth          float64 `` // В Ширина задних долей вымени
	CentralLigamentDepth            float64 `` // В Глубина центральной связки

	UdderBalance        float64 `` // В Баланс вымени
	FrontNippleDiameter float64 `` // В Диаметр передних сосков (9 балов)
	BackNippleDiameter  float64 `` // В Диаметр задних сосков
	UdderVeins          float64 `` // В Выраженность вен вымени
	//SacrumAngle         float64 `` // Угол наклона крестца

}
