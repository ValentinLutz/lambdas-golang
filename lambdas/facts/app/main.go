package main

import (
	"log/slog"
	"os"

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

	dbConfig, err := NewDatabaseConfigFromEnv()
	if err != nil {
		panic(err)
	}
	database, err = NewDatabase(dbConfig)
	if err != nil {
		panic(err)
	}

	lambda.Start(Handler)
}
