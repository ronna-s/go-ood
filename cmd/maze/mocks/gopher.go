// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Gopher is an autogenerated mock type for the Gopher type
type Gopher struct {
	mock.Mock
}

// Finished provides a mock function with given fields:
func (_m *Gopher) Finished() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Move provides a mock function with given fields:
func (_m *Gopher) Move() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TurnLeft provides a mock function with given fields:
func (_m *Gopher) TurnLeft() {
	_m.Called()
}

// TurnRight provides a mock function with given fields:
func (_m *Gopher) TurnRight() {
	_m.Called()
}

type mockConstructorTestingTNewGopher interface {
	mock.TestingT
	Cleanup(func())
}

// NewGopher creates a new instance of Gopher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGopher(t mockConstructorTestingTNewGopher) *Gopher {
	mock := &Gopher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
