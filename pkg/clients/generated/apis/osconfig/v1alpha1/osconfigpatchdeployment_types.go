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

type PatchdeploymentApt struct {
	/* Immutable. List of packages to exclude from update. These packages will be excluded. */
	// +optional
	Excludes []string `json:"excludes,omitempty"`

	/* Immutable. An exclusive list of packages to be updated. These are the only packages that will be updated.
	If these packages are not installed, they will be ignored. This field cannot be specified with
	any other patch configuration fields. */
	// +optional
	ExclusivePackages []string `json:"exclusivePackages,omitempty"`

	/* Immutable. By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead. Possible values: ["DIST", "UPGRADE"]. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type PatchdeploymentDisruptionBudget struct {
	/* Immutable. Specifies a fixed value. */
	// +optional
	Fixed *int `json:"fixed,omitempty"`

	/* Immutable. Specifies the relative value defined as a percentage, which will be multiplied by a reference value. */
	// +optional
	Percentage *int `json:"percentage,omitempty"`
}

type PatchdeploymentGcsObject struct {
	/* Immutable. Bucket of the Cloud Storage object. */
	Bucket string `json:"bucket"`

	/* Immutable. Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change. */
	GenerationNumber string `json:"generationNumber"`

	/* Immutable. Name of the Cloud Storage object. */
	Object string `json:"object"`
}

type PatchdeploymentGoo struct {
	/* Immutable. goo update settings. Use this setting to override the default goo patch rules. */
	Enabled bool `json:"enabled"`
}

type PatchdeploymentGroupLabels struct {
	/* Immutable. Compute Engine instance labels that must be present for a VM instance to be targeted by this filter. */
	Labels map[string]string `json:"labels"`
}

type PatchdeploymentInstanceFilter struct {
	/* Immutable. Target all VM instances in the project. If true, no other criteria is permitted. */
	// +optional
	All *bool `json:"all,omitempty"`

	/* Immutable. Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances. */
	// +optional
	GroupLabels []PatchdeploymentGroupLabels `json:"groupLabels,omitempty"`

	/* Immutable. Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group
	VMs when targeting configs, for example prefix="prod-". */
	// +optional
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty"`

	/* Immutable. Targets any of the VM instances specified. Instances are specified by their URI in the 'form zones/{{zone}}/instances/{{instance_name}}',
	'projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}', or
	'https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}'. */
	// +optional
	Instances []string `json:"instances,omitempty"`

	/* Immutable. Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone. */
	// +optional
	Zones []string `json:"zones,omitempty"`
}

type PatchdeploymentLinuxExecStepConfig struct {
	/* Immutable. Defaults to [0]. A list of possible return values that the execution can return to indicate a success. */
	// +optional
	AllowedSuccessCodes []int `json:"allowedSuccessCodes,omitempty"`

	/* Immutable. A Cloud Storage object containing the executable. */
	// +optional
	GcsObject *PatchdeploymentGcsObject `json:"gcsObject,omitempty"`

	/* Immutable. The script interpreter to use to run the script. If no interpreter is specified the script will
	be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]. */
	// +optional
	Interpreter *string `json:"interpreter,omitempty"`

	/* Immutable. An absolute path to the executable on the VM. */
	// +optional
	LocalPath *string `json:"localPath,omitempty"`
}

type PatchdeploymentMonthly struct {
	/* Immutable. One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.
	Months without the target day will be skipped. For example, a schedule to run "every month on the 31st"
	will not run in February, April, June, etc. */
	// +optional
	MonthDay *int `json:"monthDay,omitempty"`

	/* Immutable. Week day in a month. */
	// +optional
	WeekDayOfMonth *PatchdeploymentWeekDayOfMonth `json:"weekDayOfMonth,omitempty"`
}

