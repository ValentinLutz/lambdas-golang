package tfutil

import (
	"fmt"
	"os"
)

var s3BackendConfigs = map[string]*S3BackendConfig{
	"eu-central-1-test": {
		Profile:       "admin",
		Bucket:        "terraform-state-9f3933f9",
		DynamoDbTable: "terraform-lock-9f3933f9",
		Encrypt:       "true",
	},
}

type S3BackendConfig struct {
	Profile       string
	Bucket        string
	DynamoDbTable string
	Encrypt       string
}

var (
	ErrStageConfigNotFound  = fmt.Errorf("stage config not found")
	ErrRegionEnvNotSet      = fmt.Errorf("env REGION not set")
	ErrEnvironmentEnvNotSet = fmt.Errorf("env ENVIRONMENT not set")
)

func NewS3BackendConfig() (*S3BackendConfig, error) {
	region, ok := os.LookupEnv("REGION")
	if !ok {
		return nil, ErrRegionEnvNotSet
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return nil, ErrEnvironmentEnvNotSet
	}

	stageKey := region + "-" + env
	stage, ok := s3BackendConfigs[stageKey]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrStageConfigNotFound, stageKey)
	}

	return stage, nil
}
