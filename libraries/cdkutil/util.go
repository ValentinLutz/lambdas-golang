package cdkutil

import (
	"github.com/magefile/mage/sh"
)

func StackName(region string, env string) string {
	return region + "-" + env
}

func StackStateFile(region string, env string, resource string) string {
	return region + "-" + env + "/" + resource + "/terraform.tfstate"
}

func GitCommit() string {
	commitHash, err := sh.Output("git", "rev-parse", "HEAD")
	if err != nil {
		panic(err)
	}

	return commitHash
}
