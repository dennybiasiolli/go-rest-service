package main

import (
	"context"
	"flag"
	"fmt"
	"go-rest-service/controllers"
	"go-rest-service/middlewares"
	"go-rest-service/views"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m",
	)
	flag.Parse()

	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddleware)

	controller := controllers.GenericController("customName")
	views.GenericView(&views.GenericViewInput{
		Router:     router,
		PathPrefix: "/products",
		Controller: &controller,
	})

	views.GenericView(&views.GenericViewInput{
		Router:     router,
		PathPrefix: "articles",
	})

	views.GenericView(&views.GenericViewInput{
		Router:     router,
		PathPrefix: "/quotes",
	})

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", host, port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Starting http.Server on", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Shutting down")
	os.Exit(0)
}
