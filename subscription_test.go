package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSubscription(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"subscriptionId":"1","customerId":"1","planId":"1","status":"active"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	subscription := &Subscription{CustomerID: "1", PlanID: "1"}
	createdSubscription, err := client.CreateSubscription(subscription)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createdSubscription.Status != "active" {
		t.Errorf("Expected subscription status to be 'active', got %v", createdSubscription.Status)
	}
}

func TestGetSubscription(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"subscriptionId":"1","customerId":"1","planId":"1","status":"active"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	subscription, err := client.GetSubscription("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if subscription.Status != "active" {
		t.Errorf("Expected subscription status to be 'active', got %v", subscription.Status)
	}
}
