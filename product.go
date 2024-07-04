package main

type Product struct {
	ID       string `json:"productId"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Industry string `json:"industry"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Currency string `json:"currency"`
}

func (c *Client) GetProducts() ([]Product, error) {
	req, err := c.NewRequest("GET", "/products", nil)
	if err != nil {
		return nil, err
	}

	var products []Product
	err = c.Do(req, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c *Client) GetProduct(productID string) (*Product, error) {
	req, err := c.NewRequest("GET", "/products/"+productID, nil)
	if err != nil {
		return nil, err
	}

	var product Product
	err = c.Do(req, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
