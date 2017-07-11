package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
)

func main() {
	schema, err := ioutil.ReadFile("schema.json")
	if err != nil {
		panic(err)
	}
	schemaLoader := gojsonschema.NewBytesLoader(schema)

	document, err := ioutil.ReadFile("document.json")
	if err != nil {
		panic(err)
	}
	documentLoader := gojsonschema.NewBytesLoader(document)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err)
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
