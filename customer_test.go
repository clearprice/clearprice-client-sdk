package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"customerId":"1","name":"John Doe","email":"john.doe@example.com"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	customer := &Customer{Name: "John Doe", Email: "john.doe@example.com"}
	createdCustomer, err := client.CreateCustomer(customer)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createdCustomer.Name != "John Doe" {
		t.Errorf("Expected customer name to be 'John Doe', got %v", createdCustomer.Name)
	}
	if createdCustomer.Email != "john.doe@example.com" {
		t.Errorf("Expected customer email to be 'john.doe@example.com', got %v", createdCustomer.Email)
	}
}

func TestGetCustomer(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"customerId":"1","name":"John Doe","email":"john.doe@example.com"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	customer, err := client.GetCustomer("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if customer.Name != "John Doe" {
		t.Errorf("Expected customer name to be 'John Doe', got %v", customer.Name)
	}
	if customer.Email != "john.doe@example.com" {
		t.Errorf("Expected customer email to be 'john.doe@example.com', got %v", customer.Email)
	}
}
