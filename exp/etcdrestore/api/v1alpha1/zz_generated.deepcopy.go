//go:build !ignore_autogenerated

/*
Copyright © 2023 - 2024 SUSE LLC

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDMachineSnapshot) DeepCopyInto(out *ETCDMachineSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDMachineSnapshot.
func (in *ETCDMachineSnapshot) DeepCopy() *ETCDMachineSnapshot {
	if in == nil {
		return nil
	}
	out := new(ETCDMachineSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ETCDMachineSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDMachineSnapshotList) DeepCopyInto(out *ETCDMachineSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ETCDMachineSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDMachineSnapshotList.
func (in *ETCDMachineSnapshotList) DeepCopy() *ETCDMachineSnapshotList {
	if in == nil {
		return nil
	}
	out := new(ETCDMachineSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ETCDMachineSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDMachineSnapshotSpec) DeepCopyInto(out *ETCDMachineSnapshotSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDMachineSnapshotSpec.
func (in *ETCDMachineSnapshotSpec) DeepCopy() *ETCDMachineSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(ETCDMachineSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDMachineSnapshotStatus) DeepCopyInto(out *ETCDMachineSnapshotStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDMachineSnapshotStatus.
func (in *ETCDMachineSnapshotStatus) DeepCopy() *ETCDMachineSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(ETCDMachineSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotRestore) DeepCopyInto(out *ETCDSnapshotRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotRestore.
func (in *ETCDSnapshotRestore) DeepCopy() *ETCDSnapshotRestore {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ETCDSnapshotRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotRestoreList) DeepCopyInto(out *ETCDSnapshotRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ETCDSnapshotRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotRestoreList.
func (in *ETCDSnapshotRestoreList) DeepCopy() *ETCDSnapshotRestoreList {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ETCDSnapshotRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotRestoreSpec) DeepCopyInto(out *ETCDSnapshotRestoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotRestoreSpec.
func (in *ETCDSnapshotRestoreSpec) DeepCopy() *ETCDSnapshotRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotRestoreStatus) DeepCopyInto(out *ETCDSnapshotRestoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotRestoreStatus.
func (in *ETCDSnapshotRestoreStatus) DeepCopy() *ETCDSnapshotRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalConfig) DeepCopyInto(out *LocalConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalConfig.
func (in *LocalConfig) DeepCopy() *LocalConfig {
	if in == nil {
		return nil
	}
	out := new(LocalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2EtcdMachineSnapshotConfig) DeepCopyInto(out *RKE2EtcdMachineSnapshotConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2EtcdMachineSnapshotConfig.
func (in *RKE2EtcdMachineSnapshotConfig) DeepCopy() *RKE2EtcdMachineSnapshotConfig {
	if in == nil {
		return nil
	}
	out := new(RKE2EtcdMachineSnapshotConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKE2EtcdMachineSnapshotConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2EtcdMachineSnapshotConfigList) DeepCopyInto(out *RKE2EtcdMachineSnapshotConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ETCDSnapshotRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2EtcdMachineSnapshotConfigList.
func (in *RKE2EtcdMachineSnapshotConfigList) DeepCopy() *RKE2EtcdMachineSnapshotConfigList {
	if in == nil {
		return nil
	}
	out := new(RKE2EtcdMachineSnapshotConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKE2EtcdMachineSnapshotConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2EtcdMachineSnapshotConfigSpec) DeepCopyInto(out *RKE2EtcdMachineSnapshotConfigSpec) {
	*out = *in
	out.S3 = in.S3
	out.Local = in.Local
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2EtcdMachineSnapshotConfigSpec.
func (in *RKE2EtcdMachineSnapshotConfigSpec) DeepCopy() *RKE2EtcdMachineSnapshotConfigSpec {
	if in == nil {
		return nil
	}
	out := new(RKE2EtcdMachineSnapshotConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Config) DeepCopyInto(out *S3Config) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Config.
func (in *S3Config) DeepCopy() *S3Config {
	if in == nil {
		return nil
	}
	out := new(S3Config)
	in.DeepCopyInto(out)
	return out
}
