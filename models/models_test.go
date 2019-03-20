package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse_GetTime(t *testing.T) {
	testRequest := Request{1,[]string{"Europe/Moscow", "Europe/London", "Europe/Paris"}}
	testResponse := Response{make(map[string]string)}
	testResponse.GetTime(testRequest)
	assert.NotEqual(t,testResponse.TimeInZones["Europe/Moscow"],"","Empty answer")
}

func TestResponse_GetTimeP(t *testing.T) {
	testRequest := Request{1,[]string{"Europe/Moscow", "Europe/London", "Europe/Paris"}}
	testResponse := Response{make(map[string]string)}
	testResponse.GetTimeP(testRequest)
	assert.NotEqual(t,testResponse.TimeInZones["Europe/London"],"","Empty answer")
}