package views

import (
	"fmt"
	"go-rest-service/controllers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type GenericViewInput struct {
	Router     *mux.Router
	PathPrefix string
	Controller *controllers.GenericControllerOutput
}

type GenericViewOutput struct {
	Router     *mux.Router
	Subrouter  *mux.Router
	PathPrefix string
	Controller controllers.GenericControllerOutput
}

func GenericView(
	input *GenericViewInput,
) GenericViewOutput {
	var controller controllers.GenericControllerOutput
	if input.Controller == nil {
		controller = controllers.GenericController(input.PathPrefix)
	} else {
		controller = *input.Controller
	}
	if strings.HasPrefix(input.PathPrefix, "/") == false {
		input.PathPrefix = fmt.Sprintf("/%s", input.PathPrefix)
	}
	subrouter := input.Router.PathPrefix(input.PathPrefix).Subrouter()
	subrouter.
		HandleFunc("/", controller.GetAll).
		Methods(http.MethodGet)

	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Get).
		Methods(http.MethodGet)

	subrouter.
		HandleFunc("/", controller.Post).
		Methods(http.MethodPost)

	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Put).
		Methods(http.MethodPut)
	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Patch).
		Methods(http.MethodPatch)

	subrouter.
		HandleFunc("/{id:[0-9]+}/", controller.Delete).
		Methods(http.MethodDelete)

	return GenericViewOutput{
		Router:     input.Router,
		Subrouter:  subrouter,
		PathPrefix: input.PathPrefix,
		Controller: controller,
	}
}
