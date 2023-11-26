package testutil

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BaseURL  string `yaml:"base_url"`
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

func LoadConfig(basePath string) *Config {
	region, ok := os.LookupEnv("REGION")
	if ok != true {
		panic("env REGION is not set")
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if ok != true {
		panic("env ENVIRONMENT is not set")
	}

	parsedFile, err := ParseFile[Config](basePath + "config." + region + "-" + env + ".yaml")
	if err != nil {
		panic(err)
	}
	return parsedFile
}

func ParseFile[T any](path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	var decodedConfig *T
	err = decoder.Decode(&decodedConfig)
	if err != nil {
		return nil, err
	}

	return decodedConfig, nil
}
