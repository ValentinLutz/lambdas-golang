package apputil

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

var (
	ErrSecretUsernameNotSet = fmt.Errorf("secret username not set")
	ErrSecretPasswordNotSet = fmt.Errorf("secret password not set")
	ErrDbHostEnvNotSet      = fmt.Errorf("env DB_HOST not set")
	ErrDbPortEnvNotSet      = fmt.Errorf("env DB_PORT not set")
	ErrDbNameEnvNotSet      = fmt.Errorf("env DB_NAME not set")
)

func NewDatabaseConfig(secret Secret) (*DatabaseConfig, error) {
	if secret.Username == "" {
		return nil, ErrSecretUsernameNotSet
	}
	if secret.Password == "" {
		return nil, ErrSecretPasswordNotSet
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, ErrDbHostEnvNotSet
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, ErrDbPortEnvNotSet
	}
	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, ErrDbNameEnvNotSet
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     secret.Username,
		Password: secret.Password,
	}, nil
}

func NewDatabase(config *DatabaseConfig) (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Name, config.User, config.Password,
	)

	ctx := context.Background()

	db, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
