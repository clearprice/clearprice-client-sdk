
# Official ClearPrice Client SDK for Golang

The ClearPrice Client SDK is a Go library for interacting with the ClearPrice API. It provides convenient methods for authenticating and performing operations related to products, plans, customers, subscriptions, entitlements, and logging entitlement usage.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Authentication](#authentication)
  - [Products](#products)
  - [Plans](#plans)
  - [Customers](#customers)
  - [Subscriptions](#subscriptions)
  - [Entitlements](#entitlements)
  - [Logging Entitlement Usage](#logging-entitlement-usage)
- [Running Tests](#running-tests)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install the ClearPrice Client SDK, run:

```sh
go get github.com/clearprice/clearprice-client-sdk
```

## Configuration

Set the following environment variables to configure the SDK:

```sh
export BASE_URL=https://api.clearprice.xyz
export CLIENT_ID=your_client_id
export CLIENT_SECRET=your_client_secret
export ACCESS_TOKEN=your_access_token
export PRODUCT_KEY=your_product_key
export ORGANISATION_ID=your_organisation_id
```

## Usage

### Authentication

To obtain an OAuth2 token:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	clientID := "your_client_id"
	clientSecret := "your_client_secret"
	code := "authorization_code"
	redirectURI := "https://yourapp.com/callback"

	tokenResponse, err := GetOAuthToken(clientID, clientSecret, code, redirectURI)
	if err != nil {
		log.Fatalf("Error obtaining token: %v", err)
	}

	fmt.Printf("Access Token: %s\n", tokenResponse.AccessToken)
}
```

### Products

To list and get details of products:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := NewClient(config.BaseURL, config.AccessToken, config.ProductKey, config.OrganisationID)

	// List products
	products, err := client.GetProducts()
	if err != nil {
		log.Fatalf("Error getting products: %v", err)
	}
	fmt.Printf("Products: %+v\n", products)

	// Get a specific product
	product, err := client.GetProduct("product_id")
	if err != nil {
		log.Fatalf("Error getting product: %v", err)
	}
	fmt.Printf("Product: %+v\n", product)
}
```

### Plans

To list and get details of plans:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := NewClient(config.BaseURL, config.AccessToken, config.ProductKey, config.OrganisationID)

	// List plans
	plans, err := client.GetPlans()
	if err != nil {
		log.Fatalf("Error getting plans: %v", err)
	}
	fmt.Printf("Plans: %+v\n", plans)

	// Get a specific plan
	plan, err := client.GetPlan("plan_id")
	if err != nil {
		log.Fatalf("Error getting plan: %v", err)
	}
	fmt.Printf("Plan: %+v\n", plan)
}
```

### Customers

To create, get, update, and delete customers:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := NewClient(config.BaseURL, config.AccessToken, config.ProductKey, config.OrganisationID)

	// Create a customer
	customer := &Customer{Name: "John Doe", Email: "john.doe@example.com"}
	createdCustomer, err := client.CreateCustomer(customer)
	if err != nil {
		log.Fatalf("Error creating customer: %v", err)
	}
	fmt.Printf("Created Customer: %+v\n", createdCustomer)

	// Get a customer
	customer, err := client.GetCustomer(createdCustomer.ID)
	if err != nil {
		log.Fatalf("Error getting customer: %v", err)
	}
	fmt.Printf("Customer: %+v\n", customer)

	// Update a customer
	customer.Name = "John Doe Jr."
	updatedCustomer, err := client.UpdateCustomer(customer.ID, customer)
	if err != nil {
		log.Fatalf("Error updating customer: %v", err)
	}
	fmt.Printf("Updated Customer: %+v\n", updatedCustomer)

	// Delete a customer
	err = client.DeleteCustomer(customer.ID)
	if err != nil {
		log.Fatalf("Error deleting customer: %v", err)
	}
	fmt.Println("Customer deleted successfully")
}
```

### Subscriptions

To create, get, update, and delete subscriptions:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := NewClient(config.BaseURL, config.AccessToken, config.ProductKey, config.OrganisationID)

	// Create a subscription
	subscription := &Subscription{CustomerID: "customer_id", PlanID: "plan_id"}
	createdSubscription, err := client.CreateSubscription(subscription)
	if err != nil {
		log.Fatalf("Error creating subscription: %v", err)
	}
	fmt.Printf("Created Subscription: %+v\n", createdSubscription)

	// Get a subscription
	subscription, err := client.GetSubscription(createdSubscription.ID)
	if err != nil {
		log.Fatalf("Error getting subscription: %v", err)
	}
	fmt.Printf("Subscription: %+v\n", subscription)

	// Update a subscription
	subscription.PlanID = "new_plan_id"
	updatedSubscription, err := client.UpdateSubscription(subscription.ID, subscription)
	if err != nil {
		log.Fatalf("Error updating subscription: %v", err)
	}
	fmt.Printf("Updated Subscription: %+v\n", updatedSubscription)

	// Cancel a subscription
	err = client.DeleteSubscription(subscription.ID)
	if err != nil {
		log.Fatalf("Error canceling subscription: %v", err)
	}
	fmt.Println("Subscription canceled successfully")
}
```

### Entitlements

To get customer entitlements and log entitlement usage:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := NewClient(config.BaseURL, config.AccessToken, config.ProductKey, config.OrganisationID)

	// Get customer entitlements
	entitlements, err := client.GetCustomerEntitlements("customer_id")
	if err != nil {
		log.Fatalf("Error getting entitlements: %v", err)
	}
	fmt.Printf("Entitlements: %+v\n", entitlements)

	// Log entitlement usage
	usage := &EntitlementUsage{CustomerID: "customer_id", UsageCount: 5}
	message, err := client.LogEntitlementUsage("entitlement_id", usage)
	if err != nil {
		log.Fatalf("Error logging entitlement usage: %v", err)
	}
	fmt.Println(message)
}
```

## Running Tests

To run the tests, use:

```sh
make test
```

This will execute all the tests in the SDK.

## Contributing

If you find any issues or have suggestions for improvements, please open an issue or submit a pull request. Contributions are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.