package main

import (
	"fmt"
	"os"
)

type Config struct {
	BaseURL        string
	ClientID       string
	ClientSecret   string
	AccessToken    string
	ProductKey     string
	OrganisationID string
}

func LoadConfig() (*Config, error) {
	baseURL := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	productKey := os.Getenv("PRODUCT_KEY")
	organisationID := os.Getenv("ORGANISATION_ID")

	if baseURL == "" || clientID == "" || clientSecret == "" || accessToken == "" || productKey == "" || organisationID == "" {
		return nil, fmt.Errorf("missing environment variables")
	}

	return &Config{
		BaseURL:        baseURL,
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		AccessToken:    accessToken,
		ProductKey:     productKey,
		OrganisationID: organisationID,
	}, nil
}
