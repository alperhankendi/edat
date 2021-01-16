// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package msgmocks

import (
	context "context"

	msg "github.com/stackus/edat/msg"
	mock "github.com/stretchr/testify/mock"
)

// ReceiveMessageFunc is an autogenerated mock type for the ReceiveMessageFunc type
type ReceiveMessageFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *ReceiveMessageFunc) Execute(_a0 context.Context, _a1 msg.Message) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, msg.Message) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}