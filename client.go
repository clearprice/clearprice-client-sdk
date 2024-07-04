package main

import (
	"net/http"
)

type Client struct {
	BaseURL        string
	HTTPClient     *http.Client
	AccessToken    string
	ProductKey     string
	OrganisationID string
}

func NewClient(baseURL, accessToken, productKey, organisationID string) *Client {
	return &Client{
		BaseURL:        baseURL,
		HTTPClient:     &http.Client{},
		AccessToken:    accessToken,
		ProductKey:     productKey,
		OrganisationID: organisationID,
	}
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, encodeBody(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("X-Product-Key", c.ProductKey)
	req.Header.Set("X-Organisation-ID", c.OrganisationID)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeBody(resp.Body, v)
}
