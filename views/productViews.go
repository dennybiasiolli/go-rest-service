package views

import (
	"go-rest-service/controllers"

	"github.com/gorilla/mux"
)

func ProductViews(router *mux.Router) {
	productRoute := router.PathPrefix("/products").Subrouter()

	productRoute.
		HandleFunc("/", controllers.ProductControllerGetAll).
		Methods("GET")

	productRoute.
		HandleFunc("/{id:[0-9]+}/", controllers.ProductControllerGet).
		Methods("GET")

	productRoute.
		HandleFunc("/", controllers.ProductControllerGetAll).
		Methods("POST")

	productRoute.
		HandleFunc("/{id:[0-9]+}/", controllers.ProductControllerPut).
		Methods("PUT")
	productRoute.
		HandleFunc("/{id:[0-9]+}/", controllers.ProductControllerPatch).
		Methods("PATCH")
	productRoute.
		HandleFunc("/{id:[0-9]+}/", controllers.ProductControllerDelete).
		Methods("DELETE")

	// productRoute.HandleFunc("/ciao", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"status":  "OK",
	// 		"message": "message",
	// 	})
	// })
}
