package tfutil

import (
	"fmt"
	"os"
)

var configs = map[string]*Config{
	"eu-central-1-test": {
		Project: "monke",
		S3BackendConfig: &S3BackendConfig{
			Bucket:        "monke-eu-central-1-test-tofu-state",
			DynamoDbTable: "monke-eu-central-1-test-tofu-state-lock",
			Encrypt:       "true",
		},
	},
}

type Config struct {
	Project         string
	S3BackendConfig *S3BackendConfig
}

type S3BackendConfig struct {
	Bucket        string
	DynamoDbTable string
	Encrypt       string
}

var (
	ErrStageConfigNotFound  = fmt.Errorf("stage config not found")
	ErrRegionEnvNotSet      = fmt.Errorf("env REGION not set")
	ErrEnvironmentEnvNotSet = fmt.Errorf("env ENVIRONMENT not set")
)

func NewConfig() (*Config, error) {
	region, ok := os.LookupEnv("REGION")
	if !ok {
		return nil, ErrRegionEnvNotSet
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return nil, ErrEnvironmentEnvNotSet
	}

	stageKey := region + "-" + env
	stageConfig, ok := configs[stageKey]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrStageConfigNotFound, stageKey)
	}

	return stageConfig, nil
}
