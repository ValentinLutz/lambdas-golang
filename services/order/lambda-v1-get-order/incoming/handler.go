package incoming

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"root/libraries/apputil"
	"root/services/order/lambda-v1-get-order/core"
	"root/services/order/lambda-v1-get-order/outgoing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Database     *sqlx.DB
	OrderService *core.OrderService
}

func NewHandler() (*Handler, error) {
	apputil.NewSlogDefault(slog.LevelInfo)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to load aws default config: %w", err)
	}

	secret, err := apputil.GetSecret(cfg, os.Getenv("DB_SECRET_ID"))
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

	orderRepository := outgoing.NewOrderRepository(database)
	orderService := core.NewOrderService(orderRepository)

	return &Handler{
		Database:     database,
		OrderService: orderService,
	}, nil
}

func (handler *Handler) Invoke(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orderIdString, ok := request.PathParameters["order_id"]
	if !ok {
		slog.Error("order id not found in path parameters")

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	order, orderItems, err := handler.OrderService.GetOrder(ctx, orderIdString)
	if err != nil {
		if errors.Is(err, outgoing.ErrOrderNotFound) {
			slog.Error("order not found", apputil.ErrorAttr(err))

			return events.APIGatewayProxyResponse{
				StatusCode: 404,
			}, nil
		}

		slog.Error("failed to get order", apputil.ErrorAttr(err))

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	orderResponse := NewOrderResponse(order, orderItems)
	orderResponseBody, err := json.Marshal(orderResponse)
	if err != nil {
		slog.Error("failed to marshal order response", apputil.ErrorAttr(err))

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
