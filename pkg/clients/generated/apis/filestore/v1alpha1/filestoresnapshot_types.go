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

type FilestoreSnapshotSpec struct {
	/* A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The resource name of the filestore instance. */
	Instance string `json:"instance"`

	/* Immutable. The name of the location of the instance. This can be a region for ENTERPRISE tier instances. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type FilestoreSnapshotStatus struct {
	/* Conditions represent the latest available observations of the
	   FilestoreSnapshot's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time when the snapshot was created in RFC3339 text format. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The amount of bytes needed to allocate a full copy of the snapshot content. */
	// +optional
	FilesystemUsedBytes *string `json:"filesystemUsedBytes,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The snapshot state. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfilestoresnapshot;gcpfilestoresnapshots
// +kubebuilder:subresource:status

// FilestoreSnapshot is the Schema for the filestore API
// +k8s:openapi-gen=true
type FilestoreSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FilestoreSnapshotSpec   `json:"spec,omitempty"`
	Status FilestoreSnapshotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FilestoreSnapshotList contains a list of FilestoreSnapshot
type FilestoreSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FilestoreSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FilestoreSnapshot{}, &FilestoreSnapshotList{})
}
