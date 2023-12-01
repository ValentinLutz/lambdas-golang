package facts_v1_post_test

import (
	"context"
	"root/libraries/testutil"
	testintegration "root/resources/facts/test-integration"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Test_FactsResource_Post(t *testing.T) {
	t.Run(
		"create a new fact", func(t *testing.T) {
			// given
			handler := testintegration.FactsV1PostHandler
			testutil.LoadAndExec(handler.Database, "../truncate_tables.sql")
			body := testutil.ReadFile(t, "./post_facts_response.json")
			req := events.APIGatewayProxyRequest{
				Body: string(body),
			}

			// when
			resp, err := handler.Invoke(context.Background(), req)

			// then
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)

			expectedResponse := testutil.ReadFile(t, "./post_facts_response.json")
			assert.JSONEq(t, string(expectedResponse), resp.Body)
		},
	)

}
