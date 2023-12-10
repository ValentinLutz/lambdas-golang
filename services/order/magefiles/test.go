//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

// Integration runs the integration tests
func (Test) Integration() error {
	getOrSetDefaultStageEnvVars()

	os.Chdir("./test-integration")
	defer os.Chdir("..")

	return sh.RunV("go", "test", "-cover", "-coverpkg=../...", "-coverprofile=coverage.out", "-count=1", "-p=1", "./...")
}

// Coverage generates the coverage report
func (Test) Coverage() error {
	getOrSetDefaultStageEnvVars()

	os.Chdir("./test-integration")
	defer os.Chdir("..")

	return sh.RunV("go", "tool", "cover", "-html", "coverage.out", "-o", "coverage.html")
}
