package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/privatestate"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Set is a schema validator for types.Set attributes.
type Set interface {
	Describer

	// PlanModifySet should perform the modification.
	PlanModifySet(context.Context, SetRequest, *SetResponse)
}

// SetRequest is a request for types.Set schema plan modification.
type SetRequest struct {
	// Path contains the path of the attribute for modification. Use this path
	// for any response diagnostics.
	Path path.Path

	// PathExpression contains the expression matching the exact path
	// of the attribute for modification.
	PathExpression path.Expression

	// Config contains the entire configuration of the resource.
	Config tfsdk.Config

	// ConfigValue contains the value of the attribute for modification from the configuration.
	ConfigValue types.Set

	// Plan contains the entire proposed new state of the resource.
	Plan tfsdk.Plan

	// PlanValue contains the value of the attribute for modification from the proposed new state.
	PlanValue types.Set

	// State contains the entire prior state of the resource.
	State tfsdk.State

	// StateValue contains the value of the attribute for modification from the prior state.
	StateValue types.Set

	// Private is provider-defined resource private state data which was previously
	// stored with the resource state. This data is opaque to Terraform and does
	// not affect plan output. Any existing data is copied to
	// SetResponse.Private to prevent accidental private state data loss.
	//
	// The private state data is always the original data when the schema-based plan
	// modification began or, is updated as the logic traverses deeper into underlying
	// attributes.
	//
	// Use the GetKey method to read data. Use the SetKey method on
	// SetResponse.Private to update or remove a value.
	Private *privatestate.ProviderData
}

// SetResponse is a response to a SetRequest.
type SetResponse struct {
	// PlanValue is the planned new state for the attribute.
	PlanValue types.Set

	// RequiresReplace indicates whether a change in the attribute
	// requires replacement of the whole resource.
	RequiresReplace bool

	// Private is the private state resource data following the PlanModifySet operation.
	// This field is pre-populated from SetRequest.Private and
	// can be modified during the resource's PlanModifySet operation.
	//
	// The private state data is always the original data when the schema-based plan
	// modification began or, is updated as the logic traverses deeper into underlying
	// attributes.
	Private *privatestate.ProviderData

	// Diagnostics report errors or warnings related to validating the data
	// source configuration. An empty slice indicates success, with no warnings
	// or errors generated.
	Diagnostics diag.Diagnostics
}
