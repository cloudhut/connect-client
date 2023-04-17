package connect

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestRestartConnectors tests the success case for restarting connectors in a connect cluster.
func TestRestartConnectors(t *testing.T) {
	baseURL := "https://fake.api"

	tests := []struct {
		title   string
		url     string
		options RestartConnectorOptions
	}{
		{
			title:   "default",
			url:     baseURL + "/connectors/my-connector/restart",
			options: RestartConnectorOptions{},
		},
		{
			title:   "includeTasks=false",
			url:     baseURL + "/connectors/my-connector/restart?includeTasks=false&onlyFailed=false",
			options: RestartConnectorOptions{IncludeTasks: false},
		},
		{
			title:   "onlyFailed=false",
			url:     baseURL + "/connectors/my-connector/restart?includeTasks=false&onlyFailed=false",
			options: RestartConnectorOptions{OnlyFailed: false},
		},
		{
			title:   "both false",
			url:     baseURL + "/connectors/my-connector/restart?includeTasks=false&onlyFailed=false",
			options: RestartConnectorOptions{IncludeTasks: false, OnlyFailed: false},
		},
		{
			title:   "include tasks",
			url:     baseURL + "/connectors/my-connector/restart?includeTasks=true&onlyFailed=false",
			options: RestartConnectorOptions{IncludeTasks: true},
		},
		{
			title:   "include only failed",
			url:     baseURL + "/connectors/my-connector/restart?includeTasks=false&onlyFailed=true",
			options: RestartConnectorOptions{OnlyFailed: true},
		},
		{
			title:   "include both",
			url:     baseURL + "/connectors/my-connector/restart?includeTasks=true&onlyFailed=true",
			options: RestartConnectorOptions{IncludeTasks: true, OnlyFailed: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			c := NewClient(WithHost(baseURL))

			httpmock.ActivateNonDefault(c.client.GetClient())
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("POST", tt.url,
				newJsonStringResponder(http.StatusOK,
					`{"name": "my-connector","connector": {"state": "RUNNING","worker_id": "fakehost1:8083"},"tasks":[{"id": 0,"state": "RUNNING","worker_id": "fakehost2:8083"},{"id": 1,"state": "RESTARTING","worker_id": "fakehost3:8083"},{"id": 2,"state": "RESTARTING","worker_id": "fakehost1:8083"}]}`))

			err := c.RestartConnector(context.Background(), "my-connector", tt.options)
			assert.NoError(t, err)
		})
	}
}
