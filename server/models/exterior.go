package models

type Exterior struct {
	ID    uint `gorm:"primaryKey" json:"-"`
	CowID uint `json:"-"`

	Measures Measures `json:"-"`

	Rating float64 // Суммарный рейтинг (100 баллов)

	// 100 бальные
	MilkType float64 // Молочный тип (100 баллов)
	Body     float64 // Туловище (100 баллов)
	Limbs    float64 // Конечности (100 баллов)
	Udder    float64 // Вымя (100 баллов)
	Sacrum   float64 // Крестец (100 баллов)
	// 100 бальные

	// Телосложение (8 признаков)
	ChestWidth   float64 // Ширина груди (9 баллов)
	BodyDepth    float64 // Глубина туловища (9 баллов)
	SacrumHeight float64 // высота в крестце (9 баллов)
	SacrumWidth  float64 // ширина в крестце (9 баллов)
	BodyType     float64 // Тип телосложения (9 баллов)
	Fatness      float64 // Упитанность (9 баллов)
	RibsAngle    float64 // Угол наклона ребер (9 баллов)
	SacrumAngle  float64 // Угол наклона крестца (9 баллов)

	// Конечности (5 признаков)
	HarmonyOfMovement   float64 // Гармоничность движения (9 баллов)
	ForelegWalk         float64 // Поступь передних ног (9 баллов)
	HindLegWalkSideView float64 // Поступь задних ног сбоку (9 баллов)
	HindLegWalkBackView float64 // Поступь задних ног cзади (9 баллов)
	HoofAngle           float64 // Угол копыта (9 баллов)

	// Вымя (8 признаков)
	UdderDepth                      float64 // Глубина вымени (9 баллов)
	FrontNippleLength               float64 // Длинна передних сосков (9 баллов)
	FrontNippleLocationBackView     float64 // Расположение передних сосков вид сзади (9 баллов)
	BackNippleLocationBackView      float64 // Расположение задних сосков вид сзади (9 баллов)
	FrontUdderSegmentsLocation      float64 // Прикрепление передних долей вымени (9 баллов)
	BackUdderSegmentsLocationHeight float64 // Высота прикрепления задних долей вымени (9 баллов)
	BackUdderSegmentsWidth          float64 // Ширина задних долей вымени (9 баллов)
	CentralLigamentDepth            float64 // Глубина центральной связки (9 баллов)

	// Дополнительные (8 признаков)
	BackBoneQuality     float64 // качество костяка (9 баллов)
	Deception           float64 // Обмускульность (9 баллов)
	SacrumLength        float64 // Длина крестца (9 баллов)
	TopLine             float64 // Линия верха (9 баллов)
	UdderBalance        float64 // Баланс вымени (9 баллов)
	FrontNippleDiameter float64 // Диаметр передних сосков (9 балов)
	BackNippleDiameter  float64 // Диаметр задних сосков (9 баллов)
	UdderVeins          float64 // Выраженность вен вымени (9 баллов)

	// Самодельные 9-бальные признаки
	ForeheadWidth  float64 // Ширина лба (9 баллов)
	ForeheadHeight float64 // Высота лба (9 баллов)

	//PicturePath *string
}
