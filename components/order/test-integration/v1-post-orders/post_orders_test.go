package post_orders_test

import (
	"context"
	"encoding/json"
	"root/components/order/test-integration"
	"root/libraries/testutil"
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

			td.CmpJSON(
				t, respBody, "./post_orders_response.json", []any{
					td.Re("^[A-Z0-9]{13}-EU-[A-Z0-9]{13}$"),
					td.NotEmpty(),
				},
			)
		},
	)

}
