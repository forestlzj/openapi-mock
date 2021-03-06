// Code generated by mockery v1.0.0. DO NOT EDIT.

package serializer

import mock "github.com/stretchr/testify/mock"

// MockSerializer is an autogenerated mock type for the Serializer type
type MockSerializer struct {
	mock.Mock
}

// Serialize provides a mock function with given fields: data, format
func (_m *MockSerializer) Serialize(data interface{}, format string) ([]byte, error) {
	ret := _m.Called(data, format)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}, string) []byte); ok {
		r0 = rf(data, format)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string) error); ok {
		r1 = rf(data, format)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
