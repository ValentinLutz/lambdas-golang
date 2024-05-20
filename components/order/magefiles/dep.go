//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Dep mg.Namespace

// Generate generates the models from the open api specification
func (Dep) Generate() error {
	return sh.RunV("go", "generate", "./...")
}
