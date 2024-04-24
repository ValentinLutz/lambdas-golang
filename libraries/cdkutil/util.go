package cdkutil

import "github.com/magefile/mage/sh"

func StackName(resource string, region string, env string) string {
	return resource + "-" + region + "-" + env
}

func GitCommit() string {
	commitHash, err := sh.Output("git", "rev-parse", "HEAD")
	if err != nil {
		panic(err)
	}

	return commitHash
}
