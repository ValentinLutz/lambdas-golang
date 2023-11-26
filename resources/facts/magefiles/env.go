package main

import (
	"fmt"
	"os"
)

func getOrSetDefaultDatabaseEnvVars() {
	getValueOrSetDefault("DB_HOST", "localhost")
	getValueOrSetDefault("DB_PORT", "5432")
	getValueOrSetDefault("DB_NAME", "test")
	getValueOrSetDefault("DB_USER", "test")
	getValueOrSetDefault("DB_PASS", "test")
}

func getOrSetDefaultStage() {
	getValueOrSetDefault("REGION", "eu-central-1")
	getValueOrSetDefault("ENVIRONMENT", "local")
	getValueOrSetDefault("VERSION", "latest")
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