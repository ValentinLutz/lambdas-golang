package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jmoiron/sqlx"
)

type fact struct {
	Text string `json:"text"`
}

type factEntity struct {
	//ID   int    `db:"fact_id"`
	Text string `db:"fact_text"`
}

type App struct {
	Database *sqlx.DB
}

func (app *App) Handler(ctx context.Context, _ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var factEntities []factEntity
	err := app.Database.SelectContext(ctx, &factEntities, "SELECT fact_text FROM public.fact")
	if err != nil {
		slog.Error(
			"failed to select facts",
			slog.Any("err", err),
		)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	rand.Intn(len(factEntities))
	randomFactEntity := factEntities[rand.Intn(len(factEntities))]
	randomFactBody, err := json.Marshal(fact{Text: randomFactEntity.Text})
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
