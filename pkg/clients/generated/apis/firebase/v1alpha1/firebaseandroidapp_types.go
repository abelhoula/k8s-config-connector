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

type FirebaseAndroidAppSpec struct {
	/* The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
	If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
	This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned. */
	// +optional
	ApiKeyId *string `json:"apiKeyId,omitempty"`

	// +optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	/* The user-assigned display name of the AndroidApp. */
	DisplayName string `json:"displayName"`

	/* Immutable. The canonical package name of the Android app as would appear in the Google Play
	Developer Console. */
	// +optional
	PackageName *string `json:"packageName,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The service-generated appId of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The SHA1 certificate hashes for the AndroidApp. */
	// +optional
	Sha1Hashes []string `json:"sha1Hashes,omitempty"`

	/* The SHA256 certificate hashes for the AndroidApp. */
	// +optional
	Sha256Hashes []string `json:"sha256Hashes,omitempty"`
}

type FirebaseAndroidAppStatus struct {
	/* Conditions represent the latest available observations of the
	   FirebaseAndroidApp's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The globally unique, Firebase-assigned identifier of the AndroidApp.
	This identifier should be treated as an opaque token, as the data format is not specified. */
	// +optional
	AppId *string `json:"appId,omitempty"`

	/* This checksum is computed by the server based on the value of other fields, and it may be sent
	with update requests to ensure the client has an up-to-date value before proceeding. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* The fully qualified resource name of the AndroidApp, for example:
	projects/projectId/androidApps/appId. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirebaseAndroidApp is the Schema for the firebase API
// +k8s:openapi-gen=true
type FirebaseAndroidApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseAndroidAppSpec   `json:"spec,omitempty"`
	Status FirebaseAndroidAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirebaseAndroidAppList contains a list of FirebaseAndroidApp
type FirebaseAndroidAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirebaseAndroidApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirebaseAndroidApp{}, &FirebaseAndroidAppList{})
}
