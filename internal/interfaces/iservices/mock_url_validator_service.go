// Code generated by mockery 2.9.4. DO NOT EDIT.

package iservices

import mock "github.com/stretchr/testify/mock"

// MockURLValidatorService is an autogenerated mock type for the URLValidatorService type
type MockURLValidatorService struct {
	mock.Mock
}

// Validate provides a mock function with given fields: _a0
func (_m *MockURLValidatorService) Validate(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
