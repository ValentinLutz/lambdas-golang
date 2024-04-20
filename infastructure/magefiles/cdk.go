//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"root/infastructure/util"
)

type Cdk mg.Namespace

// Synth synthesizes the CDK stack
func (Cdk) Synth() error {
	return sh.RunV("cdktf",
		"synth",
	)
}

// Diff diffs the CDK stack
func (Cdk) Diff() error {
	stageEnvVars := getStageEnvVars()

	return sh.RunV("cdktf",
		"diff", util.StackName(stageEnvVars.Resource, stageEnvVars.Region, stageEnvVars.Environment),
	)
}

// Deploy deploys the CDK stack
func (Cdk) Deploy() error {
	stageEnvVars := getStageEnvVars()

	return sh.RunV("cdktf",
		"deploy", util.StackName(stageEnvVars.Resource, stageEnvVars.Region, stageEnvVars.Environment),
	)
}

// Destroy destroys the CDK stack
func (Cdk) Destroy() error {
	stageEnvVars := getStageEnvVars()

	return sh.RunV(
		"cdktf",
		"destroy", util.StackName(stageEnvVars.Resource, stageEnvVars.Region, stageEnvVars.Environment),
	)
}
