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

type TagtemplateAllowedValues struct {
	/* The display name of the enum value. */
	DisplayName string `json:"displayName"`
}

type TagtemplateEnumType struct {
	/* The set of allowed values for this enum. The display names of the
	values must be case-insensitively unique within this set. Currently,
	enum values can only be added to the list of allowed values. Deletion
	and renaming of enum values are not supported.
	Can have up to 500 allowed values. */
	AllowedValues []TagtemplateAllowedValues `json:"allowedValues"`
}

type TagtemplateFields struct {
	/* A description for this field. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The display name for this field. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	FieldId string `json:"fieldId"`

	/* Whether this is a required field. Defaults to false. */
	// +optional
	IsRequired *bool `json:"isRequired,omitempty"`

	/* The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* The order of this field with respect to other fields in this tag template.
	A higher value indicates a more important field. The value can be negative.
	Multiple fields can have the same order, and field orders within a tag do not have to be sequential. */
	// +optional
	Order *int `json:"order,omitempty"`

	/* The type of value this tag field can contain. */
	Type TagtemplateType `json:"type"`
}

type TagtemplateType struct {
	/* Represents an enum type.
	Exactly one of 'primitive_type' or 'enum_type' must be set. */
	// +optional
	EnumType *TagtemplateEnumType `json:"enumType,omitempty"`

	/* Represents primitive types - string, bool etc.
	Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: ["DOUBLE", "STRING", "BOOL", "TIMESTAMP"]. */
	// +optional
	PrimitiveType *string `json:"primitiveType,omitempty"`
}

type DataCatalogTagTemplateSpec struct {
	/* The display name for this template. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields. The change of field_id will be resulting in re-creating of field. The change of primitive_type will be resulting in re-creating of field, however if the field is a required, you cannot update it. */
	Fields []TagtemplateFields `json:"fields"`

	/* This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template. */
	// +optional
	ForceDelete *bool `json:"forceDelete,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Template location region. */
	// +optional
	Region *string `json:"region,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The id of the tag template to create. */
	TagTemplateId string `json:"tagTemplateId"`
}

type DataCatalogTagTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   DataCatalogTagTemplate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource name of the tag template in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogtagtemplate;gcpdatacatalogtagtemplates
// +kubebuilder:subresource:status

// DataCatalogTagTemplate is the Schema for the datacatalog API
// +k8s:openapi-gen=true
type DataCatalogTagTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogTagTemplateSpec   `json:"spec,omitempty"`
	Status DataCatalogTagTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataCatalogTagTemplateList contains a list of DataCatalogTagTemplate
type DataCatalogTagTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogTagTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogTagTemplate{}, &DataCatalogTagTemplateList{})
}
