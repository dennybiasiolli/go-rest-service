package main

import (
	"fmt"
	"go-rest-service/views"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "HOME")
// }

// func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Articles, Category: %v\n", vars["category"])
// }

// func ArticleHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Articles, Category: %v\n", vars["category"])
// }

func main() {

	// controllers.ArticleControllers()

	router := mux.NewRouter()
	views.ProductViews(router)
	// router.PathPrefix("").Handler(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"status":  "Not Found",
	// 		"message": "message",
	// 	})
	// })
	// router.HandleFunc("/ciao", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"status":  "OK",
	// 		"message": "message",
	// 	})
	// })
	// router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusNotFound)
	// })

	// router.Use(controllers.ArticlesContrlller)

	// router.HandleFunc("/", HomeHandler)
	// router.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	// // router.HandleFunc("/products", ProductsHandler)
	// http.Handle("/", router)

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
