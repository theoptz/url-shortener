package b62

import (
	"strings"

	"github.com/theoptz/url-shortener/internal/utils"
)

const base = 62
const symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const aSmall = 'a'
const aBig = 'A'
const zSmall = 'z'
const zBig = 'Z'

const zeroDigit = '0'
const nineDigit = '9'

const aSmallPosition = 10
const aBigPosition = 36

type Code struct{}

func (c *Code) Encode(val int64) string {
	if val < 0 {
		return ""
	} else if val == 0 {
		return "0"
	}

	sb := strings.Builder{}

	for val > 0 {
		sb.WriteByte(symbols[val%base])

		val /= base
	}

	res := sb.String()
	byteArr := utils.GetBytes(res)

	l := len(byteArr)
	for i := 0; i < l/2; i++ {
		byteArr[i], byteArr[l-1-i] = byteArr[l-1-i], byteArr[i]
	}

	return utils.GetString(byteArr)
}

func (c *Code) Decode(val string) int64 {
	var res int64

	bytesData := utils.GetBytes(val)
	l := len(bytesData)
	if l == 0 {
		return -1
	}

	var multiply int64 = 1

	for i := l - 1; i >= 0; i-- {
		if bytesData[i] >= aSmall && bytesData[i] <= zSmall {
			res += multiply * int64(bytesData[i]-aSmall+aSmallPosition)
		} else if bytesData[i] >= aBig && bytesData[i] <= zBig {
			res += multiply * int64(bytesData[i]-aBig+aBigPosition)
		} else if bytesData[i] >= zeroDigit && bytesData[i] <= nineDigit {
			res += multiply * int64(bytesData[i]-zeroDigit)
		} else {
			// invalid byte, force returning -1
			return -1
		}

		multiply *= base
	}

	return res
}

func NewCode() *Code {
	return &Code{}
}
