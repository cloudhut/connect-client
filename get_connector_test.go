package connect

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestGetConnector tests the success case for parsing the response for a GET request of a single connector.
func TestGetConnector(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors/datagen-product", newJsonStringResponder(http.StatusOK, `{"name":"datagen-product","config":{"connector.class":"io.confluent.kafka.connect.datagen.DatagenConnector","quickstart":"product","tasks.max":"1","value.converter.schemas.enable":"false","name":"datagen-product","kafka.topic":"product","max.interval":"1000","iterations":"10000000"},"tasks":[{"connector":"datagen-product","task":0}],"type":"source"}`))
	info, err := c.GetConnector(context.Background(), "datagen-product")
	assert.NoError(t, err)
	assert.Equal(t, "datagen-product", info.Name)
}

// TestGetConnectorNotFound tests whether the API error for a non existent connector can be parsed successfully.
func TestGetConnectorNotFound(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors/notfound-connector", newJsonStringResponder(http.StatusNotFound, `{"error_code":404,"message":"Connector notfound-connector not found"}`))
	info, err := c.GetConnector(context.Background(), "notfound-connector")
	assert.Equal(t, info, ConnectorInfo{})
	require.Error(t, err)

	apiErr, ok := err.(ApiError)
	require.Equal(t, true, ok, "Returned err is not of type ApiError. Error's content is: %v", err.Error())
	assert.Equal(t, http.StatusNotFound, apiErr.ErrorCode)
	assert.Equal(t, "Connector notfound-connector not found", apiErr.Message)
}

// TestGetConnectorNotFoundNoAPIErr tests whether the client can handle 404 responses that do not return the expected
// JSON error. This would happen if the base url is wrong and another application such as a reverse proxy answers
// the request.
func TestGetConnectorNotFoundNoAPIErr(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors/notfound-connector", httpmock.NewStringResponder(http.StatusNotFound, `404 - not found`))
	info, err := c.GetConnector(context.Background(), "notfound-connector")
	assert.Equal(t, info, ConnectorInfo{})
	require.Error(t, err)

	_, ok := err.(ApiError)
	require.Equal(t, false, ok, "Returned error should not be of type ApiErr")
}
