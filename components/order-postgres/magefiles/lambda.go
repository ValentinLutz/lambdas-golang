//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Lambda mg.Namespace

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