type PatchdeploymentOneTimeSchedule struct {
	/* Immutable. The desired patch job execution time. A timestamp in RFC3339 UTC "Zulu" format,
	accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	ExecuteTime string `json:"executeTime"`
}

type PatchdeploymentPatchConfig struct {
	/* Immutable. Apt update settings. Use this setting to override the default apt patch rules. */
	// +optional
	Apt *PatchdeploymentApt `json:"apt,omitempty"`

	/* Immutable. goo update settings. Use this setting to override the default goo patch rules. */
	// +optional
	Goo *PatchdeploymentGoo `json:"goo,omitempty"`

	/* Immutable. Allows the patch job to run on Managed instance groups (MIGs). */
	// +optional
	MigInstancesAllowed *bool `json:"migInstancesAllowed,omitempty"`

	/* Immutable. The ExecStep to run after the patch update. */
	// +optional
	PostStep *PatchdeploymentPostStep `json:"postStep,omitempty"`

	/* Immutable. The ExecStep to run before the patch update. */
	// +optional
	PreStep *PatchdeploymentPreStep `json:"preStep,omitempty"`

	/* Immutable. Post-patch reboot settings. Possible values: ["DEFAULT", "ALWAYS", "NEVER"]. */
	// +optional
	RebootConfig *string `json:"rebootConfig,omitempty"`

	/* Immutable. Windows update settings. Use this setting to override the default Windows patch rules. */
	// +optional
	WindowsUpdate *PatchdeploymentWindowsUpdate `json:"windowsUpdate,omitempty"`

	/* Immutable. Yum update settings. Use this setting to override the default yum patch rules. */
	// +optional
	Yum *PatchdeploymentYum `json:"yum,omitempty"`

	/* Immutable. zypper update settings. Use this setting to override the default zypper patch rules. */
	// +optional
	Zypper *PatchdeploymentZypper `json:"zypper,omitempty"`
}

type PatchdeploymentPostStep struct {
	/* Immutable. The ExecStepConfig for all Linux VMs targeted by the PatchJob. */
	// +optional
	LinuxExecStepConfig *PatchdeploymentLinuxExecStepConfig `json:"linuxExecStepConfig,omitempty"`

	/* Immutable. The ExecStepConfig for all Windows VMs targeted by the PatchJob. */
	// +optional
	WindowsExecStepConfig *PatchdeploymentWindowsExecStepConfig `json:"windowsExecStepConfig,omitempty"`
}

type PatchdeploymentPreStep struct {
	/* Immutable. The ExecStepConfig for all Linux VMs targeted by the PatchJob. */
	// +optional
	LinuxExecStepConfig *PatchdeploymentLinuxExecStepConfig `json:"linuxExecStepConfig,omitempty"`

	/* Immutable. The ExecStepConfig for all Windows VMs targeted by the PatchJob. */
	// +optional
	WindowsExecStepConfig *PatchdeploymentWindowsExecStepConfig `json:"windowsExecStepConfig,omitempty"`
}

type PatchdeploymentRecurringSchedule struct {
	/* Immutable. The end time at which a recurring patch deployment schedule is no longer active.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	/* The time the last patch job ran successfully.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	LastExecuteTime *string `json:"lastExecuteTime,omitempty"`

	/* Immutable. Schedule with monthly executions. */
	// +optional
	Monthly *PatchdeploymentMonthly `json:"monthly,omitempty"`

	/* The time the next patch job is scheduled to run.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	NextExecuteTime *string `json:"nextExecuteTime,omitempty"`

	/* Immutable. The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	StartTime *string `json:"startTime,omitempty"`

	/* Immutable. Time of the day to run a recurring deployment. */
	TimeOfDay PatchdeploymentTimeOfDay `json:"timeOfDay"`

	/* Immutable. Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are
	determined by the chosen time zone. */
	TimeZone PatchdeploymentTimeZone `json:"timeZone"`

	/* Immutable. Schedule with weekly executions. */
	// +optional
	Weekly *PatchdeploymentWeekly `json:"weekly,omitempty"`
}

