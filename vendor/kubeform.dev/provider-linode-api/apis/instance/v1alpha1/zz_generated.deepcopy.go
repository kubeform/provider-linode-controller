//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(InstanceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecAlerts) DeepCopyInto(out *InstanceSpecAlerts) {
	*out = *in
	if in.Cpu != nil {
		in, out := &in.Cpu, &out.Cpu
		*out = new(int64)
		**out = **in
	}
	if in.Io != nil {
		in, out := &in.Io, &out.Io
		*out = new(int64)
		**out = **in
	}
	if in.NetworkIn != nil {
		in, out := &in.NetworkIn, &out.NetworkIn
		*out = new(int64)
		**out = **in
	}
	if in.NetworkOut != nil {
		in, out := &in.NetworkOut, &out.NetworkOut
		*out = new(int64)
		**out = **in
	}
	if in.TransferQuota != nil {
		in, out := &in.TransferQuota, &out.TransferQuota
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecAlerts.
func (in *InstanceSpecAlerts) DeepCopy() *InstanceSpecAlerts {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecAlerts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecBackups) DeepCopyInto(out *InstanceSpecBackups) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]InstanceSpecBackupsSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecBackups.
func (in *InstanceSpecBackups) DeepCopy() *InstanceSpecBackups {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecBackups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecBackupsSchedule) DeepCopyInto(out *InstanceSpecBackupsSchedule) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecBackupsSchedule.
func (in *InstanceSpecBackupsSchedule) DeepCopy() *InstanceSpecBackupsSchedule {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecBackupsSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfig) DeepCopyInto(out *InstanceSpecConfig) {
	*out = *in
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = new(string)
		**out = **in
	}
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = new(InstanceSpecConfigDevices)
		(*in).DeepCopyInto(*out)
	}
	if in.Helpers != nil {
		in, out := &in.Helpers, &out.Helpers
		*out = new(InstanceSpecConfigHelpers)
		(*in).DeepCopyInto(*out)
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = make([]InstanceSpecConfigInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kernel != nil {
		in, out := &in.Kernel, &out.Kernel
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(int64)
		**out = **in
	}
	if in.RootDevice != nil {
		in, out := &in.RootDevice, &out.RootDevice
		*out = new(string)
		**out = **in
	}
	if in.RunLevel != nil {
		in, out := &in.RunLevel, &out.RunLevel
		*out = new(string)
		**out = **in
	}
	if in.VirtMode != nil {
		in, out := &in.VirtMode, &out.VirtMode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfig.
func (in *InstanceSpecConfig) DeepCopy() *InstanceSpecConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevices) DeepCopyInto(out *InstanceSpecConfigDevices) {
	*out = *in
	if in.Sda != nil {
		in, out := &in.Sda, &out.Sda
		*out = new(InstanceSpecConfigDevicesSda)
		(*in).DeepCopyInto(*out)
	}
	if in.Sdb != nil {
		in, out := &in.Sdb, &out.Sdb
		*out = new(InstanceSpecConfigDevicesSdb)
		(*in).DeepCopyInto(*out)
	}
	if in.Sdc != nil {
		in, out := &in.Sdc, &out.Sdc
		*out = new(InstanceSpecConfigDevicesSdc)
		(*in).DeepCopyInto(*out)
	}
	if in.Sdd != nil {
		in, out := &in.Sdd, &out.Sdd
		*out = new(InstanceSpecConfigDevicesSdd)
		(*in).DeepCopyInto(*out)
	}
	if in.Sde != nil {
		in, out := &in.Sde, &out.Sde
		*out = new(InstanceSpecConfigDevicesSde)
		(*in).DeepCopyInto(*out)
	}
	if in.Sdf != nil {
		in, out := &in.Sdf, &out.Sdf
		*out = new(InstanceSpecConfigDevicesSdf)
		(*in).DeepCopyInto(*out)
	}
	if in.Sdg != nil {
		in, out := &in.Sdg, &out.Sdg
		*out = new(InstanceSpecConfigDevicesSdg)
		(*in).DeepCopyInto(*out)
	}
	if in.Sdh != nil {
		in, out := &in.Sdh, &out.Sdh
		*out = new(InstanceSpecConfigDevicesSdh)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevices.
func (in *InstanceSpecConfigDevices) DeepCopy() *InstanceSpecConfigDevices {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSda) DeepCopyInto(out *InstanceSpecConfigDevicesSda) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSda.
func (in *InstanceSpecConfigDevicesSda) DeepCopy() *InstanceSpecConfigDevicesSda {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSda)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSdb) DeepCopyInto(out *InstanceSpecConfigDevicesSdb) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSdb.
func (in *InstanceSpecConfigDevicesSdb) DeepCopy() *InstanceSpecConfigDevicesSdb {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSdb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSdc) DeepCopyInto(out *InstanceSpecConfigDevicesSdc) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSdc.
func (in *InstanceSpecConfigDevicesSdc) DeepCopy() *InstanceSpecConfigDevicesSdc {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSdc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSdd) DeepCopyInto(out *InstanceSpecConfigDevicesSdd) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSdd.
func (in *InstanceSpecConfigDevicesSdd) DeepCopy() *InstanceSpecConfigDevicesSdd {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSdd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSde) DeepCopyInto(out *InstanceSpecConfigDevicesSde) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSde.
func (in *InstanceSpecConfigDevicesSde) DeepCopy() *InstanceSpecConfigDevicesSde {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSde)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSdf) DeepCopyInto(out *InstanceSpecConfigDevicesSdf) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSdf.
func (in *InstanceSpecConfigDevicesSdf) DeepCopy() *InstanceSpecConfigDevicesSdf {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSdf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSdg) DeepCopyInto(out *InstanceSpecConfigDevicesSdg) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSdg.
func (in *InstanceSpecConfigDevicesSdg) DeepCopy() *InstanceSpecConfigDevicesSdg {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSdg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigDevicesSdh) DeepCopyInto(out *InstanceSpecConfigDevicesSdh) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(int64)
		**out = **in
	}
	if in.DiskLabel != nil {
		in, out := &in.DiskLabel, &out.DiskLabel
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigDevicesSdh.
func (in *InstanceSpecConfigDevicesSdh) DeepCopy() *InstanceSpecConfigDevicesSdh {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigDevicesSdh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigHelpers) DeepCopyInto(out *InstanceSpecConfigHelpers) {
	*out = *in
	if in.DevtmpfsAutomount != nil {
		in, out := &in.DevtmpfsAutomount, &out.DevtmpfsAutomount
		*out = new(bool)
		**out = **in
	}
	if in.Distro != nil {
		in, out := &in.Distro, &out.Distro
		*out = new(bool)
		**out = **in
	}
	if in.ModulesDep != nil {
		in, out := &in.ModulesDep, &out.ModulesDep
		*out = new(bool)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(bool)
		**out = **in
	}
	if in.UpdatedbDisabled != nil {
		in, out := &in.UpdatedbDisabled, &out.UpdatedbDisabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigHelpers.
func (in *InstanceSpecConfigHelpers) DeepCopy() *InstanceSpecConfigHelpers {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigHelpers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecConfigInterface) DeepCopyInto(out *InstanceSpecConfigInterface) {
	*out = *in
	if in.IpamAddress != nil {
		in, out := &in.IpamAddress, &out.IpamAddress
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecConfigInterface.
func (in *InstanceSpecConfigInterface) DeepCopy() *InstanceSpecConfigInterface {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecConfigInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecDisk) DeepCopyInto(out *InstanceSpecDisk) {
	*out = *in
	if in.AuthorizedKeys != nil {
		in, out := &in.AuthorizedKeys, &out.AuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AuthorizedUsers != nil {
		in, out := &in.AuthorizedUsers, &out.AuthorizedUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Filesystem != nil {
		in, out := &in.Filesystem, &out.Filesystem
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	if in.RootPass != nil {
		in, out := &in.RootPass, &out.RootPass
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	if in.StackscriptData != nil {
		in, out := &in.StackscriptData, &out.StackscriptData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StackscriptID != nil {
		in, out := &in.StackscriptID, &out.StackscriptID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecDisk.
func (in *InstanceSpecDisk) DeepCopy() *InstanceSpecDisk {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecInterface) DeepCopyInto(out *InstanceSpecInterface) {
	*out = *in
	if in.IpamAddress != nil {
		in, out := &in.IpamAddress, &out.IpamAddress
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecInterface.
func (in *InstanceSpecInterface) DeepCopy() *InstanceSpecInterface {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecResource) DeepCopyInto(out *InstanceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Alerts != nil {
		in, out := &in.Alerts, &out.Alerts
		*out = new(InstanceSpecAlerts)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthorizedKeys != nil {
		in, out := &in.AuthorizedKeys, &out.AuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AuthorizedUsers != nil {
		in, out := &in.AuthorizedUsers, &out.AuthorizedUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(int64)
		**out = **in
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]InstanceSpecBackups, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackupsEnabled != nil {
		in, out := &in.BackupsEnabled, &out.BackupsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BootConfigLabel != nil {
		in, out := &in.BootConfigLabel, &out.BootConfigLabel
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]InstanceSpecConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]InstanceSpecDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = make([]InstanceSpecInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Ipv4 != nil {
		in, out := &in.Ipv4, &out.Ipv4
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ipv6 != nil {
		in, out := &in.Ipv6, &out.Ipv6
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(bool)
		**out = **in
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RootPass != nil {
		in, out := &in.RootPass, &out.RootPass
		*out = new(string)
		**out = **in
	}
	if in.Specs != nil {
		in, out := &in.Specs, &out.Specs
		*out = make([]InstanceSpecSpecs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StackscriptData != nil {
		in, out := &in.StackscriptData, &out.StackscriptData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StackscriptID != nil {
		in, out := &in.StackscriptID, &out.StackscriptID
		*out = new(int64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SwapSize != nil {
		in, out := &in.SwapSize, &out.SwapSize
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.WatchdogEnabled != nil {
		in, out := &in.WatchdogEnabled, &out.WatchdogEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecResource.
func (in *InstanceSpecResource) DeepCopy() *InstanceSpecResource {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpecSpecs) DeepCopyInto(out *InstanceSpecSpecs) {
	*out = *in
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = new(int64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(int64)
		**out = **in
	}
	if in.Transfer != nil {
		in, out := &in.Transfer, &out.Transfer
		*out = new(int64)
		**out = **in
	}
	if in.Vcpus != nil {
		in, out := &in.Vcpus, &out.Vcpus
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpecSpecs.
func (in *InstanceSpecSpecs) DeepCopy() *InstanceSpecSpecs {
	if in == nil {
		return nil
	}
	out := new(InstanceSpecSpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ip) DeepCopyInto(out *Ip) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ip.
func (in *Ip) DeepCopy() *Ip {
	if in == nil {
		return nil
	}
	out := new(Ip)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ip) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpList) DeepCopyInto(out *IpList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ip, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpList.
func (in *IpList) DeepCopy() *IpList {
	if in == nil {
		return nil
	}
	out := new(IpList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpSpec) DeepCopyInto(out *IpSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IpSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpSpec.
func (in *IpSpec) DeepCopy() *IpSpec {
	if in == nil {
		return nil
	}
	out := new(IpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpSpecResource) DeepCopyInto(out *IpSpecResource) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(int64)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(int64)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.Rdns != nil {
		in, out := &in.Rdns, &out.Rdns
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubnetMask != nil {
		in, out := &in.SubnetMask, &out.SubnetMask
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpSpecResource.
func (in *IpSpecResource) DeepCopy() *IpSpecResource {
	if in == nil {
		return nil
	}
	out := new(IpSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpStatus) DeepCopyInto(out *IpStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpStatus.
func (in *IpStatus) DeepCopy() *IpStatus {
	if in == nil {
		return nil
	}
	out := new(IpStatus)
	in.DeepCopyInto(out)
	return out
}
