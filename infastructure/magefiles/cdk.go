//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
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
		"diff", createStackName(stageEnvVars),
		"--var=commit_hash="+stageEnvVars.Version,
	)
}

// Deploy deploys the CDK stack
func (Cdk) Deploy() error {
	stageEnvVars := getStageEnvVars()

	return sh.RunV("cdktf",
		"deploy", createStackName(stageEnvVars),
		"--var=commit_hash="+stageEnvVars.Version,
	)
}

// Destroy destroys the CDK stack
func (Cdk) Destroy() error {
	stageEnvVars := getStageEnvVars()

	return sh.RunV(
		"cdktf",
		"destroy", createStackName(stageEnvVars),
		"--var=commit_hash="+stageEnvVars.Version,
	)
}

func createStackName(stageProps StageProps) string {
	return stageProps.Resource + "-" + stageProps.Region + "-" + stageProps.Environment
}
