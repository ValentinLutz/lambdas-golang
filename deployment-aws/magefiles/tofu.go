//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"root/libraries/tfutil"
)

type Tofu mg.Namespace

// Init initializes the terraform project
func (Tofu) Init() error {
	stageEnvVars := getStageEnvVars()

	backendConfig, err := tfutil.NewS3BackendConfig()
	if err != nil {
		return err
	}

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)

	return sh.RunV("tofu",
		"init",
		"-backend-config=region="+stageEnvVars.Region,
		"-backend-config=profile="+backendConfig.Profile,
		"-backend-config=bucket="+backendConfig.Bucket,
		"-backend-config=dynamodb_table="+backendConfig.DynamoDbTable,
		"-backend-config=encrypt="+backendConfig.Encrypt,
		"-backend-config=key="+tfutil.NewStateFilePath(stageEnvVars.Region, stageEnvVars.Environment, stageEnvVars.Resource),
	)
}

// Plan creates an execution plan
func (Tofu) Plan() error {
	mg.Deps(Tofu.Init)

	stageEnvVars := getStageEnvVars()

	backendConfig, err := tfutil.NewS3BackendConfig()
	if err != nil {
		return err
	}

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)

	return sh.RunV("tofu",
		"plan",
		"-out=terraform.tfplan",
		"-var-file=../terraform.tfvars",
		"-var=profile="+backendConfig.Profile,
		"-var=resource="+stageEnvVars.Resource,
	)

}

// Apply applies the changes
func (Tofu) Apply() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)

	return sh.RunV("tofu",
		"apply",
		"terraform.tfplan",
	)
}

// Plandestroy creates an execution plan for destroying the resources
func (Tofu) Plandestroy() error {
	mg.Deps(Tofu.Init)

	stageEnvVars := getStageEnvVars()

	backendConfig, err := tfutil.NewS3BackendConfig()
	if err != nil {
		return err
	}

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)

	return sh.RunV("tofu",
		"plan",
		"-out=terraform.tfplan",
		"-var-file=../terraform.tfvars",
		"-var=profile="+backendConfig.Profile,
		"-var=resource="+stageEnvVars.Resource,
		"-destroy",
	)

}

// Destroy destroys the resources
func (Tofu) Destroy() error {
	mg.Deps(Tofu.Init)

	stageEnvVars := getStageEnvVars()

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)

	return sh.RunV("tofu",
		"destroy",
		"terraform.tfplan",
	)
}

// Show shows the execution plan
func (Tofu) Show() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)

	return sh.RunV("tofu",
		"show",
		"terraform.tfplan",
	)

}
