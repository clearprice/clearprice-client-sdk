package main

type Plan struct {
	ID       string   `json:"planId"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Features []string `json:"features"`
}

func (c *Client) GetPlans() ([]Plan, error) {
	req, err := c.NewRequest("GET", "/plans", nil)
	if err != nil {
		return nil, err
	}

	var plans []Plan
	err = c.Do(req, &plans)
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (c *Client) GetPlan(planID string) (*Plan, error) {
	req, err := c.NewRequest("GET", "/plans/"+planID, nil)
	if err != nil {
		return nil, err
	}

	var plan Plan
	err = c.Do(req, &plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}