type PatchdeploymentRollout struct {
	/* Immutable. The maximum number (or percentage) of VMs per zone to disrupt at any given moment. The number of VMs calculated from multiplying the percentage by the total number of VMs in a zone is rounded up.
	During patching, a VM is considered disrupted from the time the agent is notified to begin until patching has completed. This disruption time includes the time to complete reboot and any post-patch steps.
	A VM contributes to the disruption budget if its patching operation fails either when applying the patches, running pre or post patch steps, or if it fails to respond with a success notification before timing out. VMs that are not running or do not have an active agent do not count toward this disruption budget.
	For zone-by-zone rollouts, if the disruption budget in a zone is exceeded, the patch job stops, because continuing to the next zone requires completion of the patch process in the previous zone.
	For example, if the disruption budget has a fixed value of 10, and 8 VMs fail to patch in the current zone, the patch job continues to patch 2 VMs at a time until the zone is completed. When that zone is completed successfully, patching begins with 10 VMs at a time in the next zone. If 10 VMs in the next zone fail to patch, the patch job stops. */
	DisruptionBudget PatchdeploymentDisruptionBudget `json:"disruptionBudget"`

	/* Immutable. Mode of the patch rollout. Possible values: ["ZONE_BY_ZONE", "CONCURRENT_ZONES"]. */
	Mode string `json:"mode"`
}

type PatchdeploymentTimeOfDay struct {
	/* Immutable. Hours of day in 24 hour format. Should be from 0 to 23.
	An API may choose to allow the value "24:00:00" for scenarios like business closing time. */
	// +optional
	Hours *int `json:"hours,omitempty"`

	/* Immutable. Minutes of hour of day. Must be from 0 to 59. */
	// +optional
	Minutes *int `json:"minutes,omitempty"`

	/* Immutable. Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Immutable. Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds. */
	// +optional
	Seconds *int `json:"seconds,omitempty"`
}

type PatchdeploymentTimeZone struct {
	/* Immutable. IANA Time Zone Database time zone, e.g. "America/New_York". */
	Id string `json:"id"`

	/* Immutable. IANA Time Zone Database version number, e.g. "2019a". */
	// +optional
	Version *string `json:"version,omitempty"`
}

type PatchdeploymentWeekDayOfMonth struct {
	/* Immutable. A day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]. */
	DayOfWeek string `json:"dayOfWeek"`

	/* Immutable. Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month. */
	WeekOrdinal int `json:"weekOrdinal"`
}

type PatchdeploymentWeekly struct {
	/* Immutable. IANA Time Zone Database time zone, e.g. "America/New_York". Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]. */
	DayOfWeek string `json:"dayOfWeek"`
}

type PatchdeploymentWindowsExecStepConfig struct {
	/* Immutable. Defaults to [0]. A list of possible return values that the execution can return to indicate a success. */
	// +optional
	AllowedSuccessCodes []int `json:"allowedSuccessCodes,omitempty"`

	/* Immutable. A Cloud Storage object containing the executable. */
	// +optional
	GcsObject *PatchdeploymentGcsObject `json:"gcsObject,omitempty"`

	/* Immutable. The script interpreter to use to run the script. If no interpreter is specified the script will
	be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]. */
	// +optional
	Interpreter *string `json:"interpreter,omitempty"`

	/* Immutable. An absolute path to the executable on the VM. */
	// +optional
	LocalPath *string `json:"localPath,omitempty"`
}

type PatchdeploymentWindowsUpdate struct {
	/* Immutable. Only apply updates of these windows update classifications. If empty, all updates are applied. Possible values: ["CRITICAL", "SECURITY", "DEFINITION", "DRIVER", "FEATURE_PACK", "SERVICE_PACK", "TOOL", "UPDATE_ROLLUP", "UPDATE"]. */
	// +optional
	Classifications []string `json:"classifications,omitempty"`

	/* Immutable. List of KBs to exclude from update. */
	// +optional
	Excludes []string `json:"excludes,omitempty"`

	/* Immutable. An exclusive list of kbs to be updated. These are the only patches that will be updated.
	This field must not be used with other patch configurations. */
	// +optional
	ExclusivePatches []string `json:"exclusivePatches,omitempty"`
}

