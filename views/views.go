package views

import (
	"go-rest-service/models"

	"github.com/dennybiasiolli/gorestframework"
	"github.com/gorilla/mux"
)

func SetViews(router *mux.Router) {
	gorestframework.View(&gorestframework.ViewInput{
		Router:     router,
		PathPrefix: "/products",
		ModelPtr:   &models.Product{},
	})

	gorestframework.View(&gorestframework.ViewInput{
		Router:     router,
		PathPrefix: "/res-users",
		ModelPtr:   &models.ResUser{},
	})
}
