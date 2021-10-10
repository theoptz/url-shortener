package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBytes(t *testing.T) {
	str := "example"
	expected := []byte(str)

	res := GetBytes(str)

	assert.Equal(t, expected, res)
}

func TestGetString(t *testing.T) {
	data := []byte("example")
	expected := string(data)

	res := GetString(data)

	assert.Equal(t, expected, res)
}
