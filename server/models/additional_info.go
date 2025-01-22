package models

type AdditionalInfo struct {
	ID         uint `gorm:"primary_key" json:"-"`
	ExteriorID uint `json:"-" gorm:"index"`

	AdditionalProperty1Name *string // Дополнительный параметр 1 (название)
	AdditionalProperty2Name *string // Дополнительный параметр 2 (название)

	AdditionalProperty1Value *string // Дополнительный параметр 1 (значение в баллах)
	AdditionalProperty2Value *string // Дополнительный параметр 2 (значение в баллах)

	AdditionalProperty1Measure *string // Дополнительный параметр 1 (значение в единицах измерения)
	AdditionalProperty2Measure *string // Дополнительный параметр 2 (значение в единицах измерения)

	LactationNumber  uint     // Номер лактации (целое, беззнаковое число)
	CalvingDate      DateOnly `json:",string" example:"2001-03-23" swaggertype:"string" binding:"required"` // Дата отела ГГГГ-ММ-ДД
	FirstMilkingDate DateOnly `json:",string" example:"2001-03-23" swaggertype:"string" binding:"required"` // Дата первого доения ГГГГ-ММ-ДД
}
