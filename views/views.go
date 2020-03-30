package views

import (
	"go-rest-service/models"
	"go-rest-service/routing"

	"github.com/gorilla/mux"
)

func SetViews(router *mux.Router) {
	routing.View(&routing.ViewInput{
		Router:     router,
		PathPrefix: "/products",
		ModelPtr:   &models.Product{},
	})

	routing.View(&routing.ViewInput{
		Router:     router,
		PathPrefix: "/res-users",
		ModelPtr:   &models.ResUser{},
	})
}
