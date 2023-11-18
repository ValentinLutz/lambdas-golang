//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Lambda mg.Namespace

// Build builds the lambda function
func (Lambda) Build() error {
	mg.Deps(Cdk.Synth)

	return sh.RunV(
		"sam", "build",
		"--template", "./deployment-aws/cdk.out/stack.template.json",
	)
}

// Start starts the lambda function locally
func (Lambda) Start() error {
	mg.Deps(Lambda.Build)

	return sh.RunV(
		"sam", "local", "start-api",
		"--config-file", "./deployment-docker/samconfig.yaml",
		"--docker-network", "facts-lambda-network",
	)
}
