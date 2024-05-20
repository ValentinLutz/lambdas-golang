package tofu

import (
	"fmt"
	"os"
)

type StageProps struct {
	Region      string
	Environment string
	Resource    string
}

func getStageEnvVars() StageProps {
	return StageProps{
		Region:      getEnvValueOrPanic("REGION"),
		Environment: getEnvValueOrPanic("ENVIRONMENT"),
		Resource:    getEnvValueOrPanic("RESOURCE"),
	}
}

func getEnvValueOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("env '%s' not set", key))
	}
	return value
}
