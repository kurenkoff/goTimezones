package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/time", GetTimeZone)
	return r
}

/*
	TEST DATA
	1. BODY {"id": 1, "timezones":["", ""]} EXPECTED STATUS CODE 200
	2. BODY "" EXPECTED STATUS CODE 422
*/

func TestGetTimeZone(t *testing.T) {
	testData := []struct {
		body     string
		expected int
	}{
		{
			`{
						"id": 1, 
						"timezones": [
							"America/Argentina/Salta",
        					"Europe/Moscow"
    					]
					}`,
			http.StatusOK,
		},
		{
			"",
			http.StatusUnprocessableEntity,
		},
	}

	for i := 0; i < len(testData); i++ {
		request, err := http.NewRequest("GET", "/time", bytes.NewBufferString(testData[i].body))
		if err != nil {
			t.Errorf(err.Error())
		}
		recoder := httptest.NewRecorder()
		http.HandlerFunc(GetTimeZone).ServeHTTP(recoder, request)
		assert.Equal(t, testData[i].expected, recoder.Code, fmt.Sprintf("Expected %d, got %d", testData[i].expected, recoder.Code))
	}
}
