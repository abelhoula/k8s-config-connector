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
func (in *DicomstoreBigqueryDestination) DeepCopyInto(out *DicomstoreBigqueryDestination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DicomstoreBigqueryDestination.
func (in *DicomstoreBigqueryDestination) DeepCopy() *DicomstoreBigqueryDestination {
	if in == nil {
		return nil
	}
	out := new(DicomstoreBigqueryDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DicomstoreNotificationConfig) DeepCopyInto(out *DicomstoreNotificationConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DicomstoreNotificationConfig.
func (in *DicomstoreNotificationConfig) DeepCopy() *DicomstoreNotificationConfig {
	if in == nil {
		return nil
	}
	out := new(DicomstoreNotificationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DicomstoreStreamConfigs) DeepCopyInto(out *DicomstoreStreamConfigs) {
	*out = *in
	out.BigqueryDestination = in.BigqueryDestination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DicomstoreStreamConfigs.
func (in *DicomstoreStreamConfigs) DeepCopy() *DicomstoreStreamConfigs {
	if in == nil {
		return nil
	}
	out := new(DicomstoreStreamConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FhirstoreBigqueryDestination) DeepCopyInto(out *FhirstoreBigqueryDestination) {
	*out = *in
	in.SchemaConfig.DeepCopyInto(&out.SchemaConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FhirstoreBigqueryDestination.
func (in *FhirstoreBigqueryDestination) DeepCopy() *FhirstoreBigqueryDestination {
	if in == nil {
		return nil
	}
	out := new(FhirstoreBigqueryDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FhirstoreLastUpdatedPartitionConfig) DeepCopyInto(out *FhirstoreLastUpdatedPartitionConfig) {
	*out = *in
	if in.ExpirationMs != nil {
		in, out := &in.ExpirationMs, &out.ExpirationMs
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FhirstoreLastUpdatedPartitionConfig.
func (in *FhirstoreLastUpdatedPartitionConfig) DeepCopy() *FhirstoreLastUpdatedPartitionConfig {
	if in == nil {
		return nil
	}
	out := new(FhirstoreLastUpdatedPartitionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FhirstoreNotificationConfig) DeepCopyInto(out *FhirstoreNotificationConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FhirstoreNotificationConfig.
func (in *FhirstoreNotificationConfig) DeepCopy() *FhirstoreNotificationConfig {
	if in == nil {
		return nil
	}
	out := new(FhirstoreNotificationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FhirstoreNotificationConfigs) DeepCopyInto(out *FhirstoreNotificationConfigs) {
	*out = *in
	if in.SendFullResource != nil {
		in, out := &in.SendFullResource, &out.SendFullResource
		*out = new(bool)
		**out = **in
	}
	if in.SendPreviousResourceOnDelete != nil {
		in, out := &in.SendPreviousResourceOnDelete, &out.SendPreviousResourceOnDelete
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FhirstoreNotificationConfigs.
func (in *FhirstoreNotificationConfigs) DeepCopy() *FhirstoreNotificationConfigs {
	if in == nil {
		return nil
	}
	out := new(FhirstoreNotificationConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FhirstoreSchemaConfig) DeepCopyInto(out *FhirstoreSchemaConfig) {
	*out = *in
	if in.LastUpdatedPartitionConfig != nil {
		in, out := &in.LastUpdatedPartitionConfig, &out.LastUpdatedPartitionConfig
		*out = new(FhirstoreLastUpdatedPartitionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SchemaType != nil {
		in, out := &in.SchemaType, &out.SchemaType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FhirstoreSchemaConfig.
func (in *FhirstoreSchemaConfig) DeepCopy() *FhirstoreSchemaConfig {
	if in == nil {
		return nil
	}
	out := new(FhirstoreSchemaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FhirstoreStreamConfigs) DeepCopyInto(out *FhirstoreStreamConfigs) {
	*out = *in
	in.BigqueryDestination.DeepCopyInto(&out.BigqueryDestination)
	if in.ResourceTypes != nil {
		in, out := &in.ResourceTypes, &out.ResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FhirstoreStreamConfigs.
func (in *FhirstoreStreamConfigs) DeepCopy() *FhirstoreStreamConfigs {
	if in == nil {
		return nil
	}
	out := new(FhirstoreStreamConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareConsentStore) DeepCopyInto(out *HealthcareConsentStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareConsentStore.
func (in *HealthcareConsentStore) DeepCopy() *HealthcareConsentStore {
	if in == nil {
		return nil
	}
	out := new(HealthcareConsentStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareConsentStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareConsentStoreList) DeepCopyInto(out *HealthcareConsentStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthcareConsentStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareConsentStoreList.
func (in *HealthcareConsentStoreList) DeepCopy() *HealthcareConsentStoreList {
	if in == nil {
		return nil
	}
	out := new(HealthcareConsentStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareConsentStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareConsentStoreSpec) DeepCopyInto(out *HealthcareConsentStoreSpec) {
	*out = *in
	if in.DefaultConsentTtl != nil {
		in, out := &in.DefaultConsentTtl, &out.DefaultConsentTtl
		*out = new(string)
		**out = **in
	}
	if in.EnableConsentCreateOnUpdate != nil {
		in, out := &in.EnableConsentCreateOnUpdate, &out.EnableConsentCreateOnUpdate
		*out = new(bool)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareConsentStoreSpec.
func (in *HealthcareConsentStoreSpec) DeepCopy() *HealthcareConsentStoreSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcareConsentStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareConsentStoreStatus) DeepCopyInto(out *HealthcareConsentStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareConsentStoreStatus.
func (in *HealthcareConsentStoreStatus) DeepCopy() *HealthcareConsentStoreStatus {
	if in == nil {
		return nil
	}
	out := new(HealthcareConsentStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDICOMStore) DeepCopyInto(out *HealthcareDICOMStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDICOMStore.
func (in *HealthcareDICOMStore) DeepCopy() *HealthcareDICOMStore {
	if in == nil {
		return nil
	}
	out := new(HealthcareDICOMStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareDICOMStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDICOMStoreList) DeepCopyInto(out *HealthcareDICOMStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthcareDICOMStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDICOMStoreList.
func (in *HealthcareDICOMStoreList) DeepCopy() *HealthcareDICOMStoreList {
	if in == nil {
		return nil
	}
	out := new(HealthcareDICOMStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareDICOMStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDICOMStoreSpec) DeepCopyInto(out *HealthcareDICOMStoreSpec) {
	*out = *in
	if in.NotificationConfig != nil {
		in, out := &in.NotificationConfig, &out.NotificationConfig
		*out = new(DicomstoreNotificationConfig)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.StreamConfigs != nil {
		in, out := &in.StreamConfigs, &out.StreamConfigs
		*out = make([]DicomstoreStreamConfigs, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDICOMStoreSpec.
func (in *HealthcareDICOMStoreSpec) DeepCopy() *HealthcareDICOMStoreSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcareDICOMStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDICOMStoreStatus) DeepCopyInto(out *HealthcareDICOMStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDICOMStoreStatus.
func (in *HealthcareDICOMStoreStatus) DeepCopy() *HealthcareDICOMStoreStatus {
	if in == nil {
		return nil
	}
	out := new(HealthcareDICOMStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDataset) DeepCopyInto(out *HealthcareDataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDataset.
func (in *HealthcareDataset) DeepCopy() *HealthcareDataset {
	if in == nil {
		return nil
	}
	out := new(HealthcareDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareDataset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDatasetList) DeepCopyInto(out *HealthcareDatasetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthcareDataset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDatasetList.
func (in *HealthcareDatasetList) DeepCopy() *HealthcareDatasetList {
	if in == nil {
		return nil
	}
	out := new(HealthcareDatasetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareDatasetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDatasetSpec) DeepCopyInto(out *HealthcareDatasetSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDatasetSpec.
func (in *HealthcareDatasetSpec) DeepCopy() *HealthcareDatasetSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcareDatasetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareDatasetStatus) DeepCopyInto(out *HealthcareDatasetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareDatasetStatus.
func (in *HealthcareDatasetStatus) DeepCopy() *HealthcareDatasetStatus {
	if in == nil {
		return nil
	}
	out := new(HealthcareDatasetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareFHIRStore) DeepCopyInto(out *HealthcareFHIRStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareFHIRStore.
func (in *HealthcareFHIRStore) DeepCopy() *HealthcareFHIRStore {
	if in == nil {
		return nil
	}
	out := new(HealthcareFHIRStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareFHIRStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareFHIRStoreList) DeepCopyInto(out *HealthcareFHIRStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthcareFHIRStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareFHIRStoreList.
func (in *HealthcareFHIRStoreList) DeepCopy() *HealthcareFHIRStoreList {
	if in == nil {
		return nil
	}
	out := new(HealthcareFHIRStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareFHIRStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareFHIRStoreSpec) DeepCopyInto(out *HealthcareFHIRStoreSpec) {
	*out = *in
	if in.ComplexDataTypeReferenceParsing != nil {
		in, out := &in.ComplexDataTypeReferenceParsing, &out.ComplexDataTypeReferenceParsing
		*out = new(string)
		**out = **in
	}
	if in.DefaultSearchHandlingStrict != nil {
		in, out := &in.DefaultSearchHandlingStrict, &out.DefaultSearchHandlingStrict
		*out = new(bool)
		**out = **in
	}
	if in.DisableReferentialIntegrity != nil {
		in, out := &in.DisableReferentialIntegrity, &out.DisableReferentialIntegrity
		*out = new(bool)
		**out = **in
	}
	if in.DisableResourceVersioning != nil {
		in, out := &in.DisableResourceVersioning, &out.DisableResourceVersioning
		*out = new(bool)
		**out = **in
	}
	if in.EnableHistoryImport != nil {
		in, out := &in.EnableHistoryImport, &out.EnableHistoryImport
		*out = new(bool)
		**out = **in
	}
	if in.EnableUpdateCreate != nil {
		in, out := &in.EnableUpdateCreate, &out.EnableUpdateCreate
		*out = new(bool)
		**out = **in
	}
	if in.NotificationConfig != nil {
		in, out := &in.NotificationConfig, &out.NotificationConfig
		*out = new(FhirstoreNotificationConfig)
		**out = **in
	}
	if in.NotificationConfigs != nil {
		in, out := &in.NotificationConfigs, &out.NotificationConfigs
		*out = make([]FhirstoreNotificationConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.StreamConfigs != nil {
		in, out := &in.StreamConfigs, &out.StreamConfigs
		*out = make([]FhirstoreStreamConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareFHIRStoreSpec.
func (in *HealthcareFHIRStoreSpec) DeepCopy() *HealthcareFHIRStoreSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcareFHIRStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareFHIRStoreStatus) DeepCopyInto(out *HealthcareFHIRStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareFHIRStoreStatus.
func (in *HealthcareFHIRStoreStatus) DeepCopy() *HealthcareFHIRStoreStatus {
	if in == nil {
		return nil
	}
	out := new(HealthcareFHIRStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareHL7V2Store) DeepCopyInto(out *HealthcareHL7V2Store) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareHL7V2Store.
func (in *HealthcareHL7V2Store) DeepCopy() *HealthcareHL7V2Store {
	if in == nil {
		return nil
	}
	out := new(HealthcareHL7V2Store)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareHL7V2Store) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareHL7V2StoreList) DeepCopyInto(out *HealthcareHL7V2StoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthcareHL7V2Store, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareHL7V2StoreList.
func (in *HealthcareHL7V2StoreList) DeepCopy() *HealthcareHL7V2StoreList {
	if in == nil {
		return nil
	}
	out := new(HealthcareHL7V2StoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthcareHL7V2StoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareHL7V2StoreSpec) DeepCopyInto(out *HealthcareHL7V2StoreSpec) {
	*out = *in
	if in.NotificationConfig != nil {
		in, out := &in.NotificationConfig, &out.NotificationConfig
		*out = new(Hl7v2storeNotificationConfig)
		**out = **in
	}
	if in.NotificationConfigs != nil {
		in, out := &in.NotificationConfigs, &out.NotificationConfigs
		*out = make([]Hl7v2storeNotificationConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ParserConfig != nil {
		in, out := &in.ParserConfig, &out.ParserConfig
		*out = new(Hl7v2storeParserConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareHL7V2StoreSpec.
func (in *HealthcareHL7V2StoreSpec) DeepCopy() *HealthcareHL7V2StoreSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcareHL7V2StoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcareHL7V2StoreStatus) DeepCopyInto(out *HealthcareHL7V2StoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcareHL7V2StoreStatus.
func (in *HealthcareHL7V2StoreStatus) DeepCopy() *HealthcareHL7V2StoreStatus {
	if in == nil {
		return nil
	}
	out := new(HealthcareHL7V2StoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hl7v2storeNotificationConfig) DeepCopyInto(out *Hl7v2storeNotificationConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hl7v2storeNotificationConfig.
func (in *Hl7v2storeNotificationConfig) DeepCopy() *Hl7v2storeNotificationConfig {
	if in == nil {
		return nil
	}
	out := new(Hl7v2storeNotificationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hl7v2storeNotificationConfigs) DeepCopyInto(out *Hl7v2storeNotificationConfigs) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hl7v2storeNotificationConfigs.
func (in *Hl7v2storeNotificationConfigs) DeepCopy() *Hl7v2storeNotificationConfigs {
	if in == nil {
		return nil
	}
	out := new(Hl7v2storeNotificationConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hl7v2storeParserConfig) DeepCopyInto(out *Hl7v2storeParserConfig) {
	*out = *in
	if in.AllowNullHeader != nil {
		in, out := &in.AllowNullHeader, &out.AllowNullHeader
		*out = new(bool)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.SegmentTerminator != nil {
		in, out := &in.SegmentTerminator, &out.SegmentTerminator
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hl7v2storeParserConfig.
func (in *Hl7v2storeParserConfig) DeepCopy() *Hl7v2storeParserConfig {
	if in == nil {
		return nil
	}
	out := new(Hl7v2storeParserConfig)
	in.DeepCopyInto(out)
	return out
}
