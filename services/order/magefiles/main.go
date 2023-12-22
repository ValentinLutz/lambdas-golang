package main

import (
	"os"
)

// Clean cleans generated files
func Clean() error {
	paths := []string{
		"./deployment-aws/cdk.out",
		"./.aws-sam",
		"./lambda-shared/incoming/model.gen.go",
		"./lambda-v1-post/bootstrap",
		"./lambda-v1-get/bootstrap",
	}

	for _, path := range paths {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}

	return nil
}
