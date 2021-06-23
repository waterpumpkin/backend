// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import db "github.com/veganbase/backend/services/blob-service/db"
import mock "github.com/stretchr/testify/mock"
import model "github.com/veganbase/backend/services/blob-service/model"

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// AddBlobsToItem provides a mock function with given fields: itemID, blobIDs
func (_m *DB) AddBlobsToItem(itemID string, blobIDs []string) error {
	ret := _m.Called(itemID, blobIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(itemID, blobIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlobByID provides a mock function with given fields: id
func (_m *DB) BlobByID(id string) (*model.Blob, error) {
	ret := _m.Called(id)

	var r0 *model.Blob
	if rf, ok := ret.Get(0).(func(string) *model.Blob); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlobsByUser provides a mock function with given fields: userID, tags, page, perPage
func (_m *DB) BlobsByUser(userID string, tags []string, page uint, perPage uint) ([]model.Blob, error) {
	ret := _m.Called(userID, tags, page, perPage)

	var r0 []model.Blob
	if rf, ok := ret.Get(0).(func(string, []string, uint, uint) []model.Blob); ok {
		r0 = rf(userID, tags, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, uint, uint) error); ok {
		r1 = rf(userID, tags, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearBlobOwner provides a mock function with given fields: id
func (_m *DB) ClearBlobOwner(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBlob provides a mock function with given fields: blob
func (_m *DB) CreateBlob(blob *model.Blob) error {
	ret := _m.Called(blob)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Blob) error); ok {
		r0 = rf(blob)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlob provides a mock function with given fields: id
func (_m *DB) DeleteBlob(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlobID provides a mock function with given fields:
func (_m *DB) NewBlobID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RemoveBlobsFromItem provides a mock function with given fields: id, itemIDs
func (_m *DB) RemoveBlobsFromItem(id string, itemIDs []string) ([]db.DeletedBlob, error) {
	ret := _m.Called(id, itemIDs)

	var r0 []db.DeletedBlob
	if rf, ok := ret.Get(0).(func(string, []string) []db.DeletedBlob); ok {
		r0 = rf(id, itemIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.DeletedBlob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(id, itemIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveEvent provides a mock function with given fields: topic, eventData, inTx
func (_m *DB) SaveEvent(topic string, eventData interface{}, inTx func() error) error {
	ret := _m.Called(topic, eventData, inTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, func() error) error); ok {
		r0 = rf(topic, eventData, inTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBlobTags provides a mock function with given fields: id, tags
func (_m *DB) SetBlobTags(id string, tags []string) error {
	ret := _m.Called(id, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(id, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagsForUser provides a mock function with given fields: userID
func (_m *DB) TagsForUser(userID string) ([]string, error) {
	ret := _m.Called(userID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
