package main

import (
	"fmt"
	"os"

	"github.com/aws/jsii-runtime-go"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var stageConfigs = map[string]*StageConfig{
	"eu-central-1-test": {
		account:     "489721517942",
		region:      "eu-central-1",
		environment: "test",
	},
}

type StageConfig struct {
	account     string
	region      string
	environment string
}

var (
	ErrStageConfigNotFound  = fmt.Errorf("stage config not found")
	ErrRegionEnvNotSet      = fmt.Errorf("env REGION not set")
	ErrEnvironmentEnvNotSet = fmt.Errorf("env ENVIRONMENT not set")
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
