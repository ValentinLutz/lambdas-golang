package testintegration

import (
	"os"
	v1GetOrders "root/services/order/lambda-v1-get-orders/incoming"
	v1PostOrders "root/services/order/lambda-v1-post-orders/incoming"
)

var (
	V1GetOrdersHandler  *v1GetOrders.Handler
	V1PostOrdersHandler *v1PostOrders.Handler
)

func init() {
	NewTestConfig()

	V1GetOrdersHandler = NewV1GetOrdersHandler()
	V1PostOrdersHandler = NewV1PostOrdersHandler()
}

func NewV1GetOrdersHandler() *v1GetOrders.Handler {
	handler, err := v1GetOrders.NewHandler()
	if err != nil {
		panic(err)
	}
	return handler
}

func NewV1PostOrdersHandler() *v1PostOrders.Handler {
	handler, err := v1PostOrders.NewHandler()
	if err != nil {
		panic(err)
	}
	return handler
}

func NewTestConfig() {
	envVars := map[string]string{
		"AWS_REGION":            "eu-central-1",
		"AWS_ACCOUNT":           "000000000000",
		"AWS_ACCESS_KEY_ID":     "test",
		"AWS_SECRET_ACCESS_KEY": "test",
		"AWS_ENDPOINT_URL":      "http://127.0.0.1:4566",
		"DB_HOST":               "127.0.0.1",
		"DB_PORT":               "5432",
		"DB_NAME":               "test",
		"DB_SECRET_ID":          "database-secret",
		"ORDER_REGION":          "EU",
	}

	for key, value := range envVars {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}
}
