package connect

import (
	"context"
)

type ValidateConnectorConfigOptions struct {
	Config map[string]string
}

type ValidationResultConfig struct {
	Definition map[string]string `json:"definition"`
	Value      map[string]string `json:"value"`
}

type ConnectorValidationResult struct {
	Name       string                   `json:"name"`
	ErrorCount int                      `json:"error_count"`
	Groups     []string                 `json:"groups"`
	Configs    []ValidationResultConfig `json:"configs"`
}

func (c *Client) ValidateConnectorConfig(ctx context.Context, pluginClassName string, options ValidateConnectorConfigOptions) (ValidationResultConfig, error) {
	var validationResultConfig ValidationResultConfig
	response, err := c.client.NewRequest().
		SetContext(ctx).
		SetResult(&validationResultConfig).
		SetError(ApiError{}).
		SetPathParam("pluginClassName", pluginClassName).
		SetBody(options.Config).
		Put("/connector-plugins/{pluginClassName}/config/validate")
	if err != nil {
		return ValidationResultConfig{}, err
	}

	err = getErrorFromResponse(response)
	if err != nil {
		return ValidationResultConfig{}, err
	}

	return validationResultConfig, nil
}
