// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package coremocks

import mock "github.com/stretchr/testify/mock"

// SagaData is an autogenerated mock type for the SagaData type
type SagaData struct {
	mock.Mock
}

// SagaDataName provides a mock function with given fields:
func (_m *SagaData) SagaDataName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
