// Code generated by mockery v1.0.0. DO NOT EDIT.

package post

import mock "github.com/stretchr/testify/mock"

// MockPostInterface is an autogenerated mock type for the PostInterface type
type MockPostInterface struct {
	mock.Mock
}

// GetPosts provides a mock function with given fields:
func (_m *MockPostInterface) GetPosts() ([]Post, error) {
	ret := _m.Called()

	var r0 []Post
	if rf, ok := ret.Get(0).(func() []Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
