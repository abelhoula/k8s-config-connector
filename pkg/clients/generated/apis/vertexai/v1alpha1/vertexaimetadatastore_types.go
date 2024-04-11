// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MetadatastoreEncryptionSpec struct {
	/* Immutable. Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created. */
	// +optional
	KmsKeyName *string `json:"kmsKeyName,omitempty"`
}

type VertexAIMetadataStoreSpec struct {
	/* Immutable. Description of the MetadataStore. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Customer-managed encryption key spec for a MetadataStore. If set, this MetadataStore and all sub-resources of this MetadataStore will be secured by this key. */
	// +optional
	EncryptionSpec *MetadatastoreEncryptionSpec `json:"encryptionSpec,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The region of the Metadata Store. eg us-central1. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type MetadatastoreStateStatus struct {
	/* The disk utilization of the MetadataStore in bytes. */
	// +optional
	DiskUtilizationBytes *string `json:"diskUtilizationBytes,omitempty"`
}

type VertexAIMetadataStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   VertexAIMetadataStore's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The timestamp of when the MetadataStore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* State information of the MetadataStore. */
	// +optional
	State []MetadatastoreStateStatus `json:"state,omitempty"`

	/* The timestamp of when the MetadataStore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaimetadatastore;gcpvertexaimetadatastores
// +kubebuilder:subresource:status

// VertexAIMetadataStore is the Schema for the vertexai API
// +k8s:openapi-gen=true
type VertexAIMetadataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIMetadataStoreSpec   `json:"spec,omitempty"`
	Status VertexAIMetadataStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIMetadataStoreList contains a list of VertexAIMetadataStore
type VertexAIMetadataStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIMetadataStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIMetadataStore{}, &VertexAIMetadataStoreList{})
}
