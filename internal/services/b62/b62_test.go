package b62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCode(t *testing.T) {
	code := NewCode()

	assert.Equal(t, &Code{}, code)
}

func TestCode_Encode(t *testing.T) {
	testCases := []struct {
		val int64
		res string
	}{
		{
			val: -1,
			res: "",
		},
		{
			val: 0,
			res: "0",
		},
		{
			val: 62,
			res: "10",
		},
		{
			val: 62 * 62,
			res: "100",
		},
		{
			val: 1000,
			res: "g8",
		},
		{
			val: 396548739,
			res: "qPSrp",
		},
		{
			val: 1 << 24,
			res: "18owg",
		},
	}

	for _, tc := range testCases {
		code := Code{}
		res := code.Encode(tc.val)
		assert.Equal(t, tc.res, res, "should return %s for %d", tc.res, tc.val)
	}
}

func TestCode_Decode(t *testing.T) {
	testCases := []struct {
		val string
		res int64
	}{
		{
			val: "",
			res: -1,
		},
		{
			val: "a]",
			res: -1,
		},
		{
			val: "0",
			res: 0,
		},
		{
			val: "10",
			res: 62,
		},
		{
			val: "100",
			res: 62 * 62,
		},
		{
			val: "g8",
			res: 1000,
		},
		{
			val: "qPSrp",
			res: 396548739,
		},
		{
			val: "18owg",
			res: 1 << 24,
		},
	}

	for _, tc := range testCases {
		code := Code{}
		res := code.Decode(tc.val)
		assert.Equal(t, tc.res, res, "should return %d for %s", tc.res, tc.val)
	}
}
