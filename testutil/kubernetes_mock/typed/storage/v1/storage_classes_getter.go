// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// StorageClassesGetter is an autogenerated mock type for the StorageClassesGetter type
type StorageClassesGetter struct {
	mock.Mock
}

// StorageClasses provides a mock function with given fields:
func (_m *StorageClassesGetter) StorageClasses() v1.StorageClassInterface {
	ret := _m.Called()

	var r0 v1.StorageClassInterface
	if rf, ok := ret.Get(0).(func() v1.StorageClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.StorageClassInterface)
		}
	}

	return r0
}
