package models

type Measures struct {
	ID         uint `gorm:"primaryKey" json:"-"`
	ExteriorID uint `json:"-" gorm:"index"`

	SacrumHeight float64 `binding:"required"` // высота в крестце (Сантиметры или градусы)
	ChestWidth   float64 `binding:"required"` // Ширина груди (Сантиметры или градусы)
	SacrumAngle  float64 `binding:"required"` // Угол наклона крестца (Сантиметры или градусы)
	SacrumWidth  float64 `binding:"required"` // ширина в крестце (Сантиметры или градусы)

	HindLegWalkSideView float64 `binding:"required"` // Поступь задних ног сбоку (Сантиметры или градусы)
	HoofAngle           float64 `binding:"required"` // Угол копыта (Сантиметры или градусы)

	UdderDepth                      float64 `binding:"required"` // Глубина вымени (Сантиметры или градусы)
	FrontUdderSegmentsLocation      float64 `binding:"required"` // Прикрепление передних долей вымени (Сантиметры или градусы)
	BackUdderSegmentsLocationHeight float64 `binding:"required"` // Высота прикрепления задних долей вымени (Сантиметры или градусы)
	BackUdderSegmentsWidth          float64 `binding:"required"` // Ширина задних долей вымени (Сантиметры или градусы)
	CentralLigamentDepth            float64 `binding:"required"` // Глубина центральной связки (Сантиметры или градусы)
	FrontNippleLength               float64 `binding:"required"` // Длинна передних сосков (Сантиметры или градусы)

	UdderBalance        float64 `binding:"required"` // Баланс вымени (Сантиметры или градусы)
	FrontNippleDiameter float64 `binding:"required"` // Диаметр передних сосков (Сантиметры или градусы)
	BackNippleDiameter  float64 `binding:"required"` // Диаметр задних сосков (Сантиметры или градусы)
}
