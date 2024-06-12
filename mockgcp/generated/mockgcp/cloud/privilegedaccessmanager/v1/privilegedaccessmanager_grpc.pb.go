// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/privilegedaccessmanager/v1/privilegedaccessmanager.proto

package privilegedaccessmanagerpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PrivilegedAccessManagerClient is the client API for PrivilegedAccessManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivilegedAccessManagerClient interface {
	// CheckOnboardingStatus reports the onboarding status for a
	// project/folder/organization. Any findings reported by this API need to be
	// fixed before PAM can be used on the resource.
	CheckOnboardingStatus(ctx context.Context, in *CheckOnboardingStatusRequest, opts ...grpc.CallOption) (*CheckOnboardingStatusResponse, error)
	// Lists entitlements in a given project/folder/organization and location.
	ListEntitlements(ctx context.Context, in *ListEntitlementsRequest, opts ...grpc.CallOption) (*ListEntitlementsResponse, error)
	// `SearchEntitlements` returns entitlements on which the caller has the
	// specified access.
	SearchEntitlements(ctx context.Context, in *SearchEntitlementsRequest, opts ...grpc.CallOption) (*SearchEntitlementsResponse, error)
	// Gets details of a single entitlement.
	GetEntitlement(ctx context.Context, in *GetEntitlementRequest, opts ...grpc.CallOption) (*Entitlement, error)
	// Creates a new entitlement in a given project/folder/organization and
	// location.
	CreateEntitlement(ctx context.Context, in *CreateEntitlementRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a single entitlement. This method can only be called when there
	// are no in-progress (ACTIVE/ACTIVATING/REVOKING) grants under the
	// entitlement.
	DeleteEntitlement(ctx context.Context, in *DeleteEntitlementRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates the entitlement specified in the request. Updated fields in the
	// entitlement need to be specified in an update mask. The changes made to an
	// entitlement are applicable only on future grants of the entitlement.
	// However, if new approvers are added or existing approvers are removed from
	// the approval workflow, the changes are effective on existing grants.
	//
	// The following fields are not supported for updates:
	//
	//   - All immutable fields
	//   - Entitlement name
	//   - Resource name
	//   - Resource type
	//   - Adding an approval workflow in an entitlement which previously had no
	//     approval workflow.
	//   - Deleting the approval workflow from an entitlement.
	//   - Adding or deleting a step in the approval workflow (only one step is
	//     supported)
	//
	// Note that updates are allowed on the list of approvers in an approval
	// workflow step.
	UpdateEntitlement(ctx context.Context, in *UpdateEntitlementRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists grants for a given entitlement.
	ListGrants(ctx context.Context, in *ListGrantsRequest, opts ...grpc.CallOption) (*ListGrantsResponse, error)
	// `SearchGrants` returns grants that are related to the calling user in the
	// specified way.
	SearchGrants(ctx context.Context, in *SearchGrantsRequest, opts ...grpc.CallOption) (*SearchGrantsResponse, error)
	// Get details of a single grant.
	GetGrant(ctx context.Context, in *GetGrantRequest, opts ...grpc.CallOption) (*Grant, error)
	// Creates a new grant in a given project and location.
	CreateGrant(ctx context.Context, in *CreateGrantRequest, opts ...grpc.CallOption) (*Grant, error)
	// `ApproveGrant` is used to approve a grant. This method can only be called
	// on a grant when it's in the `APPROVAL_AWAITED` state. This operation can't
	// be undone.
	ApproveGrant(ctx context.Context, in *ApproveGrantRequest, opts ...grpc.CallOption) (*Grant, error)
	// `DenyGrant` is used to deny a grant. This method can only be called on a
	// grant when it's in the `APPROVAL_AWAITED` state. This operation can't be
	// undone.
	DenyGrant(ctx context.Context, in *DenyGrantRequest, opts ...grpc.CallOption) (*Grant, error)
	// `RevokeGrant` is used to immediately revoke access for a grant. This method
	// can be called when the grant is in a non-terminal state.
	RevokeGrant(ctx context.Context, in *RevokeGrantRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type privilegedAccessManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivilegedAccessManagerClient(cc grpc.ClientConnInterface) PrivilegedAccessManagerClient {
	return &privilegedAccessManagerClient{cc}
}

func (c *privilegedAccessManagerClient) CheckOnboardingStatus(ctx context.Context, in *CheckOnboardingStatusRequest, opts ...grpc.CallOption) (*CheckOnboardingStatusResponse, error) {
	out := new(CheckOnboardingStatusResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/CheckOnboardingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) ListEntitlements(ctx context.Context, in *ListEntitlementsRequest, opts ...grpc.CallOption) (*ListEntitlementsResponse, error) {
	out := new(ListEntitlementsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/ListEntitlements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) SearchEntitlements(ctx context.Context, in *SearchEntitlementsRequest, opts ...grpc.CallOption) (*SearchEntitlementsResponse, error) {
	out := new(SearchEntitlementsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/SearchEntitlements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) GetEntitlement(ctx context.Context, in *GetEntitlementRequest, opts ...grpc.CallOption) (*Entitlement, error) {
	out := new(Entitlement)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/GetEntitlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) CreateEntitlement(ctx context.Context, in *CreateEntitlementRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/CreateEntitlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) DeleteEntitlement(ctx context.Context, in *DeleteEntitlementRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/DeleteEntitlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) UpdateEntitlement(ctx context.Context, in *UpdateEntitlementRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/UpdateEntitlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) ListGrants(ctx context.Context, in *ListGrantsRequest, opts ...grpc.CallOption) (*ListGrantsResponse, error) {
	out := new(ListGrantsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/ListGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) SearchGrants(ctx context.Context, in *SearchGrantsRequest, opts ...grpc.CallOption) (*SearchGrantsResponse, error) {
	out := new(SearchGrantsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/SearchGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) GetGrant(ctx context.Context, in *GetGrantRequest, opts ...grpc.CallOption) (*Grant, error) {
	out := new(Grant)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/GetGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) CreateGrant(ctx context.Context, in *CreateGrantRequest, opts ...grpc.CallOption) (*Grant, error) {
	out := new(Grant)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/CreateGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) ApproveGrant(ctx context.Context, in *ApproveGrantRequest, opts ...grpc.CallOption) (*Grant, error) {
	out := new(Grant)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/ApproveGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) DenyGrant(ctx context.Context, in *DenyGrantRequest, opts ...grpc.CallOption) (*Grant, error) {
	out := new(Grant)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/DenyGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privilegedAccessManagerClient) RevokeGrant(ctx context.Context, in *RevokeGrantRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/RevokeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivilegedAccessManagerServer is the server API for PrivilegedAccessManager service.
// All implementations must embed UnimplementedPrivilegedAccessManagerServer
// for forward compatibility
type PrivilegedAccessManagerServer interface {
	// CheckOnboardingStatus reports the onboarding status for a
	// project/folder/organization. Any findings reported by this API need to be
	// fixed before PAM can be used on the resource.
	CheckOnboardingStatus(context.Context, *CheckOnboardingStatusRequest) (*CheckOnboardingStatusResponse, error)
	// Lists entitlements in a given project/folder/organization and location.
	ListEntitlements(context.Context, *ListEntitlementsRequest) (*ListEntitlementsResponse, error)
	// `SearchEntitlements` returns entitlements on which the caller has the
	// specified access.
	SearchEntitlements(context.Context, *SearchEntitlementsRequest) (*SearchEntitlementsResponse, error)
	// Gets details of a single entitlement.
	GetEntitlement(context.Context, *GetEntitlementRequest) (*Entitlement, error)
	// Creates a new entitlement in a given project/folder/organization and
	// location.
	CreateEntitlement(context.Context, *CreateEntitlementRequest) (*longrunningpb.Operation, error)
	// Deletes a single entitlement. This method can only be called when there
	// are no in-progress (ACTIVE/ACTIVATING/REVOKING) grants under the
	// entitlement.
	DeleteEntitlement(context.Context, *DeleteEntitlementRequest) (*longrunningpb.Operation, error)
	// Updates the entitlement specified in the request. Updated fields in the
	// entitlement need to be specified in an update mask. The changes made to an
	// entitlement are applicable only on future grants of the entitlement.
	// However, if new approvers are added or existing approvers are removed from
	// the approval workflow, the changes are effective on existing grants.
	//
	// The following fields are not supported for updates:
	//
	//   - All immutable fields
	//   - Entitlement name
	//   - Resource name
	//   - Resource type
	//   - Adding an approval workflow in an entitlement which previously had no
	//     approval workflow.
	//   - Deleting the approval workflow from an entitlement.
	//   - Adding or deleting a step in the approval workflow (only one step is
	//     supported)
	//
	// Note that updates are allowed on the list of approvers in an approval
	// workflow step.
	UpdateEntitlement(context.Context, *UpdateEntitlementRequest) (*longrunningpb.Operation, error)
	// Lists grants for a given entitlement.
	ListGrants(context.Context, *ListGrantsRequest) (*ListGrantsResponse, error)
	// `SearchGrants` returns grants that are related to the calling user in the
	// specified way.
	SearchGrants(context.Context, *SearchGrantsRequest) (*SearchGrantsResponse, error)
	// Get details of a single grant.
	GetGrant(context.Context, *GetGrantRequest) (*Grant, error)
	// Creates a new grant in a given project and location.
	CreateGrant(context.Context, *CreateGrantRequest) (*Grant, error)
	// `ApproveGrant` is used to approve a grant. This method can only be called
	// on a grant when it's in the `APPROVAL_AWAITED` state. This operation can't
	// be undone.
	ApproveGrant(context.Context, *ApproveGrantRequest) (*Grant, error)
	// `DenyGrant` is used to deny a grant. This method can only be called on a
	// grant when it's in the `APPROVAL_AWAITED` state. This operation can't be
	// undone.
	DenyGrant(context.Context, *DenyGrantRequest) (*Grant, error)
	// `RevokeGrant` is used to immediately revoke access for a grant. This method
	// can be called when the grant is in a non-terminal state.
	RevokeGrant(context.Context, *RevokeGrantRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedPrivilegedAccessManagerServer()
}

// UnimplementedPrivilegedAccessManagerServer must be embedded to have forward compatible implementations.
type UnimplementedPrivilegedAccessManagerServer struct {
}

func (UnimplementedPrivilegedAccessManagerServer) CheckOnboardingStatus(context.Context, *CheckOnboardingStatusRequest) (*CheckOnboardingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOnboardingStatus not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) ListEntitlements(context.Context, *ListEntitlementsRequest) (*ListEntitlementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntitlements not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) SearchEntitlements(context.Context, *SearchEntitlementsRequest) (*SearchEntitlementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEntitlements not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) GetEntitlement(context.Context, *GetEntitlementRequest) (*Entitlement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntitlement not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) CreateEntitlement(context.Context, *CreateEntitlementRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntitlement not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) DeleteEntitlement(context.Context, *DeleteEntitlementRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntitlement not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) UpdateEntitlement(context.Context, *UpdateEntitlementRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntitlement not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) ListGrants(context.Context, *ListGrantsRequest) (*ListGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGrants not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) SearchGrants(context.Context, *SearchGrantsRequest) (*SearchGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGrants not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) GetGrant(context.Context, *GetGrantRequest) (*Grant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGrant not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) CreateGrant(context.Context, *CreateGrantRequest) (*Grant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGrant not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) ApproveGrant(context.Context, *ApproveGrantRequest) (*Grant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveGrant not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) DenyGrant(context.Context, *DenyGrantRequest) (*Grant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyGrant not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) RevokeGrant(context.Context, *RevokeGrantRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeGrant not implemented")
}
func (UnimplementedPrivilegedAccessManagerServer) mustEmbedUnimplementedPrivilegedAccessManagerServer() {
}

// UnsafePrivilegedAccessManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivilegedAccessManagerServer will
// result in compilation errors.
type UnsafePrivilegedAccessManagerServer interface {
	mustEmbedUnimplementedPrivilegedAccessManagerServer()
}

func RegisterPrivilegedAccessManagerServer(s grpc.ServiceRegistrar, srv PrivilegedAccessManagerServer) {
	s.RegisterService(&PrivilegedAccessManager_ServiceDesc, srv)
}

func _PrivilegedAccessManager_CheckOnboardingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOnboardingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).CheckOnboardingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/CheckOnboardingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).CheckOnboardingStatus(ctx, req.(*CheckOnboardingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_ListEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntitlementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).ListEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/ListEntitlements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).ListEntitlements(ctx, req.(*ListEntitlementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_SearchEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchEntitlementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).SearchEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/SearchEntitlements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).SearchEntitlements(ctx, req.(*SearchEntitlementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_GetEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntitlementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).GetEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/GetEntitlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).GetEntitlement(ctx, req.(*GetEntitlementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_CreateEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntitlementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).CreateEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/CreateEntitlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).CreateEntitlement(ctx, req.(*CreateEntitlementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_DeleteEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntitlementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).DeleteEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/DeleteEntitlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).DeleteEntitlement(ctx, req.(*DeleteEntitlementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_UpdateEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntitlementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).UpdateEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/UpdateEntitlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).UpdateEntitlement(ctx, req.(*UpdateEntitlementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_ListGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).ListGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/ListGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).ListGrants(ctx, req.(*ListGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_SearchGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).SearchGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/SearchGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).SearchGrants(ctx, req.(*SearchGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_GetGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).GetGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/GetGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).GetGrant(ctx, req.(*GetGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_CreateGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).CreateGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/CreateGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).CreateGrant(ctx, req.(*CreateGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_ApproveGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).ApproveGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/ApproveGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).ApproveGrant(ctx, req.(*ApproveGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_DenyGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).DenyGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/DenyGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).DenyGrant(ctx, req.(*DenyGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivilegedAccessManager_RevokeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegedAccessManagerServer).RevokeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager/RevokeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegedAccessManagerServer).RevokeGrant(ctx, req.(*RevokeGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivilegedAccessManager_ServiceDesc is the grpc.ServiceDesc for PrivilegedAccessManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivilegedAccessManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.privilegedaccessmanager.v1.PrivilegedAccessManager",
	HandlerType: (*PrivilegedAccessManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckOnboardingStatus",
			Handler:    _PrivilegedAccessManager_CheckOnboardingStatus_Handler,
		},
		{
			MethodName: "ListEntitlements",
			Handler:    _PrivilegedAccessManager_ListEntitlements_Handler,
		},
		{
			MethodName: "SearchEntitlements",
			Handler:    _PrivilegedAccessManager_SearchEntitlements_Handler,
		},
		{
			MethodName: "GetEntitlement",
			Handler:    _PrivilegedAccessManager_GetEntitlement_Handler,
		},
		{
			MethodName: "CreateEntitlement",
			Handler:    _PrivilegedAccessManager_CreateEntitlement_Handler,
		},
		{
			MethodName: "DeleteEntitlement",
			Handler:    _PrivilegedAccessManager_DeleteEntitlement_Handler,
		},
		{
			MethodName: "UpdateEntitlement",
			Handler:    _PrivilegedAccessManager_UpdateEntitlement_Handler,
		},
		{
			MethodName: "ListGrants",
			Handler:    _PrivilegedAccessManager_ListGrants_Handler,
		},
		{
			MethodName: "SearchGrants",
			Handler:    _PrivilegedAccessManager_SearchGrants_Handler,
		},
		{
			MethodName: "GetGrant",
			Handler:    _PrivilegedAccessManager_GetGrant_Handler,
		},
		{
			MethodName: "CreateGrant",
			Handler:    _PrivilegedAccessManager_CreateGrant_Handler,
		},
		{
			MethodName: "ApproveGrant",
			Handler:    _PrivilegedAccessManager_ApproveGrant_Handler,
		},
		{
			MethodName: "DenyGrant",
			Handler:    _PrivilegedAccessManager_DenyGrant_Handler,
		},
		{
			MethodName: "RevokeGrant",
			Handler:    _PrivilegedAccessManager_RevokeGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/privilegedaccessmanager/v1/privilegedaccessmanager.proto",
}
