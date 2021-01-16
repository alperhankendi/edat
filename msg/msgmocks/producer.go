// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package msgmocks

import (
	context "context"

	msg "github.com/stackus/edat/msg"
	mock "github.com/stretchr/testify/mock"
)

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *Producer) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: ctx, channel, message
func (_m *Producer) Send(ctx context.Context, channel string, message msg.Message) error {
	ret := _m.Called(ctx, channel, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, msg.Message) error); ok {
		r0 = rf(ctx, channel, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}