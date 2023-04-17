package connect

import (
	"context"
	"strconv"
)

// RestartConnectorOptions describe the available options to restart connectors.
type RestartConnectorOptions struct {
	// Specifies whether to restart the connector instance and task instances
	// or just the connector instance. Default is false.
	IncludeTasks bool
	// Specifies whether to restart just the instances with a FAILED status
	// or all instances. Default is false.
	OnlyFailed bool
}

// // RestartConnector restart the connector.
func (c *Client) RestartConnector(ctx context.Context, connectorName string, options RestartConnectorOptions) error {
	response, err := c.client.NewRequest().
		SetContext(ctx).
		SetError(ApiError{}).
		SetPathParam("connector", connectorName).
		SetQueryParams(map[string]string{
			"includeTasks": strconv.FormatBool(options.IncludeTasks),
			"onlyFailed":   strconv.FormatBool(options.OnlyFailed),
		}).
		Post("/connectors/{connector}/restart")
	if err != nil {
		return err
	}

	err = getErrorFromResponse(response)
	if err != nil {
		return err
	}

	return nil
}
