package main

import (
	"fmt"
	"log"
	"distributed-pricing-engine/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.GetDB().Close()

	fmt.Println("Database connection successful")
}