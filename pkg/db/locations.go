package db

import (
	"database/sql"
	"distributed-pricing-engine/pkg/models"
	"fmt"
)

func GetLocationByID(locationID int64) (*models.Location, error){
	location := &models.Location{}
	err := GetDB().QueryRow("SELECT id, country, city, tax_rate, currency, discount FROM Locations WHERE id = ?", locationID).Scan(&location.ID, &location.Country, &location.City, &location.TaxRate, &location.Currency, &location.Discount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("location does not exist")
		}
		return nil, err
	}
	return location, nil
}

func GetRandomLocation() (*models.Location, error) {
	location := &models.Location{}
	err := GetDB().QueryRow("SELECT id, country, city, tax_rate, currency, discount FROM Locations ORDER BY RAND() LIMIT 1").Scan(&location.ID, &location.Country, &location.City, &location.TaxRate, &location.Currency, &location.Discount)
	if err != nil {
		return nil, err
	}
	return location, nil
}