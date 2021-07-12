package connect

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestGetConnectorConfig tests the success case for parsing the response for a GET request of a single
// connector's config.
func TestGetConnectorConfig(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors/datagen-product/config", newJsonStringResponder(http.StatusOK, `{"connector.class":"io.confluent.kafka.connect.datagen.DatagenConnector","quickstart":"product","tasks.max":"1","value.converter.schemas.enable":"false","name":"datagen-product","kafka.topic":"product","max.interval":"1000","iterations":"10000000"}`))
	info, err := c.GetConnectorConfig(context.Background(), "datagen-product")
	assert.NoError(t, err)
	assert.Equal(t, "io.confluent.kafka.connect.datagen.DatagenConnector", info["connector.class"])
	assert.Equal(t, "1000", info["max.interval"])
}

func TestGetConnectorConfigNotFound(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors/notfound-connector/config", newJsonStringResponder(http.StatusNotFound, `{"error_code":404,"message":"Connector notfound-connector not found"}`))
	info, err := c.GetConnectorConfig(context.Background(), "notfound-connector")
	assert.Equal(t, map[string]string(nil), info)
	require.Error(t, err)

	apiErr, ok := err.(ApiError)
	require.Equal(t, true, ok, "Returned err is not of type ApiError. Error's content is: %v", err.Error())
	assert.Equal(t, http.StatusNotFound, apiErr.ErrorCode)
	assert.Equal(t, "Connector notfound-connector not found", apiErr.Message)
}
