// Code generated by mockery v1.0.0. DO NOT EDIT.

package order

import mock "github.com/stretchr/testify/mock"

// MockUseCase is an autogenerated mock type for the UseCase type
type MockUseCase struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: r
func (_m *MockUseCase) CreateOrder(r *CreateRequest) (*CreateResponse, int, error) {
	ret := _m.Called(r)

	var r0 *CreateResponse
	if rf, ok := ret.Get(0).(func(*CreateRequest) *CreateResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateResponse)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*CreateRequest) int); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*CreateRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrder provides a mock function with given fields: options
func (_m *MockUseCase) GetOrder(options *GetOptions) ([]*GetResponse, int, error) {
	ret := _m.Called(options)

	var r0 []*GetResponse
	if rf, ok := ret.Get(0).(func(*GetOptions) []*GetResponse); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetResponse)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*GetOptions) int); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*GetOptions) error); ok {
		r2 = rf(options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TakeOrder provides a mock function with given fields: id, r
func (_m *MockUseCase) TakeOrder(id int, r *TakeRequest) (*TakeResponse, int, error) {
	ret := _m.Called(id, r)

	var r0 *TakeResponse
	if rf, ok := ret.Get(0).(func(int, *TakeRequest) *TakeResponse); ok {
		r0 = rf(id, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TakeResponse)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, *TakeRequest) int); ok {
		r1 = rf(id, r)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, *TakeRequest) error); ok {
		r2 = rf(id, r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
