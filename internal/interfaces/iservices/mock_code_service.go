// Code generated by mockery 2.9.4. DO NOT EDIT.

package iservices

import mock "github.com/stretchr/testify/mock"

// MockCodeService is an autogenerated mock type for the CodeService type
type MockCodeService struct {
	mock.Mock
}

// Decode provides a mock function with given fields: _a0
func (_m *MockCodeService) Decode(_a0 string) int64 {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Encode provides a mock function with given fields: _a0
func (_m *MockCodeService) Encode(_a0 int64) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
