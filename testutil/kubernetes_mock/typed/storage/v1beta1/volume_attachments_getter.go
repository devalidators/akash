// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

// VolumeAttachmentsGetter is an autogenerated mock type for the VolumeAttachmentsGetter type
type VolumeAttachmentsGetter struct {
	mock.Mock
}

// VolumeAttachments provides a mock function with given fields:
func (_m *VolumeAttachmentsGetter) VolumeAttachments() v1beta1.VolumeAttachmentInterface {
	ret := _m.Called()

	var r0 v1beta1.VolumeAttachmentInterface
	if rf, ok := ret.Get(0).(func() v1beta1.VolumeAttachmentInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.VolumeAttachmentInterface)
		}
	}

	return r0
}
