package bootstrap

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
)

var stageConfigs = map[string]*StageConfig{
	"eu-central-1-test": {
		Region:      "eu-central-1",
		Environment: "test",
		Profile:     "admin",
		Bucket:      jsii.String("terraform-state-5449924400404832213"),
	},
}

type StageConfig struct {
	Region      string
	Environment string
	Profile     string
	Bucket      *string
}

var (
	ErrStageConfigNotFound = fmt.Errorf("stage config not found")
)

func NewStageConfig(region string, env string) (*StageConfig, error) {
	stageKey := region + "-" + env
	stage, ok := stageConfigs[stageKey]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrStageConfigNotFound, stageKey)
	}

	return stage, nil
}
