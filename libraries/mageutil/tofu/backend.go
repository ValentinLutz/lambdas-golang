package tofu

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Project         string           `yaml:"project"`
	S3BackendConfig *S3BackendConfig `yaml:"s3_backend"`
}

type S3BackendConfig struct {
	Bucket        string `yaml:"bucket"`
	DynamoDbTable string `yaml:"dynamodb_table"`
	Encrypt       string `yaml:"encrypt"`
}

var (
	ErrOpenConfigFile = errors.New("failed to open config file")
	ErrDecodeConfig   = errors.New("failed to decode config")
)

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Join(ErrOpenConfigFile, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var config *Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, errors.Join(ErrDecodeConfig, err)
	}

	return config, nil
}
