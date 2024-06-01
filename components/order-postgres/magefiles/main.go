package main

import (
	"os"
)

// Clean cleans generated files
func Clean() error {
	paths := []string{
		"./lambda-v1-get-order/incoming/model.gen.go",
		"./lambda-v1-get-orders/incoming/model.gen.go",
		"./lambda-v1-post-orders/incoming/model.gen.go",
		"./lambda-v1-get-order/bootstrap",
		"./lambda-v1-get-orders/bootstrap",
		"./lambda-v1-post-orders/bootstrap",
		"./test-integration/coverage.out",
		"./test-integration/coverage.html",
	}

	for _, path := range paths {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}

	return nil
}
