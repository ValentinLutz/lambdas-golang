package main

import (
	"fmt"
	"os"
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
		Name:     getValueOrSetDefault("DB_NAME", "test"),
		User:     getValueOrSetDefault("DB_USER", "test"),
		Password: getValueOrSetDefault("DB_PASS", "test"),
	}
}

type StageProps struct {
	Environment string
	Region      string
	Version     string
}

func getOrSetDefaultStageEnvVars() *StageProps {
	return &StageProps{
		Environment: getValueOrSetDefault("ENVIRONMENT", "local"),
		Region:      getValueOrSetDefault("REGION", "eu-central-1"),
		Version:     getValueOrSetDefault("VERSION", "latest"),
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
