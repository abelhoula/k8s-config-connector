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

type StorageDefaultObjectAccessControlSpec struct {
	/* Reference to the bucket. */
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* The entity holding the permission, in one of the following forms:
	* user-{{userId}}
	* user-{{email}} (such as "user-liz@example.com")
	* group-{{groupId}}
	* group-{{email}} (such as "group-example@googlegroups.com")
	* domain-{{domain}} (such as "domain-example.com")
	* project-team-{{projectId}}
	* allUsers
	* allAuthenticatedUsers. */
	Entity string `json:"entity"`

	/* The name of the object, if applied to an object. */
	// +optional
	Object *string `json:"object,omitempty"`

	/* The access permission for the entity. Possible values: ["OWNER", "READER"]. */
	Role string `json:"role"`
}

type DefaultobjectaccesscontrolProjectTeamStatus struct {
	/* The project team associated with the entity. */
	// +optional
	ProjectNumber *string `json:"projectNumber,omitempty"`

	/* The team. Possible values: ["editors", "owners", "viewers"]. */
	// +optional
	Team *string `json:"team,omitempty"`
}

type StorageDefaultObjectAccessControlStatus struct {
	/* Conditions represent the latest available observations of the
	   StorageDefaultObjectAccessControl's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The domain associated with the entity. */
	// +optional
	Domain *string `json:"domain,omitempty"`

	/* The email address associated with the entity. */
	// +optional
	Email *string `json:"email,omitempty"`

	/* The ID for the entity. */
	// +optional
	EntityId *string `json:"entityId,omitempty"`

	/* The content generation of the object, if applied to an object. */
	// +optional
	Generation *int `json:"generation,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The project team associated with the entity. */
	// +optional
	ProjectTeam *DefaultobjectaccesscontrolProjectTeamStatus `json:"projectTeam,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragedefaultobjectaccesscontrol;gcpstoragedefaultobjectaccesscontrols
// +kubebuilder:subresource:status

// StorageDefaultObjectAccessControl is the Schema for the storage API
// +k8s:openapi-gen=true
type StorageDefaultObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageDefaultObjectAccessControlSpec   `json:"spec,omitempty"`
	Status StorageDefaultObjectAccessControlStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageDefaultObjectAccessControlList contains a list of StorageDefaultObjectAccessControl
type StorageDefaultObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageDefaultObjectAccessControl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageDefaultObjectAccessControl{}, &StorageDefaultObjectAccessControlList{})
}
