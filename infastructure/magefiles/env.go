package main

import (
	"fmt"
	"os"
)

type StageProps struct {
	Region      string
	Environment string
	Resource    string
	Version     string
}

func getStageEnvVars() StageProps {
	return StageProps{
		Region:      getEnvValueOrPanic("REGION"),
		Environment: getEnvValueOrPanic("ENVIRONMENT"),
		Resource:    getEnvValueOrPanic("RESOURCE"),
		Version:     getEnvValueOrSetDefault("VERSION", "latest"),
	}
}

func getEnvValueOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("env '%s' not set", key))
	}
	return value
}

func getEnvValueOrSetDefault(key string, defaultValue string) string {
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
