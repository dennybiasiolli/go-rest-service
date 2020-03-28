package utils

import (
	"encoding/json"
	"go-rest-service/settings"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func JsonRespond(w http.ResponseWriter, data interface{}) {
	JsonRespondWithStatus(w, data, http.StatusOK)
}
func JsonRespondWithStatus(w http.ResponseWriter, data interface{}, httpStatus int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}

func DbOperation(fn func(db *gorm.DB)) {
	log.Println("Opening db conn...")
	db, err := gorm.Open(settings.DatabaseDialect, settings.DatabaseConnectionString)
	if err != nil {
		log.Fatalln(err)
		panic("failed to connect database")
	}
	defer db.Close()

	fn(db)
	defer log.Println("Closing db conn...")
}
