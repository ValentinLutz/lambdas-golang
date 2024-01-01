package incoming

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"root/libraries/apputil"
	"root/services/order/lambda-v1-get-orders/core"
	"root/services/order/lambda-v1-get-orders/outgoing"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/google/uuid"
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
	offset := 0
	offsetString, ok := request.QueryStringParameters["offset"]
	if ok {
		parsedOffset, err := strconv.Atoi(offsetString)
		if err != nil {
			slog.Error("failed to parse offset", apputil.ErrorAttr(err))

			return events.APIGatewayProxyResponse{
				StatusCode: 400,
			}, nil
		}
		offset = parsedOffset
	}

	limit := 50
	limitString, ok := request.QueryStringParameters["limit"]
	if ok {
		parsedLimit, err := strconv.Atoi(limitString)
		if err != nil {
			slog.Error("failed to parse limit", apputil.ErrorAttr(err))

			return events.APIGatewayProxyResponse{
				StatusCode: 400,
			}, nil
		}
		limit = parsedLimit
	}

	var customerId *uuid.UUID
	customerIdString, ok := request.QueryStringParameters["customer_id"]
	if ok {
		parsedCustomerId, err := uuid.Parse(customerIdString)
		if err != nil {
			slog.Error("failed to parse customerId", apputil.ErrorAttr(err))

			return events.APIGatewayProxyResponse{
				StatusCode: 400,
			}, nil
		}
		customerId = &parsedCustomerId
	}

	ordersResponse, err := handler.OrderService.GetOrders(ctx, offset, limit, customerId)
	if err != nil {
		slog.Error("failed to get orders", apputil.ErrorAttr(err))

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	ordersResponseBody, err := json.Marshal(ordersResponse)
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
		Body: string(ordersResponseBody),
	}, nil
}
