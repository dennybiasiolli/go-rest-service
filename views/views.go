package views

import (
	"go-rest-service/models"
	"go-rest-service/utils"

	"github.com/gorilla/mux"
)

func SetViews(router *mux.Router) {
	utils.GenericView(&utils.GenericViewInput{
		Router:     router,
		PathPrefix: "/products",
		ModelPtr:   &models.Product{},
	})

	utils.GenericView(&utils.GenericViewInput{
		Router:     router,
		PathPrefix: "/res-users",
		ModelPtr:   &models.ResUser{},
	})
}
