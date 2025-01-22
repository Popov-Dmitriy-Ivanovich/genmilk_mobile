package models

type Cow struct {
	ID              uint    `gorm:"primaryKey" json:"-"`
	InventoryNumber string  `example:"1213321" binding:"required"`    // Инвентарный номер коровы
	SelecsNumber    uint64  `example:"98989" binding:"required"`      // Селекс номер коровы
	RSHNNumber      string  `example:"1323323232" binding:"required"` // РСХН номер коровы
	Name            string  `example:"Дима" binding:"required"`       // Кличка коровы
	BreedName       string  `example:"Порода" binding:"required"`     // Название породы
	Bloody          float64 `example:"1.00" binding:"required"`       // Кровность

	BirthDate   DateOnly   `json:",string" example:"2001-03-23" swaggertype:"string" binding:"required"` // Дата рождения ГГГГ-ММ-ЧЧ
	HoldingName string     `binding:"required"`                                                          // Название холдинга
	HoldingInn  string     `binding:"required"`                                                          // ИНН холдинга
	HozName     string     `binding:"required"`                                                          // Название хозяйства
	Exteriors   []Exterior `json:"-"`                                                                    // Оценки экстерьера
}
