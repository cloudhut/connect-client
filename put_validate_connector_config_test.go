package connect

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// TestPutValidateConnectorConfigInvalidConfig tests the response of an invalid config.
func TestPutValidateConnectorConfigValidConfig(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	validateConnectorConfigOptions := ValidateConnectorConfigOptions{
		Config: map[string]interface{}{"connector.class": "com.github.cloudhut.kowl.RegisteredConnector", "name": "connectorName"},
	}

	responseBody := `
	{
	    "name": "com.github.cloudhut.kowl.RegisteredConnector",
	    "error_count": 0,
	    "groups": [],
	    "configs": []
	}`

	httpmock.RegisterResponder("PUT", baseURL+"/connector-plugins/registeredConnector/config/validate", newJsonStringResponder(http.StatusOK, responseBody))
	validationResult, err := c.PutValidateConnectorConfig(context.Background(), "registeredConnector", validateConnectorConfigOptions)
	assert.NoError(t, err)
	assert.Equal(t, 0, validationResult.ErrorCount)
	assert.Equal(t, []string{}, validationResult.Groups)
}

// TestPutValidateConnectorConfigInvalidConfig tests the response of an invalid config.
func TestPutValidateConnectorConfigInvalidConfig(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	validateConnectorConfigOptions := ValidateConnectorConfigOptions{
		Config: map[string]interface{}{"connector.class": "com.github.cloudhut.kowl.RegisteredConnector"},
	}

	responseBody := `
	{
	    "name": "com.github.cloudhut.kowl.RegisteredConnector",
	    "error_count": 1,
	    "groups": [
	        "Common"
	    ],
	    "configs": [
	        {
	            "definition": {
	                "name": "name",
	                "type": "STRING",
	                "required": true,
	                "default_value": null,
	                "importance": "HIGH",
	                "documentation": "Globally unique name to use for this connector.",
	                "group": "Common",
	                "width": "MEDIUM",
	                "display_name": "Connector name",
	                "dependents": [],
	                "order": 1
	            },
	            "value": {
	                "name": "name",
	                "value": null,
	                "recommended_values": [],
	                "errors": [
	                  "Missing required configuration \"name\" which has no default value."
	                ],
	                "visible": true
	            }
	        }
		]
	}`

	httpmock.RegisterResponder("PUT", baseURL+"/connector-plugins/registeredConnector/config/validate", newJsonStringResponder(http.StatusOK, responseBody))
	validationResult, err := c.PutValidateConnectorConfig(context.Background(), "registeredConnector", validateConnectorConfigOptions)
	assert.NoError(t, err)
	assert.Equal(t, 1, validationResult.ErrorCount)
	assert.Equal(t, []string{"Common"}, validationResult.Groups)
}

// TestPutValidateConnectorConfigInvalidConfig tests the response of a non-exising connect, i.e. an error and a default response.
func TestPutValidateConnectorConfigConnectorNotFound(t *testing.T) {
	baseURL := "https://fake.api"
	c := NewClient(WithHost(baseURL))

	httpmock.ActivateNonDefault(c.client.GetClient())
	defer httpmock.DeactivateAndReset()

	validateConnectorConfigOptions := ValidateConnectorConfigOptions{
		Config: map[string]interface{}{"connector.class": "com.github.cloudhut.kowl.RegisteredConnector", "name": "connectorName"},
	}

	httpmock.RegisterResponder("PUT", baseURL+"/connector-plugins/registeredConnector/config/validate", newJsonStringResponder(http.StatusNotFound, `notFound`))
	validationResult, err := c.PutValidateConnectorConfig(context.Background(), "registeredConnector", validateConnectorConfigOptions)
	assert.Error(t, err)
	assert.Equal(t, ConnectorValidationResult{}, validationResult)
}
