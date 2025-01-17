package models

type Cow struct {
	ID              uint       `gorm:"primaryKey" json:"-"`
	InventoryNumber string     `example:"1213321"`                                        // Инвентарный номер коровы
	SelecsNumber    uint64     `example:"98989"`                                          // Селекс номер коровы
	RSHNNumber      string     `example:"1323323232"`                                     // РСХН номер коровы
	Name            string     `example:"Дима"`                                           // Кличка коровы
	BreedName       string     `example:"Порода"`                                         // Название породы
	Bloody          float64    `example:"1.00"`                                           // Кровность
	BirthDate       DateOnly   `json:",string" example:"2001-03-23" swaggertype:"string"` // Дата рождения ГГГГ-ММ-ЧЧ
	Exteriors       []Exterior `json:"-"`                                                 // Оценки экстерьера
}