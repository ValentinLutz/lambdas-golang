package post_orders_test

import (
	"context"
	"encoding/json"
	"root/libraries/testutil"
	"root/services/order/test-integration"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/maxatome/go-testdeep/td"
)

func Test_Post_Orders(t *testing.T) {
	t.Run(
		"create a new order", func(t *testing.T) {
			// given
			handler := testintegration.V1PostOrdersHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")
			body := testutil.ReadFile(t, "./post_orders_request.json")
			req := events.APIGatewayProxyRequest{
				Body: string(body),
			}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			td.CmpNoError(t, err)
			td.Cmp(t, resp.StatusCode, 200)

			var respBody map[string]interface{}
			err = json.Unmarshal([]byte(resp.Body), &respBody)
			td.CmpNoError(t, err)

			expectedResponse := testutil.ReadFile(t, "./post_orders_response.json")
			td.CmpJSON(
				t, respBody, expectedResponse, []any{
					td.Re("^[A-Za-z0-9]{13}-[A-Z]{2,4}-[A-Za-z0-9]{13}$"),
					td.NotEmpty(),
				},
			)
		},
	)

}
