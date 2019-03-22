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

func TestGetTimeP(t *testing.T) {
	chn := make(chan PTime)
	go GetTimeP("Europe/Moscow", chn)
	go GetTimeP("Europe/SHDhas", chn)

	test := make(map[string]string)
	for i := 0; i < 2; i++ {
		tmp := <-chn
		test[tmp.Zone] = tmp.Time
	}
	assert.NotEqual(t, test["Europe/Moscow"], "", "No time in Moscow")
	assert.Equal(t, test["Europe/SHDhas"], "", "Europe/SHDhas exist")
}
