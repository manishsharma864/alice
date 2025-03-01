// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	types "github.com/getamis/alice/internal/message/types"
)

// MessageMain is an autogenerated mock type for the MessageMain type
type MessageMain struct {
	mock.Mock
}

// AddMessage provides a mock function with given fields: msg
func (_m *MessageMain) AddMessage(msg types.Message) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Message) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMessageMain creates a new instance of MessageMain. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageMain(t testing.TB) *MessageMain {
	mock := &MessageMain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
