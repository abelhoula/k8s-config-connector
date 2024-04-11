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

type VpngatewayVpnInterfaces struct {
	/* Immutable. The numeric ID of this VPN gateway interface. */
	// +optional
	Id *int `json:"id,omitempty"`

	/* Immutable. When this value is present, the VPN Gateway will be used
	for IPsec-encrypted Cloud Interconnect; all Egress or Ingress
	traffic for this VPN Gateway interface will go through the specified
	interconnect attachment resource. Not currently available publicly. */
	// +optional
	InterconnectAttachmentRef *v1alpha1.ResourceRef `json:"interconnectAttachmentRef,omitempty"`

	/* The external IP address for this VPN gateway interface. */
	// +optional
	IpAddress *string `json:"ipAddress,omitempty"`
}

type ComputeVPNGatewaySpec struct {
	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The network this VPN gateway is accepting traffic for. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* Immutable. The region this gateway should sit in. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The stack type for this VPN gateway to identify the IP protocols that are enabled.
	If not specified, IPV4_ONLY will be used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6"]. */
	// +optional
	StackType *string `json:"stackType,omitempty"`

	/* Immutable. A list of interfaces on this VPN gateway. */
	// +optional
	VpnInterfaces []VpngatewayVpnInterfaces `json:"vpnInterfaces,omitempty"`
}

type ComputeVPNGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeVPNGateway's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputevpngateway;gcpcomputevpngateways
// +kubebuilder:subresource:status

// ComputeVPNGateway is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeVPNGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeVPNGatewayList contains a list of ComputeVPNGateway
type ComputeVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeVPNGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeVPNGateway{}, &ComputeVPNGatewayList{})
}
