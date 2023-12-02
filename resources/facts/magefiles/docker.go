//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Build builds the lambda functions with docker
func (Docker) Build() error {
	os.Chdir("../../")
	defer os.Chdir("..")

	err := sh.RunV(
		"docker",
		"build",
		"--file", "resources/facts/lambda-v1-get/Dockerfile",
		".",
	)
	if err != nil {
		return err
	}

	return sh.RunV(
		"docker",
		"build",
		"--file", "resources/facts/lambda-v1-post/Dockerfile",
		".",
	)
}

// Up starts the docker-compose stack
func (Docker) Up() error {
	getOrSetDefaultDatabaseEnvVars()

	os.Chdir("./deployment-docker")
	defer os.Chdir("..")

	return sh.RunV(
		"docker",
		"compose",
		"--file", "./docker-compose.yaml",
		"up",
		"--force-recreate",
		//"--detach",
		"--wait",
	)
}

// Down stops the docker-compose stack
func (Docker) Down() error {
	getOrSetDefaultDatabaseEnvVars()

	os.Chdir("./deployment-docker")
	defer os.Chdir("..")

	return sh.RunV(
		"docker",
		"compose",
		"--file", "./docker-compose.yaml",
		"down",
	)
}

// Logs shows the logs of the docker-compose stack
func (Docker) Logs() error {
	getOrSetDefaultDatabaseEnvVars()

	os.Chdir("./deployment-docker")
	defer os.Chdir("..")

	return sh.RunV(
		"docker",
		"compose",
		"--file", "./docker-compose.yaml",
		"logs",
		"-f",
	)
}
