package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudhut/connect-client"
)

func main() {
	opts := []connect.ClientOption{
		connect.WithTimeout(5 * time.Second),
		connect.WithUserAgent("Redpanda Console"),
		connect.WithHost("http://localhost:8083"),
	}
	client := connect.NewClient(opts...)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 1. List all installed connector plugins
	fmt.Printf("### List Connector Plugins ###\n")
	plugins, err := client.GetConnectorPlugins(ctx)
	if err != nil {
		panic(err)
	}
	for _, plugin := range plugins {
		fmt.Printf("Class: %q\n", plugin.Class)
		fmt.Printf("Type: %q\n", plugin.Type)
		fmt.Printf("Version: %q\n", plugin.Version)
		fmt.Print("----------------------------\n")
	}

	fmt.Printf("\n\n### Validate Connector Config ###\n")

	// 2. Validate connector config
	validateOpts := connect.ValidateConnectorConfigOptions{
		Config: map[string]interface{}{
			"connector.class": "io.debezium.connector.postgresql.PostgresConnector",
		},
	}
	validationResults, err := client.PutValidateConnectorConfig(
		ctx,
		"io.debezium.connector.postgresql.PostgresConnector",
		validateOpts,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Validation results for %q\n", validationResults.Name)
	fmt.Printf("Validation errors: %d\n", validationResults.ErrorCount)
}
