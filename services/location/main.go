package main

import (
	"log"
	"net/http"
	"distributed-pricing-engine/db"
	"distributed-pricing-engine/routes"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.GetDB().Close()

	routes.InitializeLocationRoutes()


	log.Println("Location Service listening on :8081")
	http.ListenAndServe(":8081", nil)
}