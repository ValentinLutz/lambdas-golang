package main

import (
	"root/services/order/lambda-v1-get-orders/incoming"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler, err := incoming.NewHandler()
	if err != nil {
		panic(err)
	}
	lambda.Start(handler.Invoke)
}
