package main

import (
	"fmt"
	"os"
)

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
