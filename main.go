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

	c := controllers.GenericController("customName")
	views.GenericView(&views.GenericViewInput{
		Router:     router,
		PathPrefix: "/products",
		Controller: &c,
	})

	views.GenericView(&views.GenericViewInput{
		Router:     router,
		PathPrefix: "articles",
	})

	views.GenericView(&views.GenericViewInput{
		Router:     router,
		PathPrefix: "/quotes",
	})

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
