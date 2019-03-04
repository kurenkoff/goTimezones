package schema

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

type Validator struct {
	requestSchemaLoader gojsonschema.JSONLoader
	responseSchemaLoader gojsonschema.JSONLoader

}

func NewValidator() *Validator{
	return &Validator{
		gojsonschema.NewReferenceLoader("file:///home/serafim/go/src/timezones/schema/requestSchema.json"),
		gojsonschema.NewReferenceLoader("file:///home/serafim/go/src/timezones/schema/responseSchema.json"),
	}
}

func (v Validator) ValidateRequest(request string){
	requestLoader := gojsonschema.NewStringLoader(request)
	result, err := gojsonschema.Validate(v.requestSchemaLoader, requestLoader)
	if err != nil {
		fmt.Printf("ValidateRequest %v\n", err)
	}
	if result.Valid() {
		fmt.Printf("Requst is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func (v Validator) ValidateResponse(response string){
	responseLoader := gojsonschema.NewStringLoader(response)
	result, err := gojsonschema.Validate(v.responseSchemaLoader, responseLoader)
	if err != nil {
		fmt.Printf("ValidateResponse %v\n", err)
	}
	if result.Valid() {
		fmt.Printf("Response is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}