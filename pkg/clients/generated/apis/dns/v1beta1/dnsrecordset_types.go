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

type RecordsetBackupGeo struct {
	/* For A and AAAA types only. The list of targets to be health checked. These can be specified along with `rrdatas` within this item. */
	// +optional
	HealthCheckedTargets *RecordsetHealthCheckedTargets `json:"healthCheckedTargets,omitempty"`

	/* The location name defined in Google Cloud. */
	Location string `json:"location"`

	// +optional
	RrdatasRefs []RecordsetRrdatasRefs `json:"rrdatasRefs,omitempty"`
}

type RecordsetGeo struct {
	/* For A and AAAA types only. The list of targets to be health checked. These can be specified along with `rrdatas` within this item. */
	// +optional
	HealthCheckedTargets *RecordsetHealthCheckedTargets `json:"healthCheckedTargets,omitempty"`

	/* The location name defined in Google Cloud. */
	Location string `json:"location"`

	// +optional
	RrdatasRefs []RecordsetRrdatasRefs `json:"rrdatasRefs,omitempty"`
}

type RecordsetHealthCheckedTargets struct {
	/* The list of internal load balancers to health check. */
	InternalLoadBalancers []RecordsetInternalLoadBalancers `json:"internalLoadBalancers"`
}

type RecordsetInternalLoadBalancers struct {
	IpAddressRef v1alpha1.ResourceRef `json:"ipAddressRef"`

	/* The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]. */
	IpProtocol string `json:"ipProtocol"`

	/* The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb", "regionalL7ilb", "globalL7ilb"]. */
	LoadBalancerType string `json:"loadBalancerType"`

	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* The configured port of the load balancer. */
	Port string `json:"port"`

	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	// +optional
	RegionRef *v1alpha1.ResourceRef `json:"regionRef,omitempty"`
}

type RecordsetPrimary struct {
	/* The list of internal load balancers to health check. */
	InternalLoadBalancers []RecordsetInternalLoadBalancers `json:"internalLoadBalancers"`
}

type RecordsetPrimaryBackup struct {
	/* The backup geo targets, which provide a regional failover policy for the otherwise global primary targets. */
	BackupGeo []RecordsetBackupGeo `json:"backupGeo"`

	/* Specifies whether to enable fencing for backup geo queries. */
	// +optional
	EnableGeoFencingForBackups *bool `json:"enableGeoFencingForBackups,omitempty"`

	/* The list of global primary targets to be health checked. */
	Primary RecordsetPrimary `json:"primary"`

	/* Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy. */
	// +optional
	TrickleRatio *float64 `json:"trickleRatio,omitempty"`
}

type RecordsetRoutingPolicy struct {
	/* Specifies whether to enable fencing for geo queries. */
	// +optional
	EnableGeoFencing *bool `json:"enableGeoFencing,omitempty"`

	/* The configuration for Geo location based routing policy. */
	// +optional
	Geo []RecordsetGeo `json:"geo,omitempty"`

	/* The configuration for a primary-backup policy with global to regional failover. Queries are responded to with the global primary targets, but if none of the primary targets are healthy, then we fallback to a regional failover policy. */
	// +optional
	PrimaryBackup *RecordsetPrimaryBackup `json:"primaryBackup,omitempty"`

	/* The configuration for Weighted Round Robin based routing policy. */
	// +optional
	Wrr []RecordsetWrr `json:"wrr,omitempty"`
}

type RecordsetRrdatasRefs struct {
	/* Allowed value: The `address` field of a `ComputeAddress` resource. */
	// +optional
	External *string `json:"external,omitempty"`

	/* Kind of the referent. Allowed values: ComputeAddress */
	// +optional
	Kind *string `json:"kind,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

type RecordsetWrr struct {
	/* The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of `rrdatas` or `health_checked_targets` can be set. */
	// +optional
	HealthCheckedTargets *RecordsetHealthCheckedTargets `json:"healthCheckedTargets,omitempty"`

	// +optional
	RrdatasRefs []RecordsetRrdatasRefs `json:"rrdatasRefs,omitempty"`

	/* The ratio of traffic routed to the target. */
	Weight float64 `json:"weight"`
}

type DNSRecordSetSpec struct {
	ManagedZoneRef v1alpha1.ResourceRef `json:"managedZoneRef"`

	/* Immutable. The DNS name this record set will apply to. */
	Name string `json:"name"`

	/* The configuration for steering traffic based on query. You can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type. */
	// +optional
	RoutingPolicy *RecordsetRoutingPolicy `json:"routingPolicy,omitempty"`

	/* DEPRECATED. Although this field is still available, there is limited support. We recommend that you use `spec.rrdatasRefs` instead. */
	// +optional
	Rrdatas []string `json:"rrdatas,omitempty"`

	// +optional
	RrdatasRefs []RecordsetRrdatasRefs `json:"rrdatasRefs,omitempty"`

	/* The time-to-live of this record set (seconds). */
	// +optional
	Ttl *int `json:"ttl,omitempty"`

	/* The DNS record set type. */
	Type string `json:"type"`
}

type DNSRecordSetStatus struct {
	/* Conditions represent the latest available observations of the
	   DNSRecordSet's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdnsrecordset;gcpdnsrecordsets
// +kubebuilder:subresource:status

// DNSRecordSet is the Schema for the dns API
// +k8s:openapi-gen=true
type DNSRecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSRecordSetSpec   `json:"spec,omitempty"`
	Status DNSRecordSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSRecordSetList contains a list of DNSRecordSet
type DNSRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSRecordSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSRecordSet{}, &DNSRecordSetList{})
}
