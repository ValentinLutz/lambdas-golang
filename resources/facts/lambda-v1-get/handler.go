package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	shared "root/resources/facts/lambda-shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Database *sqlx.DB
}

func (app *App) Handler(ctx context.Context, _ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var factEntities []shared.FactEntity
	err := app.Database.SelectContext(ctx, &factEntities, "SELECT fact_id, fact_text FROM facts_resource.fact")
	if err != nil {
		slog.Error(
			"failed to select facts",
			slog.Any("err", err),
		)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	if len(factEntities) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
		}, nil
	}

	randomFactEntity := factEntities[rand.Intn(len(factEntities))]
	randomFactBody, err := json.Marshal(shared.FactResponse{Text: randomFactEntity.Text})
	if err != nil {
		slog.Error(
			"failed to marshal random fact",
			slog.Any("err", err),
		)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(randomFactBody),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}