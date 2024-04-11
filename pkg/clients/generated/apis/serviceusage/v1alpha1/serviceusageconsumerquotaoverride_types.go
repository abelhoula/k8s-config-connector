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

type ServiceUsageConsumerQuotaOverrideSpec struct {
	/* Immutable. If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit. */
	// +optional
	Dimensions map[string]string `json:"dimensions,omitempty"`

	/* If the new quota would decrease the existing quota by more than 10%, the request is rejected.
	If 'force' is 'true', that safety check is ignored. */
	// +optional
	Force *bool `json:"force,omitempty"`

	/* Immutable. The limit on the metric, e.g. '/project/region'.

	~> Make sure that 'limit' is in a format that doesn't start with '1/' or contain curly braces.
	E.g. use '/project/user' instead of '1/{project}/{user}'. */
	Limit string `json:"limit"`

	/* Immutable. The metric that should be limited, e.g. 'compute.googleapis.com/cpus'. */
	Metric string `json:"metric"`

	/* The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota). */
	OverrideValue string `json:"overrideValue"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The service that the metrics belong to, e.g. 'compute.googleapis.com'. */
	Service string `json:"service"`
}

type ServiceUsageConsumerQuotaOverrideStatus struct {
	/* Conditions represent the latest available observations of the
	   ServiceUsageConsumerQuotaOverride's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The server-generated name of the quota override. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpserviceusageconsumerquotaoverride;gcpserviceusageconsumerquotaoverrides
// +kubebuilder:subresource:status

// ServiceUsageConsumerQuotaOverride is the Schema for the serviceusage API
// +k8s:openapi-gen=true
type ServiceUsageConsumerQuotaOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceUsageConsumerQuotaOverrideSpec   `json:"spec,omitempty"`
	Status ServiceUsageConsumerQuotaOverrideStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceUsageConsumerQuotaOverrideList contains a list of ServiceUsageConsumerQuotaOverride
type ServiceUsageConsumerQuotaOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceUsageConsumerQuotaOverride `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceUsageConsumerQuotaOverride{}, &ServiceUsageConsumerQuotaOverrideList{})
}
