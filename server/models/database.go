package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"sync"
)

var dbConnection *gorm.DB = nil
var dbInitMutex sync.Mutex

func initDatabase() error {
	dbInitMutex.Lock()
	defer dbInitMutex.Unlock()
	if dbConnection != nil {
		return nil
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Cow{}, &AdditionalInfo{}, &DownSides{}, &Ratings{}, &Exterior{}, &Measures{}, &User{}, &Weights{}); err != nil {
		panic(err)
	}
	dbConnection = db
	return nil
}

func GetDatabase() *gorm.DB {
	if dbConnection == nil {
		if err := initDatabase(); err != nil {
			panic(err)
		}
		// Get generic database object sql.DB to use its functions
		sqlDB, err := dbConnection.DB()
		if err != nil {
			panic(err)
		}

		// Sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(10)
	}
	return dbConnection
}

func MockDatabase() {
	data, err := os.ReadFile("fixture.db")
	if err != nil {
		panic(err)
	}
	// Write data to dst
	err = os.WriteFile("test.db", data, 0644)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Cow{}, &AdditionalInfo{}, &DownSides{}, &Ratings{}, &Exterior{}, &Measures{}, &User{}, &Weights{}); err != nil {
		panic(err)
	}
	dbConnection = db
}
