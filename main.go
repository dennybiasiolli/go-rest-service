package main

import (
	"fmt"
	"go-rest-service/controllers"
	"go-rest-service/views"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	c := controllers.NewGenericController("customName")
	views.NewGenericView(router, "/products", &c)
	views.NewGenericView(router, "articles", nil)
	views.NewGenericView(router, "/quotes", nil)

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
