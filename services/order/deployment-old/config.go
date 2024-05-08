package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var stageConfigs = map[string]*StageConfig{
	"eu-central-1-test": {
		account:     "489721517942",
		region:      "eu-central-1",
		environment: "test",
		databaseProps: DatabaseConfig{
			host:   "",
			port:   "",
			name:   "",
			secret: "",
		},
		lambdaConfig: LambdaConfig{
			architecture: awslambda.Architecture_ARM_64(),
			orderRegion:  "EU",
		},
	},
	"eu-central-1-e2e": {
		account:     "489721517942",
		region:      "eu-central-1",
		environment: "e2e",
		databaseProps: DatabaseConfig{
			host:   "",
			port:   "",
			name:   "",
			secret: "",
		},
		lambdaConfig: LambdaConfig{
			architecture: awslambda.Architecture_ARM_64(),
			orderRegion:  "EU",
		},
	},
	"eu-central-1-prod": {
		account:     "489721517942",
		region:      "eu-central-1",
		environment: "prod",
		databaseProps: DatabaseConfig{
			host:   "",
			port:   "",
			name:   "",
			secret: "",
		},
		lambdaConfig: LambdaConfig{
			architecture: awslambda.Architecture_ARM_64(),
			orderRegion:  "EU",
		},
	},
}

type DatabaseConfig struct {
	host   string
	port   string
	name   string
	secret string
}

type LambdaConfig struct {
	architecture awslambda.Architecture
	orderRegion  string
}

type StageConfig struct {
	account       string
	region        string
	environment   string
	endpointUrl   *string
	databaseProps DatabaseConfig
	lambdaConfig  LambdaConfig
}

var (
	ErrStageConfigNotFound      = fmt.Errorf("stage config not found")
	ErrRegionEnvNotSet          = fmt.Errorf("env REGION not set")
	ErrEnvironmentEnvNotSet     = fmt.Errorf("env ENVIRONMENT not set")
	ErrArchitectureNotSupported = fmt.Errorf("architecture not supported")
)

func NewStageConfig() (*StageConfig, error) {
	region, ok := os.LookupEnv("REGION")
	if !ok {
		return nil, ErrRegionEnvNotSet
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return nil, ErrEnvironmentEnvNotSet
	}

	stageKey := region + "-" + env
	stage, ok := stageConfigs[stageKey]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrStageConfigNotFound, stageKey)
	}

	return stage, nil
}

func NewIdWithStage(stage *StageConfig, id string) *string {
	envTitleFormat := cases.Title(language.English).String(stage.environment)
	return jsii.String(id + envTitleFormat)
}

func GetArchitecture() awslambda.Architecture {
	switch runtime.GOARCH {
	case "amd64":
		return awslambda.Architecture_X86_64()
	case "arm64":
		return awslambda.Architecture_ARM_64()
	default:
		panic(fmt.Errorf("%w: %s", ErrArchitectureNotSupported, runtime.GOARCH))
	}
}
