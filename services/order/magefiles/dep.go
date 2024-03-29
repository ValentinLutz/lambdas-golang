//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Dep mg.Namespace

// Install installs the dependencies for mage targets
func (Dep) Install() error {
	err := sh.RunV("go", "install", "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0")
	if err != nil {
		return err
	}
	return sh.RunV("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2")
}

// Generate generates the models from the open api specification
func (Dep) Generate() error {
	return sh.RunV("go", "generate", "./...")
}
