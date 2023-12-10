package test_integration

import (
	"os"
	getv1 "root/services/order/lambda-v1-get/incoming"
	postv1 "root/services/order/lambda-v1-post/incoming"
)

var (
	orderV1PostHandler *postv1.Handler
	orderV1GetHandler  *getv1.Handler
)

func init() {
	NewTestConfig()
	orderV1GetHandler = getv1.NewHandler()
	orderV1PostHandler = postv1.NewHandler()
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
	}

	for key, value := range envVars {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}
}
