package main

type Entitlement struct {
	ID          string `json:"entitlementId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EntitlementUsage struct {
	CustomerID string `json:"customerId"`
	UsageCount int    `json:"usageCount"`
}

func (c *Client) GetCustomerEntitlements(customerID string) ([]Entitlement, error) {
	req, err := c.NewRequest("GET", "/customers/"+customerID+"/entitlements", nil)
	if err != nil {
		return nil, err
	}

	var entitlements []Entitlement
	err = c.Do(req, &entitlements)
	if err != nil {
		return nil, err
	}

	return entitlements, nil
}

func (c *Client) LogEntitlementUsage(entitlementID string, usage *EntitlementUsage) (string, error) {
	req, err := c.NewRequest("POST", "/entitlements/"+entitlementID+"/log", usage)
	if err != nil {
		return "", err
	}

	var result map[string]string
	err = c.Do(req, &result)
	if err != nil {
		return "", err
	}

	return result["message"], nil
}
