//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

func (Test) Functional() {
	getOrSetDefaultStage()

	os.Chdir("./test-functional")
	defer os.Chdir("..")

	sh.RunV("go", "test", "-count=1", "./...")
}
