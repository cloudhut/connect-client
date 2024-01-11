package connect

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestStopConnector tests the success case for stopping connectors in a connect cluster.
func TestStopConnector(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", baseURL+"/connectors/datagen-product/stop", newJsonStringResponder(http.StatusAccepted, `{}`))
	err := c.StopConnector(context.Background(), "datagen-product")
	assert.NoError(t, err)
}
