package routes

import (
	"net/http"
	"distributed-pricing-engine/api/handlers/location"
	"distributed-pricing-engine/api/handlers/product"
)


func InitializePricingRoutes() {
	http.HandleFunc("/api/price", product.GetPrice)
	http.HandleFunc("/api/products", product.GetAllProducts)
	http.HandleFunc("/api/products-prices-by-location", product.GetProductsPricesByLocation)
}

func InitializeLocationRoutes() {
	http.HandleFunc("/api/location", location.GetLocation)
	http.HandleFunc("/api/random-location", location.GetRandomLocation)
}
