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

type Hl7v2storeNotificationConfig struct {
	/* The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
	PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail. */
	PubsubTopic string `json:"pubsubTopic"`
}

type Hl7v2storeNotificationConfigs struct {
	/* Restricts notifications sent for messages matching a filter. If this is empty, all messages
	are matched. Syntax: https://cloud.google.com/appengine/docs/standard/python/search/query_strings

	Fields/functions available for filtering are:

	* messageType, from the MSH-9.1 field. For example, NOT messageType = "ADT".
	* send_date or sendDate, the YYYY-MM-DD date the message was sent in the dataset's timeZone, from the MSH-7 segment. For example, send_date < "2017-01-02".
	* sendTime, the timestamp when the message was sent, using the RFC3339 time format for comparisons, from the MSH-7 segment. For example, sendTime < "2017-01-02T00:00:00-05:00".
	* sendFacility, the care center that the message came from, from the MSH-4 segment. For example, sendFacility = "ABC".
	* PatientId(value, type), which matches if the message lists a patient having an ID of the given value and type in the PID-2, PID-3, or PID-4 segments. For example, PatientId("123456", "MRN").
	* labels.x, a string value of the label with key x as set using the Message.labels map. For example, labels."priority"="high". The operator :* can be used to assert the existence of a label. For example, labels."priority":*. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
	PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.

	If a notification cannot be published to Cloud Pub/Sub, errors will be logged to Stackdriver. */
	PubsubTopic string `json:"pubsubTopic"`
}

type Hl7v2storeParserConfig struct {
	/* Determines whether messages with no header are allowed. */
	// +optional
	AllowNullHeader *bool `json:"allowNullHeader,omitempty"`

	/* JSON encoded string for schemas used to parse messages in this
	store if schematized parsing is desired. */
	// +optional
	Schema *string `json:"schema,omitempty"`

	/* Byte(s) to be used as the segment terminator. If this is unset, '\r' will be used as segment terminator.

	A base64-encoded string. */
	// +optional
	SegmentTerminator *string `json:"segmentTerminator,omitempty"`

	/* Immutable. The version of the unschematized parser to be used when a custom 'schema' is not set. Default value: "V1" Possible values: ["V1", "V2", "V3"]. */
	// +optional
	Version *string `json:"version,omitempty"`
}

type HealthcareHL7V2StoreSpec struct {
	/* Immutable. Identifies the dataset addressed by this request. Must be in the format
	'projects/{project}/locations/{location}/datasets/{dataset}'. */
	Dataset string `json:"dataset"`

	/* DEPRECATED. `notification_config` is deprecated. Use `notification_configs` instead. A nested object resource. */
	// +optional
	NotificationConfig *Hl7v2storeNotificationConfig `json:"notificationConfig,omitempty"`

	/* A list of notification configs. Each configuration uses a filter to determine whether to publish a
	message (both Ingest & Create) on the corresponding notification destination. Only the message name
	is sent as part of the notification. Supplied by the client. */
	// +optional
	NotificationConfigs []Hl7v2storeNotificationConfigs `json:"notificationConfigs,omitempty"`

	/* A nested object resource. */
	// +optional
	ParserConfig *Hl7v2storeParserConfig `json:"parserConfig,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type HealthcareHL7V2StoreStatus struct {
	/* Conditions represent the latest available observations of the
	   HealthcareHL7V2Store's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The fully qualified name of this dataset. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcphealthcarehl7v2store;gcphealthcarehl7v2stores
// +kubebuilder:subresource:status

// HealthcareHL7V2Store is the Schema for the healthcare API
// +k8s:openapi-gen=true
type HealthcareHL7V2Store struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthcareHL7V2StoreSpec   `json:"spec,omitempty"`
	Status HealthcareHL7V2StoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HealthcareHL7V2StoreList contains a list of HealthcareHL7V2Store
type HealthcareHL7V2StoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareHL7V2Store `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HealthcareHL7V2Store{}, &HealthcareHL7V2StoreList{})
}
