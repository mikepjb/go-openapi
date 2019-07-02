package main

import (
	"fmt"
	"os"

	"go-openapi/pkg/parse"
)

func main() {
	fmt.Println("go-openapi")

	if len(os.Args) < 2 {
		fmt.Println("usage: go-openapi ./openapi-description.yml")
		os.Exit(1)
	}

	description := parse.Parse(os.Args[1])

	fmt.Println(description.OpenAPI)
}
