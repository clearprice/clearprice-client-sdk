package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"productId":"1","name":"Product 1","category":"Category 1","industry":"Industry 1","country":"Country 1","region":"Region 1","currency":"USD"}]`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	products, err := client.GetProducts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %v", len(products))
	}
	if products[0].Name != "Product 1" {
		t.Errorf("Expected product name to be 'Product 1', got %v", products[0].Name)
	}
}

func TestGetProduct(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"productId":"1","name":"Product 1","category":"Category 1","industry":"Industry 1","country":"Country 1","region":"Region 1","currency":"USD"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	product, err := client.GetProduct("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if product.Name != "Product 1" {
		t.Errorf("Expected product name to be 'Product 1', got %v", product.Name)
	}
}
