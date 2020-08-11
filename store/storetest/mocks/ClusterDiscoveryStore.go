// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// ClusterDiscoveryStore is an autogenerated mock type for the ClusterDiscoveryStore type
type ClusterDiscoveryStore struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields:
func (_m *ClusterDiscoveryStore) Cleanup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: discovery
func (_m *ClusterDiscoveryStore) Delete(discovery *model.ClusterDiscovery) (bool, error) {
	ret := _m.Called(discovery)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*model.ClusterDiscovery) bool); ok {
		r0 = rf(discovery)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ClusterDiscovery) error); ok {
		r1 = rf(discovery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: discovery
func (_m *ClusterDiscoveryStore) Exists(discovery *model.ClusterDiscovery) (bool, error) {
	ret := _m.Called(discovery)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*model.ClusterDiscovery) bool); ok {
		r0 = rf(discovery)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ClusterDiscovery) error); ok {
		r1 = rf(discovery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: discoveryType, clusterName
func (_m *ClusterDiscoveryStore) GetAll(discoveryType string, clusterName string) ([]*model.ClusterDiscovery, error) {
	ret := _m.Called(discoveryType, clusterName)

	var r0 []*model.ClusterDiscovery
	if rf, ok := ret.Get(0).(func(string, string) []*model.ClusterDiscovery); ok {
		r0 = rf(discoveryType, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ClusterDiscovery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(discoveryType, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: discovery
func (_m *ClusterDiscoveryStore) Save(discovery *model.ClusterDiscovery) error {
	ret := _m.Called(discovery)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ClusterDiscovery) error); ok {
		r0 = rf(discovery)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLastPingAt provides a mock function with given fields: discovery
func (_m *ClusterDiscoveryStore) SetLastPingAt(discovery *model.ClusterDiscovery) error {
	ret := _m.Called(discovery)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ClusterDiscovery) error); ok {
		r0 = rf(discovery)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
