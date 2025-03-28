package models

type Cow struct {
	ID              uint     `gorm:"primaryKey" json:"-"`
	InventoryNumber string   `example:"1213321" validate:"required"` // Инвентарный номер коровы
	SelecsNumber    uint64   `example:"98989" validate:"required"`   // Селекс номер коровы
	RSHNNumber      string   `example:"1323323232" `                 // РСХН номер коровы
	Name            string   `example:"Дима" `                       // Кличка коровы
	BreedName       string   `example:"Порода" validate:"required"`  // Название породы
	Bloody          *float64 `example:"1.00" `                       // Кровность

	BirthDate           DateOnly   `json:",string" example:"2001-03-23" swaggertype:"string" ` // Дата рождения ГГГГ-ММ-ЧЧ
	HoldingName         string     `validate:"required"`                                       // Название холдинга
	HoldingInn          string     `validate:"required"`                                       // ИНН холдинга
	HozName             string     `validate:"required"`                                       // Название хозяйства
	Exteriors           []Exterior // Оценки экстерьера
	FirstAssessmentDate *DateOnly  `example:"2001-03-23" swaggertype:"string"`
}
