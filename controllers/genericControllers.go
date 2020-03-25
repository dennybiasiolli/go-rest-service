package controllers

import (
	"fmt"
	"go-rest-service/utils"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type GenericControllerOutput struct {
	ControllerName string
	GetAll         func(w http.ResponseWriter, r *http.Request)
	Get            func(w http.ResponseWriter, r *http.Request)
	Post           func(w http.ResponseWriter, r *http.Request)
	Put            func(w http.ResponseWriter, r *http.Request)
	Patch          func(w http.ResponseWriter, r *http.Request)
	Delete         func(w http.ResponseWriter, r *http.Request)
}

func GenericController(controllerName string) GenericControllerOutput {
	if strings.HasPrefix(controllerName, "/") {
		controllerName = controllerName[1:]
	}
	return GenericControllerOutput{
		ControllerName: controllerName,

		GetAll: func(w http.ResponseWriter, r *http.Request) {
			utils.Respond(w, utils.Message(
				true,
				fmt.Sprintf("GetAll %s", controllerName),
			))
		},

		Get: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			utils.Respond(w, utils.Message(
				true,
				fmt.Sprintf("Get %s %s", controllerName, vars["id"]),
			))
		},

		Post: func(w http.ResponseWriter, r *http.Request) {
			utils.Respond(w, utils.Message(
				true,
				fmt.Sprintf("Post %s", controllerName),
			))
		},

		Put: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			utils.Respond(w, utils.Message(
				true,
				fmt.Sprintf("Put %s %s", controllerName, vars["id"]),
			))
		},

		Patch: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			utils.Respond(w, utils.Message(
				true,
				fmt.Sprintf("Patch %s %s", controllerName, vars["id"]),
			))
		},

		Delete: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			utils.Respond(w, utils.Message(
				true,
				fmt.Sprintf("Delete %s %s", controllerName, vars["id"]),
			))
		},
	}
}
