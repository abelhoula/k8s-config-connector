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

type AppprofileSingleClusterRouting struct {
	/* If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	It is unsafe to send these requests to the same table/row/column in multiple clusters. */
	// +optional
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty"`

	/* The cluster to which read/write requests should be routed. */
	ClusterId string `json:"clusterId"`
}

type AppprofileStandardIsolation struct {
	/* The priority of requests sent using this app profile. Possible values: ["PRIORITY_LOW", "PRIORITY_MEDIUM", "PRIORITY_HIGH"]. */
	Priority string `json:"priority"`
}

type BigtableAppProfileSpec struct {
	/* Long form description of the use case for this app profile. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The instance to create the app profile within. */
	// +optional
	InstanceRef *v1alpha1.ResourceRef `json:"instanceRef,omitempty"`

	/* The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all clusters are eligible. */
	// +optional
	MultiClusterRoutingClusterIds []string `json:"multiClusterRoutingClusterIds,omitempty"`

	/* If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	consistency to improve availability. */
	// +optional
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty"`

	/* Immutable. Optional. The appProfileId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Use a single-cluster routing policy. */
	// +optional
	SingleClusterRouting *AppprofileSingleClusterRouting `json:"singleClusterRouting,omitempty"`

	/* The standard options used for isolating this app profile's traffic from other use cases. */
	// +optional
	StandardIsolation *AppprofileStandardIsolation `json:"standardIsolation,omitempty"`
}

type BigtableAppProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   BigtableAppProfile's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The unique name of the requested app profile. Values are of the form 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableappprofile;gcpbigtableappprofiles
// +kubebuilder:subresource:status

// BigtableAppProfile is the Schema for the bigtable API
// +k8s:openapi-gen=true
type BigtableAppProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableAppProfileSpec   `json:"spec,omitempty"`
	Status BigtableAppProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigtableAppProfileList contains a list of BigtableAppProfile
type BigtableAppProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableAppProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableAppProfile{}, &BigtableAppProfileList{})
}
