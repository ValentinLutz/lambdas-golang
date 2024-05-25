package tofu

import (
	"os"
	"path/filepath"
)

func NewStateFilePath(component string) string {
	return component + "/terraform.tfstate"
}

func NewConfigPath(region string, environment string) string {
	return "../../configs/" + region + "-" + environment + ".yaml"
}

func NewRootPath(region string, environment string) string {
	return "deployment-aws/" + region + "-" + environment
}

func NewComponentName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Base(dir), nil
}
