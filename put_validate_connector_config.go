package connect

import (
	"context"
)

type ValidateConnectorConfigOptions struct {
	Config map[string]interface{}
}

type ConnectorValidationResultConfig struct {
	Definition map[string]interface{} `json:"definition"`
	Value      map[string]interface{} `json:"value"`
}

type ConnectorValidationResult struct {
	Name       string                            `json:"name"`
	ErrorCount int                               `json:"error_count"`
	Groups     []string                          `json:"groups"`
	Configs    []ConnectorValidationResultConfig `json:"configs"`
}

func (c *Client) PutValidateConnectorConfig(ctx context.Context, pluginClassName string, options ValidateConnectorConfigOptions) (ConnectorValidationResult, error) {
	var connectorValidationResult ConnectorValidationResult
	response, err := c.client.NewRequest().
		SetContext(ctx).
		SetResult(&connectorValidationResult).
		SetError(ApiError{}).
		SetPathParam("pluginClassName", pluginClassName).
		SetBody(options.Config).
		Put("/connector-plugins/{pluginClassName}/config/validate")
	if err != nil {
		return ConnectorValidationResult{}, err
	}

	err = getErrorFromResponse(response)
	if err != nil {
		return ConnectorValidationResult{}, err
	}

	return connectorValidationResult, nil
}
