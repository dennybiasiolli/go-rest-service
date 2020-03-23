package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductControllerGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": "ProductControllerGetAll",
	})
}

func ProductControllerGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": fmt.Sprintf("ProductControllerGet %s", vars["id"]),
	})
}

func ProductControllerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": "ProductControllerPost",
	})
}

func ProductControllerPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": fmt.Sprintf("ProductControllerPut %s", vars["id"]),
	})
}
func ProductControllerPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": fmt.Sprintf("ProductControllerPatch %s", vars["id"]),
	})
}

func ProductControllerDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": fmt.Sprintf("ProductControllerDelete %s", vars["id"]),
	})
}
