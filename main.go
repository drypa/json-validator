package main

import (
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <schema-file> <json-file>")
		os.Exit(1)
	}

	schemaFile := os.Args[1]
	jsonFile := os.Args[2]

	schemaLoader := gojsonschema.NewReferenceLoader(schemaFile)
	documentLoader := gojsonschema.NewReferenceLoader(jsonFile)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Printf("Failed to validate document: %v", err)
		os.Exit(1)
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		os.Exit(1)
	}
}
