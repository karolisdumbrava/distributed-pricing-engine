package location

import (
	"distributed-pricing-engine/pkg/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetLocation(w http.ResponseWriter, r *http.Request) {
	locationIDStr := r.URL.Query().Get("location_id")
	locationID, err := strconv.ParseInt(locationIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	location, err := db.GetLocationByID(locationID)
	if err != nil {
		if err.Error() == "location does not exist" {
			http.Error(w, "Location does not exist", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch location", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching location with ID: ", locationID)
	json.NewEncoder(w).Encode(location)
}

func GetRandomLocation(w http.ResponseWriter, r *http.Request) {
	location, err := db.GetRandomLocation()
	if err != nil {
		http.Error(w, "Failed to fetch location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching random location")
	json.NewEncoder(w).Encode(location)
}