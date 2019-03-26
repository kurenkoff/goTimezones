package schema

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
)

// Validator структура, необходимая для валидации запроса и ответа
type Validator struct {
	// JsonSchema запроса
	requestSchemaLoader  gojsonschema.JSONLoader
	// JsonSchema ответа
	responseSchemaLoader gojsonschema.JSONLoader
}

// NewValidator возвращает новый валидатор
func NewValidator() *Validator {
	// загрузка из файла схемы запроса
	buf, err := ioutil.ReadFile("schema/requestSchema.json")
	if err != nil {
		log.Fatal(fmt.Errorf("can't find requestSchema.json"))
	}
	request := string(buf)

	// загрузка из файла схемы ответа
	buf, err = ioutil.ReadFile("schema/responseSchema.json")
	if err != nil {
		log.Fatal(fmt.Errorf("can't find responseSchema.json"))
	}
	response := string(buf)

	return &Validator{
		gojsonschema.NewStringLoader(request),
		gojsonschema.NewStringLoader(response),
	}
}

// NewCustomValidator возвращает новый валидатор.
// NewCustomValidator Читает Json схемы из файлов расположенных по путям requestPath и responsePath
func NewCustomValidator(requestPath string, responsePath string) *Validator {
	// загрузка из файла схемы запроса
	buf, err := ioutil.ReadFile(requestPath)
	if err != nil {
		log.Fatal(fmt.Errorf("can't find requestSchema.json"))
	}
	request := string(buf)

	// загрузка из файла схемы ответа
	buf, err = ioutil.ReadFile(responsePath)
	if err != nil {
		log.Fatal(fmt.Errorf("can't find responseSchema.json"))
	}
	response := string(buf)

	return &Validator{
		gojsonschema.NewStringLoader(request),
		gojsonschema.NewStringLoader(response),
	}
}

// ValidateRequest производит валидацию запроса по схеме записанной в поле v.responseSchemaLoader
func (v Validator) ValidateRequest(request string) error {
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

// ValidateResponse производит валидацию ответа по схеме записанной в поле v.responseSchemaLoader
func (v Validator) ValidateResponse(response string) error {
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
