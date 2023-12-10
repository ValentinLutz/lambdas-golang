package order_v1_get_test

import (
	"context"
	"root/libraries/testutil"
	testintegration "root/services/order/test-integration"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Test_orderResource_Get(t *testing.T) {
	t.Run(
		"get a random fact when no fact exists", func(t *testing.T) {
			// given
			handler := testintegration.orderV1GetHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")

			// when
			resp, err := handler.Invoke(context.Background(), events.APIGatewayProxyRequest{})

			// then
			assert.NoError(t, err)
			assert.Equal(t, 404, resp.StatusCode)
		},
	)

	t.Run(
		"get a random fact when only one fact exists", func(t *testing.T) {
			// given
			handler := testintegration.orderV1GetHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")
			testutil.LoadAndExec(handler.Database, "./insert_fact.sql")

			// when
			resp, err := handler.Invoke(context.Background(), events.APIGatewayProxyRequest{})

			// then
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)

			expectedResponse := testutil.ReadFile(t, "./get_order_response.json")
			assert.JSONEq(t, string(expectedResponse), resp.Body)
		},
	)
}
