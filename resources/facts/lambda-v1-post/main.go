package main

import (
	"root/resources/facts/lambda-v1-post/incoming"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := incoming.NewHandler()
	lambda.Start(handler.Invoke)
}
