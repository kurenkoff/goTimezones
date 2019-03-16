package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestGetTime(t *testing.T) {
	data := []string{
		"Europe/Moscow",
		"America/Argentina/Salta",
	}

	for i := 0; i < len(data); i++ {
		time := GetTime(data[i])
		assert.NotEqual(t,"",time,"Empty response")
	}
}