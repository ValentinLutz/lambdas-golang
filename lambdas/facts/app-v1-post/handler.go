package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Database *sqlx.DB
}

func (app *App) Handler(ctx context.Context, _ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: Implement POST handler
	slog.Info("POST handler not implemented")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
