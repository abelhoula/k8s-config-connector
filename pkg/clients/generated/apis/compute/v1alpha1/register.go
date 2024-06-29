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

// Package v1alpha1 contains API Schema definitions for the compute v1alpha1 API group.
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/pkg/apis/compute
// +k8s:defaulter-gen=TypeMeta
// +groupName=compute.cnrm.cloud.google.com
package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is the group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme

	ComputeAutoscalerGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeAutoscaler{}).Name(),
	}

	ComputeBackendBucketSignedURLKeyGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeBackendBucketSignedURLKey{}).Name(),
	}

	ComputeBackendServiceSignedURLKeyGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeBackendServiceSignedURLKey{}).Name(),
	}

	ComputeDiskResourcePolicyAttachmentGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeDiskResourcePolicyAttachment{}).Name(),
	}

	ComputeGlobalNetworkEndpointGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeGlobalNetworkEndpoint{}).Name(),
	}

	ComputeGlobalNetworkEndpointGroupGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeGlobalNetworkEndpointGroup{}).Name(),
	}

	ComputeInstanceGroupNamedPortGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeInstanceGroupNamedPort{}).Name(),
	}

	ComputeMachineImageGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeMachineImage{}).Name(),
	}

	ComputeNetworkEndpointGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeNetworkEndpoint{}).Name(),
	}

	ComputeNetworkFirewallPolicyRuleGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeNetworkFirewallPolicyRule{}).Name(),
	}

	ComputeNetworkPeeringRoutesConfigGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeNetworkPeeringRoutesConfig{}).Name(),
	}

	ComputeOrganizationSecurityPolicyGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeOrganizationSecurityPolicy{}).Name(),
	}

	ComputeOrganizationSecurityPolicyAssociationGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeOrganizationSecurityPolicyAssociation{}).Name(),
	}

	ComputeOrganizationSecurityPolicyRuleGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeOrganizationSecurityPolicyRule{}).Name(),
	}

	ComputePerInstanceConfigGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputePerInstanceConfig{}).Name(),
	}

	ComputeRegionAutoscalerGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeRegionAutoscaler{}).Name(),
	}

	ComputeRegionDiskResourcePolicyAttachmentGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeRegionDiskResourcePolicyAttachment{}).Name(),
	}

	ComputeRegionPerInstanceConfigGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeRegionPerInstanceConfig{}).Name(),
	}

	ComputeRegionSSLPolicyGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(ComputeRegionSSLPolicy{}).Name(),
	}

	computeAPIVersion = SchemeGroupVersion.String()
)
