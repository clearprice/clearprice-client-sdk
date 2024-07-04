package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	baseURL := "https://api.clearprice.xyz"
	accessToken := "test_access_token"
	productKey := "test_product_key"
	organisationID := "test_organisation_id"

	client := NewClient(baseURL, accessToken, productKey, organisationID)

	if client.BaseURL != baseURL {
		t.Errorf("Expected baseURL to be %v, got %v", baseURL, client.BaseURL)
	}
	if client.AccessToken != accessToken {
		t.Errorf("Expected accessToken to be %v, got %v", accessToken, client.AccessToken)
	}
}

func TestDo(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"success"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	req, err := client.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var result map[string]string
	err = client.Do(req, &result)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result["message"] != "success" {
		t.Errorf("Expected message to be 'success', got %v", result["message"])
	}
}
