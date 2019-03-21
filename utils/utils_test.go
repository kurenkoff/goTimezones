package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTime(t *testing.T) {
	testData := []struct {
		location string
		expected string
	}{
		{
			"Europe/Moscow",
			"",
		},
		{
			"/dust",
			"bad status code",
		},
	}

	for i := 0; i < len(testData); i++ {
		_, err := GetTime(testData[i].location)
		assert.Equal(t, testData[i].expected, err.Error(), "unexpected error")
	}
}
