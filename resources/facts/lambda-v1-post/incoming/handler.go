package incoming

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"root/libraries/apputil"
	shared "root/resources/facts/lambda-shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Database *sqlx.DB
}

func NewHandler() *Handler {
	apputil.NewSlogDefault(slog.LevelInfo)

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

	return &Handler{
		Database: database,
	}
}

type FactRequest struct {
	Text string `json:"text"`
}

type FactResponse struct {
	Text string `json:"text"`
}

func (handler *Handler) Invoke(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var factRequest FactRequest
	err := json.Unmarshal([]byte(r.Body), &factRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	factEntity := shared.FactEntity{
		Text: factRequest.Text,
	}
	_, err = handler.Database.NamedExecContext(ctx, "INSERT INTO facts_resource.fact (fact_text) VALUES (:fact_text)", factEntity)
	if err != nil {
		slog.Error("failed to insert fact", slog.Any("err", err))
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	factResponseBody, err := json.Marshal(FactResponse{Text: factRequest.Text})
	return events.APIGatewayProxyResponse{
		Body:       string(factResponseBody),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
