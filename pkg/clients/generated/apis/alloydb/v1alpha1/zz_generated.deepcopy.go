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
func (in *AlloyDBBackup) DeepCopyInto(out *AlloyDBBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBBackup.
func (in *AlloyDBBackup) DeepCopy() *AlloyDBBackup {
	if in == nil {
		return nil
	}
	out := new(AlloyDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlloyDBBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBBackupList) DeepCopyInto(out *AlloyDBBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlloyDBBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBBackupList.
func (in *AlloyDBBackupList) DeepCopy() *AlloyDBBackupList {
	if in == nil {
		return nil
	}
	out := new(AlloyDBBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlloyDBBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBBackupSpec) DeepCopyInto(out *AlloyDBBackupSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(BackupEncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBBackupSpec.
func (in *AlloyDBBackupSpec) DeepCopy() *AlloyDBBackupSpec {
	if in == nil {
		return nil
	}
	out := new(AlloyDBBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBBackupStatus) DeepCopyInto(out *AlloyDBBackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.EncryptionInfo != nil {
		in, out := &in.EncryptionInfo, &out.EncryptionInfo
		*out = make([]BackupEncryptionInfoStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.Reconciling != nil {
		in, out := &in.Reconciling, &out.Reconciling
		*out = new(bool)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBBackupStatus.
func (in *AlloyDBBackupStatus) DeepCopy() *AlloyDBBackupStatus {
	if in == nil {
		return nil
	}
	out := new(AlloyDBBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBCluster) DeepCopyInto(out *AlloyDBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBCluster.
func (in *AlloyDBCluster) DeepCopy() *AlloyDBCluster {
	if in == nil {
		return nil
	}
	out := new(AlloyDBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlloyDBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBClusterList) DeepCopyInto(out *AlloyDBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlloyDBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBClusterList.
func (in *AlloyDBClusterList) DeepCopy() *AlloyDBClusterList {
	if in == nil {
		return nil
	}
	out := new(AlloyDBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlloyDBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBClusterSpec) DeepCopyInto(out *AlloyDBClusterSpec) {
	*out = *in
	if in.AutomatedBackupPolicy != nil {
		in, out := &in.AutomatedBackupPolicy, &out.AutomatedBackupPolicy
		*out = new(ClusterAutomatedBackupPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ContinuousBackupConfig != nil {
		in, out := &in.ContinuousBackupConfig, &out.ContinuousBackupConfig
		*out = new(ClusterContinuousBackupConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(ClusterEncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.InitialUser != nil {
		in, out := &in.InitialUser, &out.InitialUser
		*out = new(ClusterInitialUser)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RestoreBackupSource != nil {
		in, out := &in.RestoreBackupSource, &out.RestoreBackupSource
		*out = new(ClusterRestoreBackupSource)
		**out = **in
	}
	if in.RestoreContinuousBackupSource != nil {
		in, out := &in.RestoreContinuousBackupSource, &out.RestoreContinuousBackupSource
		*out = new(ClusterRestoreContinuousBackupSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBClusterSpec.
func (in *AlloyDBClusterSpec) DeepCopy() *AlloyDBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AlloyDBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBClusterStatus) DeepCopyInto(out *AlloyDBClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.BackupSource != nil {
		in, out := &in.BackupSource, &out.BackupSource
		*out = make([]ClusterBackupSourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContinuousBackupInfo != nil {
		in, out := &in.ContinuousBackupInfo, &out.ContinuousBackupInfo
		*out = make([]ClusterContinuousBackupInfoStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DatabaseVersion != nil {
		in, out := &in.DatabaseVersion, &out.DatabaseVersion
		*out = new(string)
		**out = **in
	}
	if in.EncryptionInfo != nil {
		in, out := &in.EncryptionInfo, &out.EncryptionInfo
		*out = make([]ClusterEncryptionInfoStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MigrationSource != nil {
		in, out := &in.MigrationSource, &out.MigrationSource
		*out = make([]ClusterMigrationSourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBClusterStatus.
func (in *AlloyDBClusterStatus) DeepCopy() *AlloyDBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AlloyDBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBInstance) DeepCopyInto(out *AlloyDBInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBInstance.
func (in *AlloyDBInstance) DeepCopy() *AlloyDBInstance {
	if in == nil {
		return nil
	}
	out := new(AlloyDBInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlloyDBInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBInstanceList) DeepCopyInto(out *AlloyDBInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlloyDBInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBInstanceList.
func (in *AlloyDBInstanceList) DeepCopy() *AlloyDBInstanceList {
	if in == nil {
		return nil
	}
	out := new(AlloyDBInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlloyDBInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBInstanceSpec) DeepCopyInto(out *AlloyDBInstanceSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AvailabilityType != nil {
		in, out := &in.AvailabilityType, &out.AvailabilityType
		*out = new(string)
		**out = **in
	}
	if in.DatabaseFlags != nil {
		in, out := &in.DatabaseFlags, &out.DatabaseFlags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.GceZone != nil {
		in, out := &in.GceZone, &out.GceZone
		*out = new(string)
		**out = **in
	}
	if in.MachineConfig != nil {
		in, out := &in.MachineConfig, &out.MachineConfig
		*out = new(InstanceMachineConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadPoolConfig != nil {
		in, out := &in.ReadPoolConfig, &out.ReadPoolConfig
		*out = new(InstanceReadPoolConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBInstanceSpec.
func (in *AlloyDBInstanceSpec) DeepCopy() *AlloyDBInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(AlloyDBInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlloyDBInstanceStatus) DeepCopyInto(out *AlloyDBInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.Reconciling != nil {
		in, out := &in.Reconciling, &out.Reconciling
		*out = new(bool)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlloyDBInstanceStatus.
func (in *AlloyDBInstanceStatus) DeepCopy() *AlloyDBInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AlloyDBInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEncryptionConfig) DeepCopyInto(out *BackupEncryptionConfig) {
	*out = *in
	if in.KmsKeyName != nil {
		in, out := &in.KmsKeyName, &out.KmsKeyName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEncryptionConfig.
func (in *BackupEncryptionConfig) DeepCopy() *BackupEncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(BackupEncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEncryptionInfoStatus) DeepCopyInto(out *BackupEncryptionInfoStatus) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyVersions != nil {
		in, out := &in.KmsKeyVersions, &out.KmsKeyVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEncryptionInfoStatus.
func (in *BackupEncryptionInfoStatus) DeepCopy() *BackupEncryptionInfoStatus {
	if in == nil {
		return nil
	}
	out := new(BackupEncryptionInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutomatedBackupPolicy) DeepCopyInto(out *ClusterAutomatedBackupPolicy) {
	*out = *in
	if in.BackupWindow != nil {
		in, out := &in.BackupWindow, &out.BackupWindow
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(ClusterEncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.QuantityBasedRetention != nil {
		in, out := &in.QuantityBasedRetention, &out.QuantityBasedRetention
		*out = new(ClusterQuantityBasedRetention)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeBasedRetention != nil {
		in, out := &in.TimeBasedRetention, &out.TimeBasedRetention
		*out = new(ClusterTimeBasedRetention)
		(*in).DeepCopyInto(*out)
	}
	if in.WeeklySchedule != nil {
		in, out := &in.WeeklySchedule, &out.WeeklySchedule
		*out = new(ClusterWeeklySchedule)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutomatedBackupPolicy.
func (in *ClusterAutomatedBackupPolicy) DeepCopy() *ClusterAutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterAutomatedBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBackupSourceStatus) DeepCopyInto(out *ClusterBackupSourceStatus) {
	*out = *in
	if in.BackupName != nil {
		in, out := &in.BackupName, &out.BackupName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBackupSourceStatus.
func (in *ClusterBackupSourceStatus) DeepCopy() *ClusterBackupSourceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBackupSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContinuousBackupConfig) DeepCopyInto(out *ClusterContinuousBackupConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(ClusterEncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RecoveryWindowDays != nil {
		in, out := &in.RecoveryWindowDays, &out.RecoveryWindowDays
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContinuousBackupConfig.
func (in *ClusterContinuousBackupConfig) DeepCopy() *ClusterContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterContinuousBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContinuousBackupInfoStatus) DeepCopyInto(out *ClusterContinuousBackupInfoStatus) {
	*out = *in
	if in.EarliestRestorableTime != nil {
		in, out := &in.EarliestRestorableTime, &out.EarliestRestorableTime
		*out = new(string)
		**out = **in
	}
	if in.EnabledTime != nil {
		in, out := &in.EnabledTime, &out.EnabledTime
		*out = new(string)
		**out = **in
	}
	if in.EncryptionInfo != nil {
		in, out := &in.EncryptionInfo, &out.EncryptionInfo
		*out = make([]ClusterEncryptionInfoStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContinuousBackupInfoStatus.
func (in *ClusterContinuousBackupInfoStatus) DeepCopy() *ClusterContinuousBackupInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterContinuousBackupInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEncryptionConfig) DeepCopyInto(out *ClusterEncryptionConfig) {
	*out = *in
	if in.KmsKeyName != nil {
		in, out := &in.KmsKeyName, &out.KmsKeyName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEncryptionConfig.
func (in *ClusterEncryptionConfig) DeepCopy() *ClusterEncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterEncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEncryptionInfoStatus) DeepCopyInto(out *ClusterEncryptionInfoStatus) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyVersions != nil {
		in, out := &in.KmsKeyVersions, &out.KmsKeyVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEncryptionInfoStatus.
func (in *ClusterEncryptionInfoStatus) DeepCopy() *ClusterEncryptionInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterEncryptionInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInitialUser) DeepCopyInto(out *ClusterInitialUser) {
	*out = *in
	in.Password.DeepCopyInto(&out.Password)
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInitialUser.
func (in *ClusterInitialUser) DeepCopy() *ClusterInitialUser {
	if in == nil {
		return nil
	}
	out := new(ClusterInitialUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMigrationSourceStatus) DeepCopyInto(out *ClusterMigrationSourceStatus) {
	*out = *in
	if in.HostPort != nil {
		in, out := &in.HostPort, &out.HostPort
		*out = new(string)
		**out = **in
	}
	if in.ReferenceId != nil {
		in, out := &in.ReferenceId, &out.ReferenceId
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMigrationSourceStatus.
func (in *ClusterMigrationSourceStatus) DeepCopy() *ClusterMigrationSourceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterMigrationSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPassword) DeepCopyInto(out *ClusterPassword) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ClusterValueFrom)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPassword.
func (in *ClusterPassword) DeepCopy() *ClusterPassword {
	if in == nil {
		return nil
	}
	out := new(ClusterPassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQuantityBasedRetention) DeepCopyInto(out *ClusterQuantityBasedRetention) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuantityBasedRetention.
func (in *ClusterQuantityBasedRetention) DeepCopy() *ClusterQuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := new(ClusterQuantityBasedRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRestoreBackupSource) DeepCopyInto(out *ClusterRestoreBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRestoreBackupSource.
func (in *ClusterRestoreBackupSource) DeepCopy() *ClusterRestoreBackupSource {
	if in == nil {
		return nil
	}
	out := new(ClusterRestoreBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRestoreContinuousBackupSource) DeepCopyInto(out *ClusterRestoreContinuousBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRestoreContinuousBackupSource.
func (in *ClusterRestoreContinuousBackupSource) DeepCopy() *ClusterRestoreContinuousBackupSource {
	if in == nil {
		return nil
	}
	out := new(ClusterRestoreContinuousBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStartTimes) DeepCopyInto(out *ClusterStartTimes) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(int)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(int)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(int)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStartTimes.
func (in *ClusterStartTimes) DeepCopy() *ClusterStartTimes {
	if in == nil {
		return nil
	}
	out := new(ClusterStartTimes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTimeBasedRetention) DeepCopyInto(out *ClusterTimeBasedRetention) {
	*out = *in
	if in.RetentionPeriod != nil {
		in, out := &in.RetentionPeriod, &out.RetentionPeriod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTimeBasedRetention.
func (in *ClusterTimeBasedRetention) DeepCopy() *ClusterTimeBasedRetention {
	if in == nil {
		return nil
	}
	out := new(ClusterTimeBasedRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterValueFrom) DeepCopyInto(out *ClusterValueFrom) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(k8sv1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterValueFrom.
func (in *ClusterValueFrom) DeepCopy() *ClusterValueFrom {
	if in == nil {
		return nil
	}
	out := new(ClusterValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWeeklySchedule) DeepCopyInto(out *ClusterWeeklySchedule) {
	*out = *in
	if in.DaysOfWeek != nil {
		in, out := &in.DaysOfWeek, &out.DaysOfWeek
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StartTimes != nil {
		in, out := &in.StartTimes, &out.StartTimes
		*out = make([]ClusterStartTimes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWeeklySchedule.
func (in *ClusterWeeklySchedule) DeepCopy() *ClusterWeeklySchedule {
	if in == nil {
		return nil
	}
	out := new(ClusterWeeklySchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMachineConfig) DeepCopyInto(out *InstanceMachineConfig) {
	*out = *in
	if in.CpuCount != nil {
		in, out := &in.CpuCount, &out.CpuCount
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMachineConfig.
func (in *InstanceMachineConfig) DeepCopy() *InstanceMachineConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceMachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceReadPoolConfig) DeepCopyInto(out *InstanceReadPoolConfig) {
	*out = *in
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceReadPoolConfig.
func (in *InstanceReadPoolConfig) DeepCopy() *InstanceReadPoolConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceReadPoolConfig)
	in.DeepCopyInto(out)
	return out
}
