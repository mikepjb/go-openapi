package parse_test

import (
	"fmt"
	"testing"

	"go-openapi/pkg/parse"
)

func TestOpenAPIVersion(t *testing.T) {
	description := parse.Parse("example.yml")

	if description.OpenAPI != "3.0.1" {
		fmt.Errorf("wrong version, got %v", description.OpenAPI)
	}
}
