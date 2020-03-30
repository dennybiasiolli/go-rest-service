package main

import (
	"go-rest-service/database"
	"go-rest-service/models"
	"go-rest-service/routing"
	"go-rest-service/views"
)

func main() {
	database.InitDbConn(models.MigrateModels)
	defer database.CloseDbConn()

	routing.StartHTTPListener(views.SetViews)
}
