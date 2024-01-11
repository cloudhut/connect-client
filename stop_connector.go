package connect

import (
	"context"
)

func (c *Client) StopConnector(ctx context.Context, connectorName string) error {
	response, err := c.client.NewRequest().
		SetContext(ctx).
		SetError(ApiError{}).
		SetPathParam("connector", connectorName).
		Put("/connectors/{connector}/stop")
	if err != nil {
		return err
	}

	err = getErrorFromResponse(response)
	if err != nil {
		return err
	}

	return nil
}
