package main

import (
	"context"
	"flag"
	"fmt"
	"go-rest-service/controllers"
	"go-rest-service/middlewares"
	"go-rest-service/models"
	"go-rest-service/utils"
	"go-rest-service/views"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	utils.DbOperation(func(db *gorm.DB) {
		// Migrate the schema
		db.AutoMigrate(
			&models.Product{},
			&models.ResUser{},
		)

		// Create
		db.Create(&models.Product{Code: "L1212", Price: 123})

		// Read
		var product models.Product
		var products []models.Product
		db.First(&product, 1)                   // find product with id 1
		db.First(&product, "code = ?", "L1212") // find product with code l1212
		db.Find(&products, "price > ?", 10)     // find products with price > 10

		// Update - update product's price to 2000
		db.Model(&product).Update("Price", 2000)

		// Delete - delete product
		db.Delete(&product)
	})

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
	// enabling CORS
	router.Methods(http.MethodOptions)
	router.Use(middlewares.CORSMiddleware)

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
