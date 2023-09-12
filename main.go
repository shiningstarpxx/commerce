package main

import (
	"log"
	"net/http"

	"commerce/controllers"
	"commerce/database"
	"commerce/datalayer"
	"commerce/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func main() {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Pass the database connection to the controller
	daatalayer := &datalayer.UserDatalayer{DB: db}
	userController := controllers.UserController{Datalayer: daatalayer}

	// Setup Metrics Server
	initPrometheus()

	// Setup the router and routes
	router := router.SetupRoutes(userController)

	// Start the server
	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
