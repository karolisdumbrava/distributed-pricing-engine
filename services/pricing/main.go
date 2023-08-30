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

	// Initialize routes
	routes.InitializePricingRoutes()

	log.Println("Pricing Service listening on :8080")
	http.ListenAndServe(":8080", nil)
}
