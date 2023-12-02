//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Lambda mg.Namespace

// Start starts the lambda function locally
func (Lambda) Start() error {
	mg.Deps(Lambda.Build)

	return sh.RunV(
		"sam", "local", "start-api",
		"--template", "./deployment-aws/cdk.out/facts-resource-eu-central-1-local.template.json",
		"--docker-network", "lambda-network",
		"--warm-containers", "EAGER",
	)
}

// Build builds all lambda functions
func (Lambda) Build() error {
	lambdas := []string{
		"./lambda-v1-get",
		"./lambda-v1-post",
	}

	for _, lambda := range lambdas {
		err := build(lambda)
		if err != nil {
			return err
		}
	}

	return nil
}

func build(path string) error {
	os.Chdir(path)
	defer os.Chdir("..")

	return sh.RunV(
		"go", "build", "-tags", "lambda.norpc", "-o", "bootstrap",
	)
}
