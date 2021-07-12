package connect

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestListConnectors tests the success case for listing all connectors in a connect cluster.
func TestListConnectors(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors", newJsonStringResponder(http.StatusOK, `["datagen-product"]`))
	connectors, err := c.ListConnectors(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, []string{"datagen-product"}, connectors)
}

// TestListConnectorsExpanded tests the success case for listing all connectors with expanded info (info and status)
// in a connect cluster.
func TestListConnectorsExpanded(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors?expand=info&expand=status", newJsonStringResponder(http.StatusOK, `{"datagen-product":{"info":{"name":"datagen-product","config":{"connector.class":"io.confluent.kafka.connect.datagen.DatagenConnector","quickstart":"product","tasks.max":"1","value.converter.schemas.enable":"false","name":"datagen-product","kafka.topic":"product","max.interval":"1000","iterations":"10000000"},"tasks":[{"connector":"datagen-product","task":0}],"type":"source"},"status":{"name":"datagen-product","connector":{"state":"RUNNING","worker_id":"connect:8083"},"tasks":[{"id":0,"state":"RUNNING","worker_id":"connect:8083"}],"type":"source"}}}`))
	connectors, err := c.ListConnectorsExpanded(context.Background())
	assert.NoError(t, err)
	require.Len(t, connectors, 1)

	mockConnector := connectors["datagen-product"]
	require.NotNil(t, mockConnector)
	assert.Equal(t, mockConnector.Info.Name, "datagen-product")

	require.NotNil(t, mockConnector.Info.Config)
	assert.Equal(t, mockConnector.Info.Config["connector.class"], "io.confluent.kafka.connect.datagen.DatagenConnector")

	assert.Equal(t, mockConnector.Status.Name, "datagen-product")
	assert.Equal(t, mockConnector.Status.Connector.State, "RUNNING")
}

// TestListConnectorsExpandedEmpty tests the success case for listing all connectors with expanded info (info and status)
// in a connect cluster. However in this specific case the response is an empty object because no connectors may be existent.
func TestListConnectorsExpandedEmpty(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", baseURL+"/connectors?expand=info&expand=status", newJsonStringResponder(http.StatusOK, `{}`))
	connectors, err := c.ListConnectorsExpanded(context.Background())
	assert.NoError(t, err)
	assert.Len(t, connectors, 0)
}
