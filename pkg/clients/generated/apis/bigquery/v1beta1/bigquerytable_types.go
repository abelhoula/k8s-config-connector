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

type TableAvroOptions struct {
	/* If sourceFormat is set to "AVRO", indicates whether to interpret logical types as the corresponding BigQuery data type (for example, TIMESTAMP), instead of using the raw type (for example, INTEGER). */
	UseAvroLogicalTypes bool `json:"useAvroLogicalTypes"`
}

type TableCsvOptions struct {
	/* Indicates if BigQuery should accept rows that are missing trailing optional columns. */
	// +optional
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty"`

	/* Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file. The default value is false. */
	// +optional
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty"`

	/* The character encoding of the data. The supported values are UTF-8 or ISO-8859-1. */
	// +optional
	Encoding *string `json:"encoding,omitempty"`

	/* The separator for fields in a CSV file. */
	// +optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty"`

	Quote string `json:"quote"`

	/* The number of rows at the top of a CSV file that BigQuery will skip when reading the data. */
	// +optional
	SkipLeadingRows *int `json:"skipLeadingRows,omitempty"`
}

type TableEncryptionConfiguration struct {
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef"`

	/* The self link or full name of the kms key version used to encrypt this table. */
	// +optional
	KmsKeyVersion *string `json:"kmsKeyVersion,omitempty"`
}

type TableExternalDataConfiguration struct {
	/* Let BigQuery try to autodetect the schema and format of the table. */
	Autodetect bool `json:"autodetect"`

	/* Additional options if source_format is set to "AVRO". */
	// +optional
	AvroOptions *TableAvroOptions `json:"avroOptions,omitempty"`

	/* The compression type of the data source. Valid values are "NONE" or "GZIP". */
	// +optional
	Compression *string `json:"compression,omitempty"`

	/* The connection specifying the credentials to be used to read external storage, such as Azure Blob, Cloud Storage, or S3. The connectionId can have the form "{{project}}.{{location}}.{{connection_id}}" or "projects/{{project}}/locations/{{location}}/connections/{{connection_id}}". */
	// +optional
	ConnectionId *string `json:"connectionId,omitempty"`

	/* Additional properties to set if source_format is set to "CSV". */
	// +optional
	CsvOptions *TableCsvOptions `json:"csvOptions,omitempty"`

	/* Specifies how source URIs are interpreted for constructing the file set to load.  By default source URIs are expanded against the underlying storage.  Other options include specifying manifest files. Only applicable to object storage systems. */
	// +optional
	FileSetSpecType *string `json:"fileSetSpecType,omitempty"`

	/* Additional options if source_format is set to "GOOGLE_SHEETS". */
	// +optional
	GoogleSheetsOptions *TableGoogleSheetsOptions `json:"googleSheetsOptions,omitempty"`

	/* When set, configures hive partitioning support. Not all storage formats support hive partitioning -- requesting hive partitioning on an unsupported format will lead to an error, as will providing an invalid specification. */
	// +optional
	HivePartitioningOptions *TableHivePartitioningOptions `json:"hivePartitioningOptions,omitempty"`

	/* Indicates if BigQuery should allow extra values that are not represented in the table schema. If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false. */
	// +optional
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty"`

	/* Additional properties to set if sourceFormat is set to JSON.". */
	// +optional
	JsonOptions *TableJsonOptions `json:"jsonOptions,omitempty"`

	/* The maximum number of bad records that BigQuery can ignore when reading data. */
	// +optional
	MaxBadRecords *int `json:"maxBadRecords,omitempty"`

	/* Metadata Cache Mode for the table. Set this to enable caching of metadata from external data source. */
	// +optional
	MetadataCacheMode *string `json:"metadataCacheMode,omitempty"`

	/* Object Metadata is used to create Object Tables. Object Tables contain a listing of objects (with their metadata) found at the sourceUris. If ObjectMetadata is set, sourceFormat should be omitted. */
	// +optional
	ObjectMetadata *string `json:"objectMetadata,omitempty"`

	/* Additional properties to set if sourceFormat is set to PARQUET.". */
	// +optional
	ParquetOptions *TableParquetOptions `json:"parquetOptions,omitempty"`

	/* When creating an external table, the user can provide a reference file with the table schema. This is enabled for the following formats: AVRO, PARQUET, ORC. */
	// +optional
	ReferenceFileSchemaUri *string `json:"referenceFileSchemaUri,omitempty"`

	/* Immutable. A JSON schema for the external table. Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables. */
	// +optional
	Schema *string `json:"schema,omitempty"`

	/* Please see sourceFormat under ExternalDataConfiguration in Bigquery's public API documentation (https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#externaldataconfiguration) for supported formats. To use "GOOGLE_SHEETS" the scopes must include "googleapis.com/auth/drive.readonly". */
	// +optional
	SourceFormat *string `json:"sourceFormat,omitempty"`

	/* A list of the fully-qualified URIs that point to your data in Google Cloud. */
	SourceUris []string `json:"sourceUris"`
}

type TableGoogleSheetsOptions struct {
	/* Range of a sheet to query from. Only used when non-empty. At least one of range or skip_leading_rows must be set. Typical format: "sheet_name!top_left_cell_id:bottom_right_cell_id" For example: "sheet1!A1:B20". */
	// +optional
	Range *string `json:"range,omitempty"`

	/* The number of rows at the top of the sheet that BigQuery will skip when reading the data. At least one of range or skip_leading_rows must be set. */
	// +optional
	SkipLeadingRows *int `json:"skipLeadingRows,omitempty"`
}

