package connect

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestResetConnectorTopics tests the success case for resetting connector topics.
func TestResetConnectorTopics(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", baseURL+"/connectors/datagen-product/topics/reset", newJsonStringResponder(http.StatusOK, `{}`))
	err := c.ResetConnectorTopics(context.Background(), "datagen-product")
	assert.NoError(t, err)
}
