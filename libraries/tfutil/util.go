package tfutil

import (
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
)

func NewStateFilePath(region string, env string, component string) string {
	return region + "-" + env + "/" + component + "/terraform.tfstate"
}

func NewComponentName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Base(dir), nil
}

func GitCommit() string {
	commitHash, err := sh.Output("git", "rev-parse", "HEAD")
	if err != nil {
		panic(err)
	}

	return commitHash
}
