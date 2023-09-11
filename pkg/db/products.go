package db

import (
	"database/sql"
	"distributed-pricing-engine/pkg/models"
	"errors"
	"log"
)

func GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product

	query := "SELECT product_id, product_name, base_price FROM Products"
	rows, err := GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.BasePrice)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func GetProductByID(productID int) (*models.Product, error) {
	var product models.Product

	query := "SELECT product_id, product_name, base_price FROM Products WHERE product_id = ?"
	err := GetDB().QueryRow(query, productID).Scan(&product.ID, &product.Name, &product.BasePrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product does not exist")
		}
		return nil, err
	}
	
	return &product, nil
}

func CalculateFinalPrice(basePrice float64, taxRate float64, discount float64) float64 {
    // Calculating the amount of tax
    taxAmount := basePrice * (taxRate / 100)
    
    // Calculating the discount amount
    discountAmount := basePrice * (discount / 100)

    // Final price = base price + tax - discount
    finalPrice := basePrice + taxAmount - discountAmount
    
    return finalPrice
}


// Write function to get all products and based on given location, calculate final price for each product
func GetProductsPricesByLocation(locationID int64) ([]*models.Product, error) {
	var products []*models.Product

	location, err := GetLocationByID(locationID)
	if err != nil {
		return nil, err
	}

	query := "SELECT product_id, product_name, base_price FROM Products"
	rows, err := GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.BasePrice)
		if err != nil {
			log.Printf("Failed to scan product: %v", err)
			continue
		}

		product.FinalPrice = CalculateFinalPrice(product.BasePrice, location.TaxRate, location.Discount)
		products = append(products, &product)
	}


	return products, nil
}