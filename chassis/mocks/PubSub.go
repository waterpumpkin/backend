// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	pubsub "github.com/veganbase/backend/chassis/pubsub"
)

// PubSub is an autogenerated mock type for the PubSub type
type PubSub struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *PubSub) Close() {
	_m.Called()
}

// Publish provides a mock function with given fields: topic, data
func (_m *PubSub) Publish(topic string, data interface{}) error {
	ret := _m.Called(topic, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: topic, sub, mode
func (_m *PubSub) Subscribe(topic string, sub string, mode pubsub.SubscriptionMode) (chan []byte, func(), error) {
	ret := _m.Called(topic, sub, mode)

	var r0 chan []byte
	if rf, ok := ret.Get(0).(func(string, string, pubsub.SubscriptionMode) chan []byte); ok {
		r0 = rf(topic, sub, mode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan []byte)
		}
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func(string, string, pubsub.SubscriptionMode) func()); ok {
		r1 = rf(topic, sub, mode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, pubsub.SubscriptionMode) error); ok {
		r2 = rf(topic, sub, mode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}