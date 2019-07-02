package parse

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Description represents the OpenAPI file that describes your REST endpoint.
type Description struct {
	OpenAPI string `yaml:"openapi"`
}

func Parse(filename string) Description {
	exampleBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("problem reading file %v\n")
		os.Exit(1)
	}

	var description Description

	err = yaml.Unmarshal(exampleBytes, &description)

	return description
}
