package schema

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
)

type Validator struct {
	requestSchemaLoader gojsonschema.JSONLoader
	responseSchemaLoader gojsonschema.JSONLoader

}

func NewValidator() *Validator{
	// загрузка из файла схемы запроса
	buf, err := ioutil.ReadFile("schema/requestSchema.json")
	if err != nil{
		log.Fatal(fmt.Errorf("can't find requestSchema.json"))
	}
	request := string(buf)

	// загрузка из файла схемы ответа
	buf, err = ioutil.ReadFile("schema/responseSchema.json")
	if err != nil{
		log.Fatal(fmt.Errorf("can't find requestSchema.json"))
	}
	response := string(buf)

	return &Validator{
		gojsonschema.NewStringLoader(request),
		gojsonschema.NewStringLoader(response),
	}
}

func (v Validator) ValidateRequest(request string) error{
	requestLoader := gojsonschema.NewStringLoader(request)
	result, err := gojsonschema.Validate(v.requestSchemaLoader, requestLoader)
	if err != nil {
		fmt.Printf("ValidateRequest %v\n", err)
	}
	if result.Valid() {
		fmt.Printf("Requst is valid\n")
	} else {
		return fmt.Errorf("ValidateRequest: request is not valid")
	}
	return nil
}

func (v Validator) ValidateResponse(response string) error{
	responseLoader := gojsonschema.NewStringLoader(response)
	result, err := gojsonschema.Validate(v.responseSchemaLoader, responseLoader)
	if err != nil {
		fmt.Printf("ValidateResponse %v\n", err)
	}
	if result.Valid() {
		fmt.Printf("Response is valid\n")
	} else {
		return fmt.Errorf("ValidateResponse: response is not valid")
	}
	return nil
}