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

// Upgrade upgrade the terraform providers
func (Tofu) Upgrade() error {
	stageEnvVars := getStageEnvVars()

	backendConfig, err := tfutil.NewS3BackendConfig()
	if err != nil {
		return err
	}

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)
	defer os.Chdir("../../..")

	return sh.RunV("tofu",
		"init",
		"-upgrade",
		"-backend-config=region="+stageEnvVars.Region,
		"-backend-config=bucket="+backendConfig.Bucket,
		"-backend-config=dynamodb_table="+backendConfig.DynamoDbTable,
		"-backend-config=encrypt="+backendConfig.Encrypt,
		"-backend-config=key="+tfutil.NewStateFilePath(stageEnvVars.Region, stageEnvVars.Environment, stageEnvVars.Resource),
	)
}

// Init initializes the terraform project
func (Tofu) Init() error {
	stageEnvVars := getStageEnvVars()

	backendConfig, err := tfutil.NewS3BackendConfig()
	if err != nil {
		return err
	}

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)
	defer os.Chdir("../../..")

	return sh.RunV("tofu",
		"init",
		"-lockfile=readonly",
		"-backend-config=region="+stageEnvVars.Region,
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
	defer os.Chdir("../../..")

	return sh.RunV("tofu",
		"plan",
		"-out=terraform.tfplan",
		"-var-file=../terraform.tfvars",
		"-var=profile="+backendConfig.Profile,
		"-var=resource="+stageEnvVars.Resource,
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
	defer os.Chdir("../../..")

	return sh.RunV("tofu",
		"plan",
		"-out=terraform.tfplan",
		"-var-file=../terraform.tfvars",
		"-var=profile="+backendConfig.Profile,
		"-var=resource="+stageEnvVars.Resource,
		"-destroy",
	)

}

// Show shows the execution plan
func (Tofu) Show() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)
	defer os.Chdir("../../..")

	return sh.RunV("tofu",
		"show",
		"terraform.tfplan",
	)

}

// Apply applies the changes
func (Tofu) Apply() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./environments/" + stageEnvVars.Region + "-" + stageEnvVars.Environment + "/" + stageEnvVars.Resource)
	defer os.Chdir("../../..")

	return sh.RunV("tofu",
		"apply",
		"terraform.tfplan",
	)
}
