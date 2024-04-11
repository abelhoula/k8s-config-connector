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

type EntryGcsFilesetSpec struct {
	/* Patterns to identify a set of files in Google Cloud Storage.
	See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
	for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:

	* gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
	* gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
	* gs://bucket_name/file*: matches files prefixed by file in bucket_name
	* gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
	* gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
	* gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
	* gs://bucket_name/a/* /b: matches all files in bucket_name that match a/* /b pattern, such as a/c/b, a/d/b
	* gs://another_bucket/a.txt: matches gs://another_bucket/a.txt. */
	FilePatterns []string `json:"filePatterns"`

	/* Sample files contained in this fileset, not all files contained in this fileset are represented here. */
	// +optional
	SampleGcsFileSpecs []EntrySampleGcsFileSpecs `json:"sampleGcsFileSpecs,omitempty"`
}

type EntrySampleGcsFileSpecs struct {
	/* The full file path. */
	// +optional
	FilePath *string `json:"filePath,omitempty"`

	/* The size of the file, in bytes. */
	// +optional
	SizeBytes *int `json:"sizeBytes,omitempty"`
}

type DataCatalogEntrySpec struct {
	/* Entry description, which can consist of several sentences or paragraphs that describe entry contents. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Display information such as title and description. A short name to identify the entry,
	for example, "Analytics Data - Jan 2011". */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The name of the entry group this entry is in. */
	EntryGroup string `json:"entryGroup"`

	/* Immutable. The id of the entry to create. */
	EntryId string `json:"entryId"`

	/* Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET. */
	// +optional
	GcsFilesetSpec *EntryGcsFilesetSpec `json:"gcsFilesetSpec,omitempty"`

	/* The resource this metadata entry refers to.
	For Google Cloud Platform resources, linkedResource is the full name of the resource.
	For example, the linkedResource for a table resource from BigQuery is:
	//bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	this field is optional and defaults to an empty string. */
	// +optional
	LinkedResource *string `json:"linkedResource,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	attached to it. See
	https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	for what fields this schema can contain. */
	// +optional
	Schema *string `json:"schema,omitempty"`

	/* Immutable. The type of the entry. Only used for Entries with types in the EntryType enum.
	Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType. Possible values: ["FILESET"]. */
	// +optional
	Type *string `json:"type,omitempty"`

	/* This field indicates the entry's source system that Data Catalog does not integrate with.
	userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. */
	// +optional
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty"`

	/* Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	When creating an entry, users should check the enum values first, if nothing matches the entry
	to be created, then provide a custom value, for example "my_special_type".
	userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. */
	// +optional
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty"`
}

type EntryBigqueryDateShardedSpecStatus struct {
	/* The Data Catalog resource name of the dataset entry the current table belongs to, for example,
	projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}. */
	// +optional
	Dataset *string `json:"dataset,omitempty"`

	/* Total number of shards. */
	// +optional
	ShardCount *int `json:"shardCount,omitempty"`

	/* The table name prefix of the shards. The name of any given shard is [tablePrefix]YYYYMMDD,
	for example, for shard MyTable20180101, the tablePrefix is MyTable. */
	// +optional
	TablePrefix *string `json:"tablePrefix,omitempty"`
}

type EntryBigqueryTableSpecStatus struct {
	/* The table source type. */
	// +optional
	TableSourceType *string `json:"tableSourceType,omitempty"`

	/* Spec of a BigQuery table. This field should only be populated if tableSourceType is BIGQUERY_TABLE. */
	// +optional
	TableSpec []EntryTableSpecStatus `json:"tableSpec,omitempty"`

	/* Table view specification. This field should only be populated if tableSourceType is BIGQUERY_VIEW. */
	// +optional
	ViewSpec []EntryViewSpecStatus `json:"viewSpec,omitempty"`
}

type EntryTableSpecStatus struct {
	/* If the table is a dated shard, i.e., with name pattern [prefix]YYYYMMDD, groupedEntry is the
	Data Catalog resource name of the date sharded grouped entry, for example,
	projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}.
	Otherwise, groupedEntry is empty. */
	// +optional
	GroupedEntry *string `json:"groupedEntry,omitempty"`
}

type EntryViewSpecStatus struct {
	/* The query that defines the table view. */
	// +optional
	ViewQuery *string `json:"viewQuery,omitempty"`
}

type DataCatalogEntryStatus struct {
	/* Conditions represent the latest available observations of the
	   DataCatalogEntry's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD.
	Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding. */
	// +optional
	BigqueryDateShardedSpec []EntryBigqueryDateShardedSpecStatus `json:"bigqueryDateShardedSpec,omitempty"`

	/* Specification that applies to a BigQuery table. This is only valid on entries of type TABLE. */
	// +optional
	BigqueryTableSpec []EntryBigqueryTableSpecStatus `json:"bigqueryTableSpec,omitempty"`

	/* This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub. */
	// +optional
	IntegratedSystem *string `json:"integratedSystem,omitempty"`

	/* The Data Catalog resource name of the entry in URL format.
	Example: projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}.
	Note that this Entry and its child resources may not actually be stored in the location in this name. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogentry;gcpdatacatalogentries
// +kubebuilder:subresource:status

// DataCatalogEntry is the Schema for the datacatalog API
// +k8s:openapi-gen=true
type DataCatalogEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogEntrySpec   `json:"spec,omitempty"`
	Status DataCatalogEntryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataCatalogEntryList contains a list of DataCatalogEntry
type DataCatalogEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogEntry{}, &DataCatalogEntryList{})
}
