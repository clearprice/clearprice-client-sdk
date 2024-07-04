package main

type Subscription struct {
	ID         string `json:"subscriptionId"`
	CustomerID string `json:"customerId"`
	PlanID     string `json:"planId"`
	Status     string `json:"status"`
}

func (c *Client) CreateSubscription(subscription *Subscription) (*Subscription, error) {
	req, err := c.NewRequest("POST", "/subscriptions", subscription)
	if err != nil {
		return nil, err
	}

	var createdSubscription Subscription
	err = c.Do(req, &createdSubscription)
	if err != nil {
		return nil, err
	}

	return &createdSubscription, nil
}

func (c *Client) GetSubscription(subscriptionID string) (*Subscription, error) {
	req, err := c.NewRequest("GET", "/subscriptions/"+subscriptionID, nil)
	if err != nil {
		return nil, err
	}

	var subscription Subscription
	err = c.Do(req, &subscription)
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (c *Client) UpdateSubscription(subscriptionID string, subscription *Subscription) (*Subscription, error) {
	req, err := c.NewRequest("PUT", "/subscriptions/"+subscriptionID, subscription)
	if err != nil {
		return nil, err
	}

	var updatedSubscription Subscription
	err = c.Do(req, &updatedSubscription)
	if err != nil {
		return nil, err
	}

	return &updatedSubscription, nil
}

func (c *Client) DeleteSubscription(subscriptionID string) error {
	req, err := c.NewRequest("DELETE", "/subscriptions/"+subscriptionID, nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}
