package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct{
	Data string
	Expected string
}

func TestNewValidator(t *testing.T) {
	val := NewCustomValidator("requestSchema.json", "responseSchema.json")

	assert.NotNil(t, val, "Nil pointer")

}

func TestValidator_ValidateRequest(t *testing.T) {
	val := NewCustomValidator("requestSchema.json", "responseSchema.json")
	testData := []TestData{
		{
			"{\"id\":1, \"timezones\": [\"Europe/Moscow\"] } ",
			"",

		},
		{
			"{}",
			"ValidateRequest: request is not valid",

		},
		{
			"\"id\": 200",
			"ValidateRequest: request is not valid",
		},
	}

	for i := 0; i < len(testData); i++  {
		acctual := val.ValidateRequest(testData[i].Data)
		if testData[i].Expected == ""{
			assert.Nil(t, acctual, "Error isn't nil")
		} else{
			assert.EqualError(t, acctual, testData[i].Expected, "Error")
		}
	}



}

func TestValidator_ValidateResponse(t *testing.T) {
	val := NewCustomValidator("requestSchema.json", "responseSchema.json")
	testData := []TestData{
		{
			"{\"timezones\":{\"America/Argentina/Salta\":\"\"," +
				"\"Europe/Moscow\":\"2019-03-14T13:54:23.818762+03:00\"}}",
			"",

		},
		{
			"{}",
			"ValidateResponse: response is not valid",

		},
	}
	for i := 0; i < len(testData); i++  {
		acctual := val.ValidateResponse(testData[i].Data)
		if testData[i].Expected == ""{
			assert.Nil(t, acctual, "Error isn't nil")
		} else{
			assert.EqualError(t, acctual, testData[i].Expected, "Error")
		}
	}

}