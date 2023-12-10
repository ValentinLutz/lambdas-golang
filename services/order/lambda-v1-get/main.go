package main

import (
	"root/services/order/lambda-v1-get/incoming"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := incoming.NewHandler()
	lambda.Start(handler.Invoke)
}
