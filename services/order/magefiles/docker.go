//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Up starts the docker-compose stack
func (Docker) Up() error {
	getOrSetDefaultDatabaseEnvVars()

	os.Chdir("./deployment-local")
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

	os.Chdir("./deployment-local")
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

	os.Chdir("./deployment-local")
	defer os.Chdir("..")

	return sh.RunV(
		"docker",
		"compose",
		"--file", "./docker-compose.yaml",
		"logs",
	)
}
