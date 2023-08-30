package db

import (
	"distributed-pricing-engine/models"
	"errors"
	"database/sql"
)

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