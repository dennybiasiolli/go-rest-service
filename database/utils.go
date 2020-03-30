package database

import (
	"go-rest-service/settings"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var IsDbInitialized = false

func InitDbConn(fnDbAutoMigrate func(db *gorm.DB)) {
	log.Println("Opening db conn...")
	db, err = gorm.Open(settings.DatabaseDialect, settings.DatabaseConnectionString)
	if err != nil {
		log.Fatalln(err)
		panic("failed to connect database")
	}
	fnDbAutoMigrate(db)
}

func CloseDbConn() {
	db.Close()
	log.Println("Closing db conn...")
}

func DbOperation(fn func(db *gorm.DB)) {
	fn(db)
}
