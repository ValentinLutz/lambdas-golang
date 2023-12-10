//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Dep mg.Namespace

// Generate generates the models from the OpenAPI specification
func (Dep) Generate() error {
	return sh.RunV("oapi-codegen", "--config", "./api-definition/oapi-codgen.yaml", "./api-definition/order-api-v1.yaml")
}
