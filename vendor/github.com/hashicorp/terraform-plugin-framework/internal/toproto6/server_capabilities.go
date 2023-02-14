package toproto6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ServerCapabilities returns the *tfprotov6.ServerCapabilities for a
// *fwserver.ServerCapabilities.
func ServerCapabilities(ctx context.Context, fw *fwserver.ServerCapabilities) *tfprotov6.ServerCapabilities {
	if fw == nil {
		return nil
	}

	return &tfprotov6.ServerCapabilities{
		PlanDestroy: fw.PlanDestroy,
	}
}
