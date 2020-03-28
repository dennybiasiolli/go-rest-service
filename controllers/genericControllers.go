package controllers

import (
	"fmt"
	"go-rest-service/utils"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type GenericControllerOutput struct {
	ControllerName string
	ModelPtr       interface{}
	GetAll         func(w http.ResponseWriter, r *http.Request)
	Get            func(w http.ResponseWriter, r *http.Request)
	Post           func(w http.ResponseWriter, r *http.Request)
	Put            func(w http.ResponseWriter, r *http.Request)
	Patch          func(w http.ResponseWriter, r *http.Request)
	Delete         func(w http.ResponseWriter, r *http.Request)
}

func GenericController(controllerName string, modelPtr interface{}) GenericControllerOutput {
	if strings.HasPrefix(controllerName, "/") {
		controllerName = controllerName[1:]
	}
	T := reflect.TypeOf(modelPtr)
	if T.Kind() == reflect.Ptr {
		T = T.Elem()
	}
	return GenericControllerOutput{
		ControllerName: controllerName,
		ModelPtr:       modelPtr,

		GetAll: func(w http.ResponseWriter, r *http.Request) {
			utils.DbOperation(func(db *gorm.DB) {
				resultsPtr := reflect.New(reflect.SliceOf(T)).Interface()
				db.Find(resultsPtr)
				utils.JsonRespond(w, resultsPtr)
			})
		},

		Get: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			utils.DbOperation(func(db *gorm.DB) {
				resultPtr := reflect.New(T).Interface()
				if db.First(resultPtr, vars["id"]).RecordNotFound() {
					utils.JsonRespondWithStatus(w, map[string]interface{}{
						"error":   true,
						"message": fmt.Sprintf("%s with ID %s not found", controllerName, vars["id"]),
					}, http.StatusNotFound)
				} else {
					utils.JsonRespond(w, resultPtr)
				}
			})
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
