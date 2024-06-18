//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPool) DeepCopyInto(out *CloudBuildWorkerPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPool.
func (in *CloudBuildWorkerPool) DeepCopy() *CloudBuildWorkerPool {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudBuildWorkerPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPoolList) DeepCopyInto(out *CloudBuildWorkerPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudBuildWorkerPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPoolList.
func (in *CloudBuildWorkerPoolList) DeepCopy() *CloudBuildWorkerPoolList {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudBuildWorkerPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPoolSpec) DeepCopyInto(out *CloudBuildWorkerPoolSpec) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	in.PrivatePoolV1Config.DeepCopyInto(&out.PrivatePoolV1Config)
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPoolSpec.
func (in *CloudBuildWorkerPoolSpec) DeepCopy() *CloudBuildWorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPoolStatus) DeepCopyInto(out *CloudBuildWorkerPoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(WorkerpoolObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPoolStatus.
func (in *CloudBuildWorkerPoolStatus) DeepCopy() *CloudBuildWorkerPoolStatus {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolNetworkConfig) DeepCopyInto(out *WorkerpoolNetworkConfig) {
	*out = *in
	if in.PeeredNetworkIPRange != nil {
		in, out := &in.PeeredNetworkIPRange, &out.PeeredNetworkIPRange
		*out = new(string)
		**out = **in
	}
	if in.EgressOption != nil {
		in, out := &in.EgressOption, &out.EgressOption
		*out = new(string)
		**out = **in
	}
	out.PeeredNetworkRef = in.PeeredNetworkRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolNetworkConfig.
func (in *WorkerpoolNetworkConfig) DeepCopy() *WorkerpoolNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolNetworkConfigStatus) DeepCopyInto(out *WorkerpoolNetworkConfigStatus) {
	*out = *in
	if in.EgressOption != nil {
		in, out := &in.EgressOption, &out.EgressOption
		*out = new(string)
		**out = **in
	}
	if in.PeeredNetwork != nil {
		in, out := &in.PeeredNetwork, &out.PeeredNetwork
		*out = new(string)
		**out = **in
	}
	if in.PeeredNetworkIPRange != nil {
		in, out := &in.PeeredNetworkIPRange, &out.PeeredNetworkIPRange
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolNetworkConfigStatus.
func (in *WorkerpoolNetworkConfigStatus) DeepCopy() *WorkerpoolNetworkConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolNetworkConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolObservedStateStatus) DeepCopyInto(out *WorkerpoolObservedStateStatus) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(WorkerpoolNetworkConfigStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WorkerConfig != nil {
		in, out := &in.WorkerConfig, &out.WorkerConfig
		*out = new(WorkerpoolWorkerConfigStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolObservedStateStatus.
func (in *WorkerpoolObservedStateStatus) DeepCopy() *WorkerpoolObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolPrivatePoolV1Config) DeepCopyInto(out *WorkerpoolPrivatePoolV1Config) {
	*out = *in
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(WorkerpoolNetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	in.WorkerConfig.DeepCopyInto(&out.WorkerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolPrivatePoolV1Config.
func (in *WorkerpoolPrivatePoolV1Config) DeepCopy() *WorkerpoolPrivatePoolV1Config {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolPrivatePoolV1Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolWorkerConfig) DeepCopyInto(out *WorkerpoolWorkerConfig) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolWorkerConfig.
func (in *WorkerpoolWorkerConfig) DeepCopy() *WorkerpoolWorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolWorkerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolWorkerConfigStatus) DeepCopyInto(out *WorkerpoolWorkerConfigStatus) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolWorkerConfigStatus.
func (in *WorkerpoolWorkerConfigStatus) DeepCopy() *WorkerpoolWorkerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolWorkerConfigStatus)
	in.DeepCopyInto(out)
	return out
}
