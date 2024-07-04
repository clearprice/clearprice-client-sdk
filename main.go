package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := NewClient(config.BaseURL, config.AccessToken, config.ProductKey, config.OrganisationID)

	// Example: Create a new customer
	newCustomer := &Customer{Name: "John Doe", Email: "john.doe@example.com"}
	createdCustomer, err := client.CreateCustomer(newCustomer)
	if err != nil {
		log.Fatalf("Error creating customer: %v", err)
	}
	fmt.Printf("Created Customer: %+v\n", createdCustomer)

	// Example: Get customer details
	customer, err := client.GetCustomer(createdCustomer.ID)
	if err != nil {
		log.Fatalf("Error getting customer: %v", err)
	}
	fmt.Printf("Customer Details: %+v\n", customer)
}
