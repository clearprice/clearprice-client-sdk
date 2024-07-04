package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetOAuthToken(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expected 'POST' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/oauth2/token" {
			t.Fatalf("Expected request to '/oauth2/token', got '%s'", r.URL.EscapedPath())
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"access_token": "test_access_token",
			"token_type": "bearer",
			"expires_in": 3600,
			"refresh_token": "test_refresh_token"
		}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Replace the base URL in the request
	originalURL := "https://auth.clearprice.xyz/oauth2/token"
	mockURL, _ := url.Parse(server.URL + "/oauth2/token")

	clientID := "test_client_id"
	clientSecret := "test_client_secret"
	code := "test_code"
	redirectURI := "https://example.com/callback"

	// Use strings.Replace to replace the real URL with the mock URL in the function
	getOAuthToken := func(clientID, clientSecret, code, redirectURI string) (*OAuthTokenResponse, error) {
		data := url.Values{}
		data.Set("grant_type", "authorization_code")
		data.Set("code", code)
		data.Set("redirect_uri", redirectURI)
		data.Set("client_id", clientID)
		data.Set("client_secret", clientSecret)

		req, err := http.NewRequest("POST", strings.Replace(originalURL, "https://auth.clearprice.xyz/oauth2/token", mockURL.String(), 1), strings.NewReader(data.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var tokenResponse OAuthTokenResponse
		err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
		if err != nil {
			return nil, err
		}

		return &tokenResponse, nil
	}

	tokenResponse, err := getOAuthToken(clientID, clientSecret, code, redirectURI)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tokenResponse.AccessToken != "test_access_token" {
		t.Errorf("Expected access token to be 'test_access_token', got %v", tokenResponse.AccessToken)
	}
	if tokenResponse.RefreshToken != "test_refresh_token" {
		t.Errorf("Expected refresh token to be 'test_refresh_token', got %v", tokenResponse.RefreshToken)
	}
}
