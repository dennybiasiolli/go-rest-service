package main

import (
	"go-rest-service/models"
	"go-rest-service/settings"
	"go-rest-service/views"

	"github.com/dennybiasiolli/gorestframework"
)

func main() {
	gorestframework.InitDbConn(
		settings.DatabaseDialect,
		settings.DatabaseConnectionString,
		models.MigrateModels,
	)
	defer gorestframework.CloseDbConn()

	gorestframework.StartHTTPListener(
		settings.RouterActivateLog,
		settings.RouterUseCORS,
		views.SetViews,
	)
}
