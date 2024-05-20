package order_v1_get_test

import (
	"context"
	"encoding/json"
	testintegration "root/components/order/test-integration"
	"root/libraries/testutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/maxatome/go-testdeep/td"
)

func Test_Get_Order(t *testing.T) {
	t.Run(
		"get order", func(t *testing.T) {
			// given
			handler := testintegration.V1GetOrderHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")
			testutil.LoadAndExec(handler.Database, "./init_get_order.sql")
			req := events.APIGatewayProxyRequest{
				PathParameters: map[string]string{
					"order_id": "01HK5W7JF32CM-EU-0GVJSF5RFM7PN",
				},
			}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			td.CmpNoError(t, err)
			td.Cmp(t, resp.StatusCode, 200)

			var respBody map[string]interface{}
			err = json.Unmarshal([]byte(resp.Body), &respBody)
			td.CmpNoError(t, err)

			td.CmpJSON(
				t, respBody, "./get_order_response.json", []any{},
			)
		},
	)

	t.Run(
		"get order that does not exists", func(t *testing.T) {
			// given
			handler := testintegration.V1GetOrderHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")
			req := events.APIGatewayProxyRequest{
				PathParameters: map[string]string{
					"order_id": "01HK5W7JF32CM-EU-0GVJSF5RFM7PN",
				},
			}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			td.CmpNoError(t, err)
			td.Cmp(t, resp.StatusCode, 404)
		},
	)

	t.Run(
		"get order with no order id", func(t *testing.T) {
			// given
			handler := testintegration.V1GetOrderHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")
			req := events.APIGatewayProxyRequest{}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			td.CmpNoError(t, err)
			td.Cmp(t, resp.StatusCode, 400)
		},
	)
}
