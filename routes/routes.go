package routes

import (
	"net/http"
	"distributed-pricing-engine/handlers"
)


func InitializePricingRoutes() {
	http.HandleFunc("/api/price", handlers.GetPrice)
}

func InitializeLocationRoutes() {
	http.HandleFunc("/api/location", handlers.GetLocation)
}
