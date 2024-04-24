//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"root/libraries/cdkutil"
)

type Cdk mg.Namespace

// Synth synthesizes the CDK stack
func (Cdk) Synth() error {
	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV("cdktf",
		"synth",
	)
}

// Diff diffs the CDK stack
func (Cdk) Diff() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV("cdktf",
		"diff", cdkutil.StackName(stageEnvVars.Resource, stageEnvVars.Region, stageEnvVars.Environment),
	)
}

// Deploy deploys the CDK stack
func (Cdk) Deploy() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV("cdktf",
		"deploy", cdkutil.StackName(stageEnvVars.Resource, stageEnvVars.Region, stageEnvVars.Environment),
	)
}

// Destroy destroys the CDK stack
func (Cdk) Destroy() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./deployment-aws")
	defer os.Chdir("..")

	return sh.RunV(
		"cdktf",
		"destroy", cdkutil.StackName(stageEnvVars.Resource, stageEnvVars.Region, stageEnvVars.Environment),
	)
}
