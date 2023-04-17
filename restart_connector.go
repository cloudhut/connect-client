package connect

import (
	"context"
	"strconv"
)

// ListConnectorsOptions describe the available options to list connectors. Either Status or Info must be set to true.
type RestartConnectorOptions struct {
	IncludeTasks bool
	OnlyFailed   bool
}

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
