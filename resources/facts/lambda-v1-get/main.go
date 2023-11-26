package main

import (
	"context"
	"log/slog"
	"os"
	"root/libraries/apputil"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			},
		),
	)
	slog.SetDefault(logger)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	secret, err := apputil.GetSecret(cfg, os.Getenv("DB_SECRET_ID"))
	if err != nil {
		panic(err)
	}

	dbConfig, err := apputil.NewDatabaseConfig(secret)
	if err != nil {
		panic(err)
	}
	database, err := apputil.NewDatabase(dbConfig)
	if err != nil {
		panic(err)
	}

	app := &App{
		Database: database,
	}

	lambda.Start(app.Handler)
}
