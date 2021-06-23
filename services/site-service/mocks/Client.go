// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import client "github.com/veganbase/backend/services/site-service/client"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// SiteUpdates provides a mock function with given fields:
func (_m *Client) SiteUpdates() (chan bool, func(), error) {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func() func()); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Sites provides a mock function with given fields:
func (_m *Client) Sites() client.SiteMap {
	ret := _m.Called()

	var r0 client.SiteMap
	if rf, ok := ret.Get(0).(func() client.SiteMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SiteMap)
		}
	}

	return r0
}