// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// DaemonSetsGetter is an autogenerated mock type for the DaemonSetsGetter type
type DaemonSetsGetter struct {
	mock.Mock
}

// DaemonSets provides a mock function with given fields: namespace
func (_m *DaemonSetsGetter) DaemonSets(namespace string) v1.DaemonSetInterface {
	ret := _m.Called(namespace)

	var r0 v1.DaemonSetInterface
	if rf, ok := ret.Get(0).(func(string) v1.DaemonSetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.DaemonSetInterface)
		}
	}

	return r0
}
