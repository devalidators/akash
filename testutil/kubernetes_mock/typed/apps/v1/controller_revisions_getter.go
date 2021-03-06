// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// ControllerRevisionsGetter is an autogenerated mock type for the ControllerRevisionsGetter type
type ControllerRevisionsGetter struct {
	mock.Mock
}

// ControllerRevisions provides a mock function with given fields: namespace
func (_m *ControllerRevisionsGetter) ControllerRevisions(namespace string) v1.ControllerRevisionInterface {
	ret := _m.Called(namespace)

	var r0 v1.ControllerRevisionInterface
	if rf, ok := ret.Get(0).(func(string) v1.ControllerRevisionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ControllerRevisionInterface)
		}
	}

	return r0
}
