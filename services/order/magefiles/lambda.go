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
	mg.Deps(Cdk.Synth)

	return sh.RunV(
		"sam", "local", "start-api",
		"--template", "./deployment-aws/cdk.out/order-resource-eu-central-1-local.template.json",
		"--docker-network", "lambda-network",
		"--warm-containers", "EAGER",
	)
}

// Build builds all lambda functions
func (Lambda) Build() error {
	getOrSetDefaultBuildEnvVars()

	lambdas := []string{
		"./lambda-v1-get-order",
		"./lambda-v1-get-orders",
		"./lambda-v1-post-orders",
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
		"go",
		"build",
		"-ldflags=-s",
		"-ldflags=-w",
		"-trimpath",
		"-buildvcs=false",
		"-tags", "lambda.norpc",
		"-o", "bootstrap",
	)
}
