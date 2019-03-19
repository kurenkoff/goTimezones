package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse_GetTime(t *testing.T) {
	testRequest := Request{1,[]string{"Europe/Moscow"}}
	testResponse := Response{make(map[string]string)}
	testResponse.GetTime(testRequest)
	assert.NotEqual(t,testResponse.TimeInZones["Europe/Moscow"],"","Empty answer")
}
