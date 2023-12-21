//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Dep mg.Namespace

// Install installs the dependencies to generate the api models
func (Dep) Install() error {
	return sh.RunV("go", "install", "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0")
}

// Generate generates the models from the open api specification
func (Dep) Generate() error {
	err := sh.RunV("mkdir", "-p", "./lambda-shared/incoming")
	if err != nil {
		return err
	}
	return sh.RunV("oapi-codegen", "--config", "./api-definition/oapi-codgen.yaml", "./api-definition/order-api-v1.yaml")
}
