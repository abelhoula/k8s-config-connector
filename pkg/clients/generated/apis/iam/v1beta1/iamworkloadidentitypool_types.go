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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IAMWorkloadIdentityPoolSpec struct {
	/* A description of the pool. Cannot exceed 256 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* A display name for the pool. Cannot exceed 32 characters. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type IAMWorkloadIdentityPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   IAMWorkloadIdentityPool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. The state of the pool. Possible values: STATE_UNSPECIFIED, ACTIVE, DELETED */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamworkloadidentitypool;gcpiamworkloadidentitypools
// +kubebuilder:subresource:status

// IAMWorkloadIdentityPool is the Schema for the iam API
// +k8s:openapi-gen=true
type IAMWorkloadIdentityPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMWorkloadIdentityPoolSpec   `json:"spec,omitempty"`
	Status IAMWorkloadIdentityPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMWorkloadIdentityPoolList contains a list of IAMWorkloadIdentityPool
type IAMWorkloadIdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMWorkloadIdentityPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMWorkloadIdentityPool{}, &IAMWorkloadIdentityPoolList{})
}
