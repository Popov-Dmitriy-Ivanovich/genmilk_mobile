package models

type Cow struct {
	ID              uint     `gorm:"primaryKey" json:"-"`
	InventoryNumber *string   `example:"1213321" ` // Инвентарный номер коровы
	SelecsNumber    *uint64   `example:"98989" `   // Селекс номер коровы
	RSHNNumber      *string   `example:"1323323232" `                 // РСХН номер коровы
	Name            *string   `example:"Дима" `                       // Кличка коровы
	BreedName       *string   `example:"Порода" `  // Название породы
	Bloody          *float64 `example:"1.00" `                       // Кровность

	Defects 			*string // пороки (если несколько, то разделитель /)

	BirthDate           *DateOnly   `json:",string" example:"2001-03-23" swaggertype:"string" ` // Дата рождения ГГГГ-ММ-ЧЧ
	HoldingName         *string     ``                                       // Название холдинга
	HoldingInn          *string     ``                                       // ИНН холдинга
	HozName             *string     ``                                       // Название хозяйства
	Exteriors           []Exterior // Оценки экстерьера
	FirstAssessmentDate *DateOnly  `example:"2001-03-23" swaggertype:"string"`
}

//TODO: Добавить опциональную строку порок(defect) NULL - нет порока, иначе нет информации о корове
