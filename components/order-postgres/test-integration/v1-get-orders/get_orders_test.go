package order_v1_get_test

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/maxatome/go-testdeep/td"
	testintegration "root/components/order-postgres/test-integration"
	"root/libraries/testutil"
	"testing"
)

func Test_Get_Orders(t *testing.T) {
	t.Run(
		"get orders for customer", func(t *testing.T) {
			// given
			handler := testintegration.V1GetOrdersHandler
			testutil.MustLoadAndExec(handler.Database, "../truncate_tables.sql")
			testutil.MustLoadAndExec(handler.Database, "./init_get_orders.sql")
			req := events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{
					"offset":       "0",
					"limit":        "10",
					"customer_key": "44bd6239-7e3d-4d4a-90a0-7d4676a00f5c",
				},
			}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			td.CmpNoError(t, err)
			td.Cmp(t, resp.StatusCode, 200)

			var respBody []map[string]interface{}
			err = json.Unmarshal([]byte(resp.Body), &respBody)
			td.CmpNoError(t, err)

			td.CmpJSON(
				t, respBody, "./get_orders_response.json", []any{},
			)
		},
	)

	t.Run(
		"get orders for customer when no orders exists for customer", func(t *testing.T) {
			// given
			handler := testintegration.V1GetOrdersHandler
			testutil.MustLoadAndExec(handler.Database, "../truncate_tables.sql")
			req := events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{
					"offset":       "0",
					"limit":        "10",
					"customer_key": "00000000-0000-0000-0000-000000000000",
				},
			}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			td.CmpNoError(t, err)
			td.Cmp(t, resp.StatusCode, 200)

			var respBody []map[string]interface{}
			err = json.Unmarshal([]byte(resp.Body), &respBody)
			td.CmpNoError(t, err)

			td.CmpJSON(
				t, respBody, "./get_orders_empty_response.json", []any{},
			)
		},
	)
}
