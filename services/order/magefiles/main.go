package main

import (
	"os"
)

// Clean cleans generated files
func Clean() error {
	err := os.RemoveAll("./deployment-aws/cdk.out")
	if err != nil {
		return err
	}
	err = os.RemoveAll("./.aws-sam")
	if err != nil {
		return err
	}

	return nil
}
