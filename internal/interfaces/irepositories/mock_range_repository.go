// Code generated by mockery 2.9.4. DO NOT EDIT.

package irepositories

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	repositories "github.com/theoptz/url-shortener/internal/repositories"
)

// MockRangeRepository is an autogenerated mock type for the RangeRepository type
type MockRangeRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *MockRangeRepository) Get(_a0 context.Context) (*repositories.RangeItem, error) {
	ret := _m.Called(_a0)

	var r0 *repositories.RangeItem
	if rf, ok := ret.Get(0).(func(context.Context) *repositories.RangeItem); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repositories.RangeItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNext provides a mock function with given fields: _a0
func (_m *MockRangeRepository) GetNext(_a0 context.Context) (*repositories.RangeItem, error) {
	ret := _m.Called(_a0)

	var r0 *repositories.RangeItem
	if rf, ok := ret.Get(0).(func(context.Context) *repositories.RangeItem); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repositories.RangeItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
