package models

type User struct {
	ID                    uint    `gorm:"primaryKey" json:"-"`
	NameSurnamePatronymic string  // ФИО
	Email                 string  `gorm:"uniqueIndex"` // Почта
	LicenceNumber         *string `gorm:"index"`       // Номер лицензии бонетера
	Password              []byte  `json:"-"`
}
