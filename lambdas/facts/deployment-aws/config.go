package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

var stageConfigs = map[string]*StageConfig{
	"eu-central-1-dev": {
		account:     "000000000000",
		region:      "eu-central-1",
		environment: "dev",
		endpointUrl: jsii.String("http://aws-localstack:4566"),
		databaseProps: DatabaseConfig{
			host:   "database-postgres",
			port:   "5432",
			name:   "test",
			secret: "database-secret",
		},
		lambdaConfig: LambdaConfig{
			// use your local platform for faster builds
			architecture: awslambda.Architecture_X86_64(),
		},
	},
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
}

type StageConfig struct {
	account       string
	region        string
	environment   string
	endpointUrl   *string
	databaseProps DatabaseConfig
	lambdaConfig  LambdaConfig
}

func NewStageConfig() *StageConfig {
	region, ok := os.LookupEnv("REGION")
	if !ok {
		panic("env REGION not set")
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		panic("env ENVIRONMENT not set")
	}

	stageKey := region + "-" + env
	stage, ok := stageConfigs[stageKey]
	if !ok {
		panic(fmt.Errorf("stage config %s not found", stageKey))
	}
	return stage
}
