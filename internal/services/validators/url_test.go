package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewURLValidator(t *testing.T) {
	v := NewURLValidator()
	expected := &URLValidator{}

	assert.Equal(t, expected, v)
}

func TestURLValidator_Validate(t *testing.T) {
	testCases := []struct {
		val string
		res bool
	}{
		{
			val: "http://example.com/",
			res: true,
		},
		{
			val: "https://example.com/example.php",
			res: true,
		},
		{
			val: "//example.com/",
			res: false,
		},
		{
			val: "example.com",
			res: false,
		},
		{
			val: "/example/com",
			res: false,
		},
		{
			val: "",
			res: false,
		},
		{
			val: "https://",
			res: false,
		},
	}

	for _, tc := range testCases {
		validator := NewURLValidator()
		res := validator.Validate(tc.val)

		assert.Equal(t, tc.res, res)
	}
}
