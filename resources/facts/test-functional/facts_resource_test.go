package test_functional

import (
	"net/http"
	"root/libraries/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FactsResource(t *testing.T) {
	testConfig := testutil.LoadConfig("")
	database := testutil.NewDatabase(testConfig)
	testutil.LoadAndExec(database, "./truncate_tables.sql")

	t.Run(
		"get a random fact when no fact exists", func(t *testing.T) {
			// given

			// when
			resp, err := http.Get(testConfig.BaseURL + "/facts")

			// then
			assert.NoError(t, err)
			assert.Equal(t, 404, resp.StatusCode)
		},
	)

	t.Run(
		"get a random fact when only one fact exists", func(t *testing.T) {
			// given
			testutil.LoadAndExec(database, "./files/insert_fact.sql")

			// when
			resp, err := http.Get(testConfig.BaseURL + "/facts")

			// then
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)

			actualResponse := testutil.ReadBody(t, resp.Body)
			expectedResponse := testutil.ReadFile(t, "./files/get_facts_response.json")
			assert.JSONEq(t, string(expectedResponse), string(actualResponse))
		},
	)

	t.Run(
		"create a new fact", func(t *testing.T) {
			// given
			body := testutil.CreateBody(t, "./files/post_facts_response.json")

			// when
			resp, err := http.Post(testConfig.BaseURL+"/facts", "application/json", body)

			// then
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)

			actualResponse := testutil.ReadBody(t, resp.Body)
			expectedResponse := testutil.ReadFile(t, "./files/post_facts_response.json")
			assert.JSONEq(t, string(expectedResponse), string(actualResponse))
		},
	)

}