type PatchdeploymentYum struct {
	/* Immutable. List of packages to exclude from update. These packages will be excluded. */
	// +optional
	Excludes []string `json:"excludes,omitempty"`

	/* Immutable. An exclusive list of packages to be updated. These are the only packages that will be updated.
	If these packages are not installed, they will be ignored. This field cannot be specified with
	any other patch configuration fields. */
	// +optional
	ExclusivePackages []string `json:"exclusivePackages,omitempty"`

	/* Immutable. Will cause patch to run yum update-minimal instead. */
	// +optional
	Minimal *bool `json:"minimal,omitempty"`

	/* Immutable. Adds the --security flag to yum update. Not supported on all platforms. */
	// +optional
	Security *bool `json:"security,omitempty"`
}

type PatchdeploymentZypper struct {
	/* Immutable. Install only patches with these categories. Common categories include security, recommended, and feature. */
	// +optional
	Categories []string `json:"categories,omitempty"`

	/* Immutable. List of packages to exclude from update. */
	// +optional
	Excludes []string `json:"excludes,omitempty"`

	/* Immutable. An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command.
	This field must not be used with any other patch configuration fields. */
	// +optional
	ExclusivePatches []string `json:"exclusivePatches,omitempty"`

	/* Immutable. Install only patches with these severities. Common severities include critical, important, moderate, and low. */
	// +optional
	Severities []string `json:"severities,omitempty"`

	/* Immutable. Adds the --with-optional flag to zypper patch. */
	// +optional
	WithOptional *bool `json:"withOptional,omitempty"`

	/* Immutable. Adds the --with-update flag, to zypper patch. */
	// +optional
	WithUpdate *bool `json:"withUpdate,omitempty"`
}

type OSConfigPatchDeploymentSpec struct {
	/* Immutable. Description of the patch deployment. Length of the description is limited to 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Duration of the patch. After the duration ends, the patch times out.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	Duration *string `json:"duration,omitempty"`

	/* Immutable. VM instances to patch. */
	InstanceFilter PatchdeploymentInstanceFilter `json:"instanceFilter"`

	/* Immutable. Schedule a one-time execution. */
	// +optional
	OneTimeSchedule *PatchdeploymentOneTimeSchedule `json:"oneTimeSchedule,omitempty"`

	/* Immutable. Patch configuration that is applied. */
	// +optional
	PatchConfig *PatchdeploymentPatchConfig `json:"patchConfig,omitempty"`

	/* Immutable. A name for the patch deployment in the project. When creating a name the following rules apply:
	* Must contain only lowercase letters, numbers, and hyphens.
	* Must start with a letter.
	* Must be between 1-63 characters.
	* Must end with a number or a letter.
	* Must be unique within the project. */
	PatchDeploymentId string `json:"patchDeploymentId"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Schedule recurring executions. */
	// +optional
	RecurringSchedule *PatchdeploymentRecurringSchedule `json:"recurringSchedule,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Rollout strategy of the patch job. */
	// +optional
	Rollout *PatchdeploymentRollout `json:"rollout,omitempty"`
}

type OSConfigPatchDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   OSConfigPatchDeployment's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time the patch deployment was created. Timestamp is in RFC3339 text format.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The last time a patch job was started by this deployment. Timestamp is in RFC3339 text format.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	LastExecuteTime *string `json:"lastExecuteTime,omitempty"`

	/* Unique name for the patch deployment resource in a project.
	The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
	A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcposconfigpatchdeployment;gcposconfigpatchdeployments
// +kubebuilder:subresource:status

// OSConfigPatchDeployment is the Schema for the osconfig API
// +k8s:openapi-gen=true
type OSConfigPatchDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSConfigPatchDeploymentSpec   `json:"spec,omitempty"`
	Status OSConfigPatchDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSConfigPatchDeploymentList contains a list of OSConfigPatchDeployment
type OSConfigPatchDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSConfigPatchDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSConfigPatchDeployment{}, &OSConfigPatchDeploymentList{})
}
