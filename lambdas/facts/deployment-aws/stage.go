package main

import (
	"fmt"
	"os"
)

var stages = map[string]*StageProps{
	"eu-central-1-dev": {
		account:     "000000000000",
		region:      "eu-central-1",
		environment: "dev",
	},
	"eu-central-1-prod": {
		account:     "489721517942",
		region:      "eu-central-1",
		environment: "test",
	},
}

type StageProps struct {
	account     string
	region      string
	environment string
}

func NewStageConfig() *StageProps {
	region, ok := os.LookupEnv("REGION")
	if !ok {
		panic("env REGION not set")
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		panic("env ENVIRONMENT not set")
	}

	stageKey := region + "-" + env
	stage, ok := stages[stageKey]
	if !ok {
		panic(fmt.Errorf("stage config %s not found", stageKey))
	}
	return stage
}
