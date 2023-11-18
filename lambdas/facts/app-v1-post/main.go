package main

import (
	"log/slog"
	"os"
	"root/libraries/apputil"

	"github.com/aws/aws-lambda-go/lambda"
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

	dbConfig, err := apputil.NewDatabaseConfigFromEnv()
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
