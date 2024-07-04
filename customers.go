package main

type Customer struct {
	ID     string `json:"customerId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func (c *Client) CreateCustomer(customer *Customer) (*Customer, error) {
	req, err := c.NewRequest("POST", "/customers", customer)
	if err != nil {
		return nil, err
	}

	var createdCustomer Customer
	err = c.Do(req, &createdCustomer)
	if err != nil {
		return nil, err
	}

	return &createdCustomer, nil
}

func (c *Client) GetCustomer(customerID string) (*Customer, error) {
	req, err := c.NewRequest("GET", "/customers/"+customerID, nil)
	if err != nil {
		return nil, err
	}

	var customer Customer
	err = c.Do(req, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *Client) UpdateCustomer(customerID string, customer *Customer) (*Customer, error) {
	req, err := c.NewRequest("PUT", "/customers/"+customerID, customer)
	if err != nil {
		return nil, err
	}

	var updatedCustomer Customer
	err = c.Do(req, &updatedCustomer)
	if err != nil {
		return nil, err
	}

	return &updatedCustomer, nil
}

func (c *Client) DeleteCustomer(customerID string) error {
	req, err := c.NewRequest("DELETE", "/customers/"+customerID, nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}
