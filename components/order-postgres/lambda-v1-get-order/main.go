package main

import (
	"root/components/order-postgres/lambda-v1-get-order/incoming"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler, err := incoming.NewHandler()
	if err != nil {
		panic(err)
	}
	lambda.Start(handler.Invoke)
}
