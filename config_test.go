package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("BASE_URL", "https://api.clearprice.xyz")
	os.Setenv("CLIENT_ID", "test_client_id")
	os.Setenv("CLIENT_SECRET", "test_client_secret")
	os.Setenv("ACCESS_TOKEN", "test_access_token")
	os.Setenv("PRODUCT_KEY", "test_product_key")
	os.Setenv("ORGANISATION_ID", "test_organisation_id")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.BaseURL != "https://api.clearprice.xyz" {
		t.Errorf("Expected BaseURL to be 'https://api.clearprice.xyz', got %v", config.BaseURL)
	}
	if config.ClientID != "test_client_id" {
		t.Errorf("Expected ClientID to be 'test_client_id', got %v", config.ClientID)
	}
}
