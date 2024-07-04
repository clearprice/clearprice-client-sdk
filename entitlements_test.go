package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomerEntitlements(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"entitlementId":"1","name":"Entitlement 1","description":"Description 1"}]`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	entitlements, err := client.GetCustomerEntitlements("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(entitlements) != 1 {
		t.Errorf("Expected 1 entitlement, got %v", len(entitlements))
	}
	if entitlements[0].Name != "Entitlement 1" {
		t.Errorf("Expected entitlement name to be 'Entitlement 1', got %v", entitlements[0].Name)
	}
}

func TestLogEntitlementUsage(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Entitlement usage logged successfully"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	usage := &EntitlementUsage{CustomerID: "1", UsageCount: 10}
	message, err := client.LogEntitlementUsage("1", usage)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedMessage := "Entitlement usage logged successfully"
	if message != expectedMessage {
		t.Errorf("Expected message to be '%s', got '%s'", expectedMessage, message)
	}
}
