package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlans(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"planId":"1","name":"Plan 1","price":10.0,"features":["Feature 1"]}]`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	plans, err := client.GetPlans()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(plans) != 1 {
		t.Errorf("Expected 1 plan, got %v", len(plans))
	}
	if plans[0].Name != "Plan 1" {
		t.Errorf("Expected plan name to be 'Plan 1', got %v", plans[0].Name)
	}
}

func TestGetPlan(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"planId":"1","name":"Plan 1","price":10.0,"features":["Feature 1"]}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, "test_access_token", "test_product_key", "test_organisation_id")

	plan, err := client.GetPlan("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if plan.Name != "Plan 1" {
		t.Errorf("Expected plan name to be 'Plan 1', got %v", plan.Name)
	}
}
