package incoming

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"root/libraries/apputil"
	"root/services/order/lambda-shared/incoming"
	"root/services/order/lambda-v1-post-orders/core"
	"root/services/order/lambda-v1-post-orders/outgoing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Database     *sqlx.DB
	OrderService *core.OrderService
}

var (
	ErrDbSecretIdEnvNotSet  = fmt.Errorf("env DB_SECRET_ID not set")
	ErrOrderRegionEnvNotSet = fmt.Errorf("env ORDER_REGION not set")
)

func NewHandler() (*Handler, error) {
	apputil.NewSlogDefault(slog.LevelInfo)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to load aws default config: %w", err)
	}

	DbSecretIdEnv, ok := os.LookupEnv("DB_SECRET_ID")
	if !ok {
		return nil, ErrDbSecretIdEnvNotSet
	}

	secret, err := apputil.GetSecret(cfg, DbSecretIdEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to get database secret: %w", err)
	}

	dbConfig, err := apputil.NewDatabaseConfig(secret)
	if err != nil {
		return nil, fmt.Errorf("failed to create database config: %w", err)
	}

	database, err := apputil.NewDatabase(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	regionEnv, ok := os.LookupEnv("ORDER_REGION")
	if !ok {
		return nil, ErrOrderRegionEnvNotSet
	}

	orderRepository := outgoing.NewOrderRepository(database)
	region, err := core.NewRegion(regionEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create region: %w", err)
	}
	orderService := core.NewOrderService(region, orderRepository)

	return &Handler{
		Database:     database,
		OrderService: orderService,
	}, nil
}

func (handler *Handler) Invoke(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var orderRequest incoming.OrderRequest
	err := json.Unmarshal([]byte(request.Body), &orderRequest)
	if err != nil {
		slog.Error("failed to unmarshal order request", apputil.ErrorAttr(err))

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	orderResponse, err := handler.OrderService.PlaceOrder(ctx, orderRequest)
	if err != nil {
		slog.Error("failed to place order", apputil.ErrorAttr(err))

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	orderResponseBody, err := json.Marshal(orderResponse)
	if err != nil {
		slog.Error("failed to marshal orders response", apputil.ErrorAttr(err))

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(orderResponseBody),
	}, nil
}
