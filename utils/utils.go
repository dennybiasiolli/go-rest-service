package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DbOperation(fn func(db *gorm.DB)) {
	log.Println("Opening db conn...")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
		panic("failed to connect database")
	}
	defer db.Close()

	fn(db)
	defer log.Println("Closing db conn...")
}
