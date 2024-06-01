package main

import (
	"fmt"
	"os"
	"runtime"
)

type DatabaseProps struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func getOrSetDefaultDatabaseEnvVars() *DatabaseProps {
	return &DatabaseProps{
		Host:     getValueOrSetDefault("DB_HOST", "localhost"),
		Port:     getValueOrSetDefault("DB_PORT", "5432"),
		Name:     getValueOrSetDefault("DB_NAME", "postgres"),
		User:     getValueOrSetDefault("DB_USER", "test"),
		Password: getValueOrSetDefault("DB_PASS", "test"),
	}
}

type BuildProps struct {
	OperatingSystem string
	Architecture    string
}

func getOrSetDefaultBuildEnvVars() *BuildProps {
	stageEnvVars := getOrSetDefaultStageEnvVars()

	getValueOrSetDefault("CGO_ENABLED", "0")

	if stageEnvVars.Environment == "local" {
		return &BuildProps{
			OperatingSystem: getValueOrSetDefault("GOOS", runtime.GOOS),
			Architecture:    getValueOrSetDefault("GOARCH", runtime.GOARCH),
		}
	}
	return &BuildProps{
		OperatingSystem: getValueOrSetDefault("GOOS", "linux"),
		Architecture:    getValueOrSetDefault("GOARCH", "arm64"),
	}
}

type StageProps struct {
	Environment string
	Region      string
}

func getOrSetDefaultStageEnvVars() *StageProps {
	return &StageProps{
		Environment: getValueOrSetDefault("ENVIRONMENT", "local"),
		Region:      getValueOrSetDefault("REGION", "eu-central-1"),
	}
}

func getValueOrSetDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("env '%s' not set, defaulting to '%s'\n", key, defaultValue)
		err := os.Setenv(key, defaultValue)
		if err != nil {
			panic(err)
		}
		return defaultValue
	}
	return value
}
