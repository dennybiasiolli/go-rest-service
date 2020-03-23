package views

import (
	"fmt"
	"go-rest-service/controllers"
	"strings"

	"github.com/gorilla/mux"
)

type GenericView struct {
	router     *mux.Router
	subrouter  *mux.Router
	pathPrefix string
	controller controllers.GenericController
}

func NewGenericView(
	router *mux.Router,
	pathPrefix string,
	controller *controllers.GenericController,
) GenericView {
	if controller == nil {
		innerController := controllers.NewGenericController(pathPrefix)
		controller = &innerController
	}
	if strings.HasPrefix(pathPrefix, "/") == false {
		pathPrefix = fmt.Sprintf("/%s", pathPrefix)
	}
	subrouter := router.PathPrefix(pathPrefix).Subrouter()
	subrouter.
		HandleFunc("/", controller.GetAll).
		Methods("GET")

	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Get).
		Methods("GET")

	subrouter.
		HandleFunc("/", controller.Post).
		Methods("POST")

	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Put).
		Methods("PUT")
	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Patch).
		Methods("PATCH")

	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Delete).
		Methods("DELETE")

	return GenericView{
		router:     router,
		subrouter:  subrouter,
		pathPrefix: pathPrefix,
		controller: *controller,
	}
}
