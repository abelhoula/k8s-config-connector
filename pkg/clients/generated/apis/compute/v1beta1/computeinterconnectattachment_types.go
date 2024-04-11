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

type ComputeInterconnectAttachmentSpec struct {
	/* Whether the VLAN attachment is enabled or disabled.  When using
	PARTNER type this will Pre-Activate the interconnect attachment. */
	// +optional
	AdminEnabled *bool `json:"adminEnabled,omitempty"`

	/* Provisioned bandwidth capacity for the interconnect attachment.
	For attachments of type DEDICATED, the user can set the bandwidth.
	For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	Defaults to BPS_10G Possible values: ["BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G"]. */
	// +optional
	Bandwidth *string `json:"bandwidth,omitempty"`

	/* Immutable. Up to 16 candidate prefixes that can be used to restrict the allocation
	of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	All prefixes must be within link-local address space (169.254.0.0/16)
	and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	an unused /29 from the supplied candidate prefix(es). The request will
	fail if all possible /29s are in use on Google's edge. If not supplied,
	Google will randomly select an unused /29 from all of link-local space. */
	// +optional
	CandidateSubnets []string `json:"candidateSubnets,omitempty"`

	/* An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Desired availability domain for the attachment. Only available for type
	PARTNER, at creation time. For improved reliability, customers should
	configure a pair of attachments with one per availability domain. The
	selected availability domain will be provided to the Partner via the
	pairing key so that the provisioned circuit will lie in the specified
	domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY. */
	// +optional
	EdgeAvailabilityDomain *string `json:"edgeAvailabilityDomain,omitempty"`

	/* Immutable. Indicates the user-supplied encryption option of this interconnect
	attachment. Can only be specified at attachment creation for PARTNER or
	DEDICATED attachments.

	* NONE - This is the default value, which means that the VLAN attachment
	carries unencrypted traffic. VMs are able to send traffic to, or receive
	traffic from, such a VLAN attachment.

	* IPSEC - The VLAN attachment carries only encrypted traffic that is
	encrypted by an IPsec device, such as an HA VPN gateway or third-party
	IPsec VPN. VMs cannot directly send traffic to, or receive traffic from,
	such a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN
	attachment must be created with this option. Default value: "NONE" Possible values: ["NONE", "IPSEC"]. */
	// +optional
	Encryption *string `json:"encryption,omitempty"`

	/* Immutable. URL of the underlying Interconnect object that this attachment's
	traffic will traverse through. Required if type is DEDICATED, must not
	be set if type is PARTNER. */
	// +optional
	Interconnect *string `json:"interconnect,omitempty"`

	// +optional
	IpsecInternalAddresses []v1alpha1.ResourceRef `json:"ipsecInternalAddresses,omitempty"`

	/* Maximum Transmission Unit (MTU), in bytes, of packets passing through
	this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440. */
	// +optional
	Mtu *string `json:"mtu,omitempty"`

	/* Region where the regional interconnect attachment resides. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The Cloud Router to be used for dynamic routing. This router must
	be in the same region as this ComputeInterconnectAttachment. The
	ComputeInterconnectAttachment will automatically connect the
	interconnect to the network & region within which the Cloud Router
	is configured. */
	RouterRef v1alpha1.ResourceRef `json:"routerRef"`

	/* Immutable. The type of InterconnectAttachment you wish to create. Defaults to
	DEDICATED. Possible values: ["DEDICATED", "PARTNER", "PARTNER_PROVIDER"]. */
	// +optional
	Type *string `json:"type,omitempty"`

	/* Immutable. The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	using PARTNER type this will be managed upstream. */
	// +optional
	VlanTag8021q *int `json:"vlanTag8021q,omitempty"`
}

type InterconnectattachmentPrivateInterconnectInfoStatus struct {
	/* 802.1q encapsulation tag to be used for traffic between
	Google and the customer, going to and from this network and region. */
	// +optional
	Tag8021q *int `json:"tag8021q,omitempty"`
}

type ComputeInterconnectAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeInterconnectAttachment's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* IPv4 address + prefix length to be configured on Cloud Router
	Interface for this interconnect attachment. */
	// +optional
	CloudRouterIpAddress *string `json:"cloudRouterIpAddress,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* IPv4 address + prefix length to be configured on the customer
	router subinterface for this interconnect attachment. */
	// +optional
	CustomerRouterIpAddress *string `json:"customerRouterIpAddress,omitempty"`

	/* Google reference ID, to be used when raising support tickets with
	Google or otherwise to debug backend connectivity issues. */
	// +optional
	GoogleReferenceId *string `json:"googleReferenceId,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* [Output only for type PARTNER. Not present for DEDICATED]. The opaque
	identifier of an PARTNER attachment used to initiate provisioning with
	a selected partner. Of the form "XXXXX/region/domain". */
	// +optional
	PairingKey *string `json:"pairingKey,omitempty"`

	/* [Output only for type PARTNER. Not present for DEDICATED]. Optional
	BGP ASN for the router that should be supplied by a layer 3 Partner if
	they configured BGP on behalf of the customer. */
	// +optional
	PartnerAsn *string `json:"partnerAsn,omitempty"`

	/* Information specific to an InterconnectAttachment. This property
	is populated if the interconnect that this is attached to is of type DEDICATED. */
	// +optional
	PrivateInterconnectInfo *InterconnectattachmentPrivateInterconnectInfoStatus `json:"privateInterconnectInfo,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* [Output Only] The current state of this attachment's functionality. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinterconnectattachment;gcpcomputeinterconnectattachments
// +kubebuilder:subresource:status

// ComputeInterconnectAttachment is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInterconnectAttachmentList contains a list of ComputeInterconnectAttachment
type ComputeInterconnectAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInterconnectAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInterconnectAttachment{}, &ComputeInterconnectAttachmentList{})
}
