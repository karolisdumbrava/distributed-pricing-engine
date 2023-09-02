package product

import (
	"net/http"
	"encoding/json"
	"distributed-pricing-engine/pkg/db"
	"strconv"
	"log"
)

// Product functions

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching all products")
	json.NewEncoder(w).Encode(products)
}

func GetProductsPricesByLocation(w http.ResponseWriter, r *http.Request) {
	locationIDStr := r.URL.Query().Get("location_id")
	locationID, err := strconv.ParseInt(locationIDStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	products, err := db.GetProductsPricesByLocation(locationID)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching products with location ID: ", locationID)
	json.NewEncoder(w).Encode(products)
}

// Price functions

func GetPrice(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	productID, err := strconv.Atoi(productIDStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := db.GetProductByID(productID)
	if err != nil {
		if err.Error() == "product does not exist" {
			http.Error(w, "Product does not exist", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		}
		return
	}

	product.FinalPrice = product.BasePrice

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching product with ID: ", productID)
	json.NewEncoder(w).Encode(product)
}

