//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

func (Test) Integration() {
	getOrSetDefaultStage()

	os.Chdir("./test-integration")
	defer os.Chdir("..")

	sh.RunV("go", "test", "-cover", "-coverpkg=../...", "-coverprofile=coverage.out", "-count=1", "-p=1", "./...")
}

func (Test) Coverage() {
	getOrSetDefaultStage()

	os.Chdir("./test-integration")
	defer os.Chdir("..")

	sh.RunV("go", "tool", "cover", "-html", "coverage.out", "-o", "coverage.html")
}
