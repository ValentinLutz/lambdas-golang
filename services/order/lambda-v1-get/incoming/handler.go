package incoming

import (
	"context"
	"log/slog"
	"os"
	"root/libraries/apputil"

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

type FactResponse struct {
	Text string `json:"text"`
}

func (handler *Handler) Invoke(ctx context.Context, _ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var factEntities []shared.FactEntity
	//err := handler.Database.SelectContext(ctx, &factEntities, "SELECT fact_id, fact_text FROM order_resource.fact")
	//if err != nil {
	//	slog.Error(
	//		"failed to select order",
	//		slog.Any("err", err),
	//	)
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 500,
	//	}, nil
	//}
	//
	//if len(factEntities) == 0 {
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 404,
	//	}, nil
	//}
	//
	//randomFactEntity := factEntities[rand.Intn(len(factEntities))]
	//randomFactBody, err := json.Marshal(FactResponse{Text: randomFactEntity.Text})
	//if err != nil {
	//	slog.Error(
	//		"failed to marshal random fact",
	//		slog.Any("err", err),
	//	)
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 500,
	//	}, nil
	//}
	//
	//return events.APIGatewayProxyResponse{
	//	Body:       string(randomFactBody),
	//	StatusCode: 200,
	//	Headers: map[string]string{
	//		"Content-Type": "application/json",
	//	},
	//}, nil
	return events.APIGatewayProxyResponse{}, nil
}
