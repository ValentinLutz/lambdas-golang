package tofu

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"root/libraries/tfutil"
)

type Tofu mg.Namespace

// Upgrade upgrade the tofu providers
func (Tofu) Upgrade() error {
	stageEnvVars := getStageEnvVars()

	config, err := tfutil.NewConfig()
	if err != nil {
		return err
	}

	component, err := tfutil.NewComponentName()
	if err != nil {
		return err
	}

	os.Chdir("./deployment-aws/" + stageEnvVars.Region + "-" + stageEnvVars.Environment)
	defer os.Chdir("../..")

	err = sh.RunV("tofu",
		"init",
		"-upgrade",
		"-backend-config=region="+stageEnvVars.Region,
		"-backend-config=bucket="+config.S3BackendConfig.Bucket,
		"-backend-config=dynamodb_table="+config.S3BackendConfig.DynamoDbTable,
		"-backend-config=encrypt="+config.S3BackendConfig.Encrypt,
		"-backend-config=key="+tfutil.NewStateFilePath(stageEnvVars.Region, stageEnvVars.Environment, component),
	)
	if err != nil {
		return err
	}

	return sh.RunV("tofu",
		"providers",
		"lock",
		"-platform=darwin_amd64",
		"-platform=linux_amd64",
		"-platform=windows_amd64",
	)
}

// Init initializes the tofu project
func (Tofu) Init() error {
	stageEnvVars := getStageEnvVars()

	config, err := tfutil.NewConfig()
	if err != nil {
		return err
	}

	component, err := tfutil.NewComponentName()
	if err != nil {
		return err
	}

	os.Chdir("./deployment-aws/" + stageEnvVars.Region + "-" + stageEnvVars.Environment)
	defer os.Chdir("../..")

	return sh.RunV("tofu",
		"init",
		"-lockfile=readonly",
		"-backend-config=region="+stageEnvVars.Region,
		"-backend-config=bucket="+config.S3BackendConfig.Bucket,
		"-backend-config=dynamodb_table="+config.S3BackendConfig.DynamoDbTable,
		"-backend-config=encrypt="+config.S3BackendConfig.Encrypt,
		"-backend-config=key="+tfutil.NewStateFilePath(stageEnvVars.Region, stageEnvVars.Environment, component),
	)
}

// Plan creates an execution plan
func (Tofu) Plan() error {
	mg.Deps(Tofu.Init)

	stageEnvVars := getStageEnvVars()

	config, err := tfutil.NewConfig()
	if err != nil {
		return err
	}

	component, err := tfutil.NewComponentName()
	if err != nil {
		return err
	}

	os.Chdir("./deployment-aws/" + stageEnvVars.Region + "-" + stageEnvVars.Environment)
	defer os.Chdir("../..")

	return sh.RunV("tofu",
		"plan",
		"-out=terraform.tfplan",
		"-var=region="+stageEnvVars.Region,
		"-var=environment="+stageEnvVars.Environment,
		"-var=project="+config.Project,
		"-var=component="+component,
	)

}

// Plandestroy creates an execution plan to destroy
func (Tofu) Plandestroy() error {
	mg.Deps(Tofu.Init)

	stageEnvVars := getStageEnvVars()

	config, err := tfutil.NewConfig()
	if err != nil {
		return err
	}

	component, err := tfutil.NewComponentName()
	if err != nil {
		return err
	}

	os.Chdir("./deployment-aws/" + stageEnvVars.Region + "-" + stageEnvVars.Environment)
	defer os.Chdir("../..")

	return sh.RunV("tofu",
		"plan",
		"-out=terraform.tfplan",
		"-var=region="+stageEnvVars.Region,
		"-var=environment="+stageEnvVars.Environment,
		"-var=project="+config.Project,
		"-var=component="+component,
		"-destroy",
	)

}

// Show shows the planned changes
func (Tofu) Show() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./deployment-aws/" + stageEnvVars.Region + "-" + stageEnvVars.Environment)
	defer os.Chdir("../..")

	return sh.RunV("tofu",
		"show",
		"terraform.tfplan",
	)

}

// Apply applies the planned changes
func (Tofu) Apply() error {
	stageEnvVars := getStageEnvVars()

	os.Chdir("./deployment-aws/" + stageEnvVars.Region + "-" + stageEnvVars.Environment)
	defer os.Chdir("../..")

	return sh.RunV("tofu",
		"apply",
		"terraform.tfplan",
	)
}
