// Code generated by mockery v1.0.0. DO NOT EDIT.

package generator

import mock "github.com/stretchr/testify/mock"

// mockKeyGenerator is an autogenerated mock type for the keyGenerator type
type mockKeyGenerator struct {
	mock.Mock
}

// AddKey provides a mock function with given fields: key
func (_m *mockKeyGenerator) AddKey(key string) {
	_m.Called(key)
}

// GenerateKey provides a mock function with given fields:
func (_m *mockKeyGenerator) GenerateKey() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
