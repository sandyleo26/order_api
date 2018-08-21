// Code generated by mockery v1.0.0. DO NOT EDIT.

package order

import mock "github.com/stretchr/testify/mock"

// MockDBAdaptor is an autogenerated mock type for the DBAdaptor type
type MockDBAdaptor struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MockDBAdaptor) Create(_a0 *Order) (*Order, error) {
	ret := _m.Called(_a0)

	var r0 *Order
	if rf, ok := ret.Get(0).(func(*Order) *Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Order) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0
func (_m *MockDBAdaptor) Get(_a0 *GetOptions) ([]*Order, error) {
	ret := _m.Called(_a0)

	var r0 []*Order
	if rf, ok := ret.Get(0).(func(*GetOptions) []*Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*GetOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, status
func (_m *MockDBAdaptor) Update(id uint, status Status) (*Order, error) {
	ret := _m.Called(id, status)

	var r0 *Order
	if rf, ok := ret.Get(0).(func(uint, Status) *Order); ok {
		r0 = rf(id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, Status) error); ok {
		r1 = rf(id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
