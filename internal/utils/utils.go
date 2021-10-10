package utils

import "unsafe"

func GetBytes(val string) []byte {
	return *(*[]byte)(unsafe.Pointer(&val))
}

func GetString(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}