type TableHivePartitioningOptions struct {
	/* When set, what mode of hive partitioning to use when reading data. */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/* If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified. */
	// +optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	/* When hive partition detection is requested, a common for all source uris must be required. The prefix must end immediately before the partition key encoding begins. */
	// +optional
	SourceUriPrefix *string `json:"sourceUriPrefix,omitempty"`
}

type TableJsonOptions struct {
	/* The character encoding of the data. The supported values are UTF-8, UTF-16BE, UTF-16LE, UTF-32BE, and UTF-32LE. The default value is UTF-8. */
	// +optional
	Encoding *string `json:"encoding,omitempty"`
}

type TableMaterializedView struct {
	/* Specifies if BigQuery should automatically refresh materialized view when the base table is updated. The default is true. */
	// +optional
	EnableRefresh *bool `json:"enableRefresh,omitempty"`

	/* Immutable. A query whose result is persisted. */
	Query string `json:"query"`

	/* Specifies maximum frequency at which this materialized view will be refreshed. The default is 1800000. */
	// +optional
	RefreshIntervalMs *int `json:"refreshIntervalMs,omitempty"`
}

type TableParquetOptions struct {
	/* Indicates whether to use schema inference specifically for Parquet LIST logical type. */
	// +optional
	EnableListInference *bool `json:"enableListInference,omitempty"`

	/* Indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default. */
	// +optional
	EnumAsString *bool `json:"enumAsString,omitempty"`
}

type TableRange struct {
	/* End of the range partitioning, exclusive. */
	End int `json:"end"`

	/* The width of each range within the partition. */
	Interval int `json:"interval"`

	/* Start of the range partitioning, inclusive. */
	Start int `json:"start"`
}

type TableRangePartitioning struct {
	/* Immutable. The field used to determine how to create a range-based partition. */
	Field string `json:"field"`

	/* Information required to partition based on ranges. Structure is documented below. */
	Range TableRange `json:"range"`
}

type TableTimePartitioning struct {
	/* Number of milliseconds for which to keep the storage for a partition. */
	// +optional
	ExpirationMs *int `json:"expirationMs,omitempty"`

	/* Immutable. The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time. */
	// +optional
	Field *string `json:"field,omitempty"`

	/* If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified. */
	// +optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	/* The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively. */
	Type string `json:"type"`
}

type TableView struct {
	/* A query that BigQuery executes when the view is referenced. */
	Query string `json:"query"`

	/* Specifies whether to use BigQuery's legacy SQL for this view. The default value is true. If set to false, the view will use BigQuery's standard SQL. */
	// +optional
	UseLegacySql *bool `json:"useLegacySql,omitempty"`
}

type BigQueryTableSpec struct {
	/* Specifies column names to use for data clustering. Up to four top-level columns are allowed, and should be specified in descending priority order. */
	// +optional
	Clustering []string `json:"clustering,omitempty"`

	DatasetRef v1alpha1.ResourceRef `json:"datasetRef"`

	/* The field description. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Specifies how the table should be encrypted. If left blank, the table will be encrypted with a Google-managed key; that process is transparent to the user. */
	// +optional
	EncryptionConfiguration *TableEncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	/* The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. */
	// +optional
	ExpirationTime *int `json:"expirationTime,omitempty"`

	/* Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table. */
	// +optional
	ExternalDataConfiguration *TableExternalDataConfiguration `json:"externalDataConfiguration,omitempty"`

	/* A descriptive name for the table. */
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	/* If specified, configures this table as a materialized view. */
	// +optional
	MaterializedView *TableMaterializedView `json:"materializedView,omitempty"`

	/* The maximum staleness of data that could be returned when the table (or stale MV) is queried. Staleness encoded as a string encoding of sql IntervalValue type. */
	// +optional
	MaxStaleness *string `json:"maxStaleness,omitempty"`

	/* If specified, configures range-based partitioning for this table. */
	// +optional
	RangePartitioning *TableRangePartitioning `json:"rangePartitioning,omitempty"`

	/* Immutable. Optional. The tableId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* A JSON schema for the table. */
	// +optional
	Schema *string `json:"schema,omitempty"`

	/* If specified, configures time-based partitioning for this table. */
	// +optional
	TimePartitioning *TableTimePartitioning `json:"timePartitioning,omitempty"`

	/* If specified, configures this table as a view. */
	// +optional
	View *TableView `json:"view,omitempty"`
}

type BigQueryTableStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryTable's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time when this table was created, in milliseconds since the epoch. */
	// +optional
	CreationTime *int `json:"creationTime,omitempty"`

	/* A hash of the resource. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* The time when this table was last modified, in milliseconds since the epoch. */
	// +optional
	LastModifiedTime *int `json:"lastModifiedTime,omitempty"`

	/* The geographic location where the table resides. This value is inherited from the dataset. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* The geographic location where the table resides. This value is inherited from the dataset. */
	// +optional
	NumBytes *int `json:"numBytes,omitempty"`

	/* The number of bytes in the table that are considered "long-term storage". */
	// +optional
	NumLongTermBytes *int `json:"numLongTermBytes,omitempty"`

	/* The number of rows of data in this table, excluding any data in the streaming buffer. */
	// +optional
	NumRows *int `json:"numRows,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The URI of the created resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* Describes the table type. */
	// +optional
	Type *string `json:"type,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryTable is the Schema for the bigquery API
// +k8s:openapi-gen=true
type BigQueryTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryTableSpec   `json:"spec,omitempty"`
	Status BigQueryTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryTableList contains a list of BigQueryTable
type BigQueryTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryTable{}, &BigQueryTableList{})
}
