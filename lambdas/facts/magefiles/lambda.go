//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Lambda mg.Namespace

// Build builds the lambda function
func (Lambda) Build() error {
	return sh.RunV("sam", "build")
}

// Invoke invokes the lambda function locally
func (Lambda) Invoke() error {
	return sh.RunV("sam", "local", "invoke", "--docker-network", "facts-lambda-network")
}

// Start starts the lambda function locally
func (Lambda) Start() error {
	return sh.RunV("sam", "local", "start-api", "--docker-network", "facts-lambda-network")
}
