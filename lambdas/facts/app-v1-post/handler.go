package main

import (
	"context"
	"encoding/json"
	"log/slog"
	shared "root/lambdas/facts/app-shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Database *sqlx.DB
}

func (app *App) Handler(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var factRequest shared.FactRequest
	err := json.Unmarshal([]byte(r.Body), &factRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	factEntity := shared.FactEntity{
		Text: factRequest.Text,
	}
	_, err = app.Database.NamedExecContext(ctx, "INSERT INTO public.fact (fact_text) VALUES (:fact_text)", factEntity)
	if err != nil {
		slog.Error("failed to insert fact", slog.Any("err", err))
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	factResponseBody, err := json.Marshal(shared.FactResponse{Text: factRequest.Text})
	return events.APIGatewayProxyResponse{
		Body:       string(factResponseBody),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
