//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Cdk mg.Namespace

// Synth synthesizes the CDK stack
func (Cdk) Synth() error {
	mg.Deps(Clean)

	getOrSetDefaultStage()

	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV(
		"cdk",
		"synth",
	)
}

// Diff diffs the CDK stack
func (Cdk) Diff() error {
	mg.Deps(Clean)

	getOrSetDefaultStage()

	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV(
		"cdk",
		"diff",
	)
}

// Deploy deploys the CDK stack
func (Cdk) Deploy() error {
	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV(
		"cdk",
		"deploy",
	)
}

// Destroy destroys the CDK stack
func (Cdk) Destroy() error {
	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV(
		"cdk",
		"destroy",
	)
}
