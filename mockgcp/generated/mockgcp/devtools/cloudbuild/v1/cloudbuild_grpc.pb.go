// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/devtools/cloudbuild/v1/cloudbuild.proto

package cloudbuildpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CloudBuildClient is the client API for CloudBuild service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudBuildClient interface {
	// Starts a build with the specified configuration.
	//
	// This method returns a long-running `Operation`, which includes the build
	// ID. Pass the build ID to `GetBuild` to determine the build status (such as
	// `SUCCESS` or `FAILURE`).
	CreateBuild(ctx context.Context, in *CreateBuildRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns information about a previously requested build.
	//
	// The `Build` that is returned includes its status (such as `SUCCESS`,
	// `FAILURE`, or `WORKING`), and timing information.
	GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error)
	// Lists previously requested builds.
	//
	// Previously requested builds may still be in-progress, or may have finished
	// successfully or unsuccessfully.
	ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error)
	// Cancels a build in progress.
	CancelBuild(ctx context.Context, in *CancelBuildRequest, opts ...grpc.CallOption) (*Build, error)
	// Creates a new build based on the specified build.
	//
	// This method creates a new build using the original build request, which may
	// or may not result in an identical build.
	//
	// For triggered builds:
	//
	// * Triggered builds resolve to a precise revision; therefore a retry of a
	// triggered build will result in a build that uses the same revision.
	//
	// For non-triggered builds that specify `RepoSource`:
	//
	// * If the original build built from the tip of a branch, the retried build
	// will build from the tip of that branch, which may not be the same revision
	// as the original build.
	// * If the original build specified a commit sha or revision ID, the retried
	// build will use the identical source.
	//
	// For builds that specify `StorageSource`:
	//
	// * If the original build pulled source from Cloud Storage without
	// specifying the generation of the object, the new build will use the current
	// object, which may be different from the original build source.
	// * If the original build pulled source from Cloud Storage and specified the
	// generation of the object, the new build will attempt to use the same
	// object, which may or may not be available depending on the bucket's
	// lifecycle management settings.
	RetryBuild(ctx context.Context, in *RetryBuildRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Approves or rejects a pending build.
	//
	// If approved, the returned LRO will be analogous to the LRO returned from
	// a CreateBuild call.
	//
	// If rejected, the returned LRO will be immediately done.
	ApproveBuild(ctx context.Context, in *ApproveBuildRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Creates a new `BuildTrigger`.
	//
	// This API is experimental.
	CreateBuildTrigger(ctx context.Context, in *CreateBuildTriggerRequest, opts ...grpc.CallOption) (*BuildTrigger, error)
	// Returns information about a `BuildTrigger`.
	//
	// This API is experimental.
	GetBuildTrigger(ctx context.Context, in *GetBuildTriggerRequest, opts ...grpc.CallOption) (*BuildTrigger, error)
	// Lists existing `BuildTrigger`s.
	//
	// This API is experimental.
	ListBuildTriggers(ctx context.Context, in *ListBuildTriggersRequest, opts ...grpc.CallOption) (*ListBuildTriggersResponse, error)
	// Deletes a `BuildTrigger` by its project ID and trigger ID.
	//
	// This API is experimental.
	DeleteBuildTrigger(ctx context.Context, in *DeleteBuildTriggerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates a `BuildTrigger` by its project ID and trigger ID.
	//
	// This API is experimental.
	UpdateBuildTrigger(ctx context.Context, in *UpdateBuildTriggerRequest, opts ...grpc.CallOption) (*BuildTrigger, error)
	// Runs a `BuildTrigger` at a particular source revision.
	//
	// To run a regional or global trigger, use the POST request
	// that includes the location endpoint in the path (ex.
	// v1/projects/{projectId}/locations/{region}/triggers/{triggerId}:run). The
	// POST request that does not include the location endpoint in the path can
	// only be used when running global triggers.
	RunBuildTrigger(ctx context.Context, in *RunBuildTriggerRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// ReceiveTriggerWebhook [Experimental] is called when the API receives a
	// webhook request targeted at a specific trigger.
	ReceiveTriggerWebhook(ctx context.Context, in *ReceiveTriggerWebhookRequest, opts ...grpc.CallOption) (*ReceiveTriggerWebhookResponse, error)
	// Creates a `WorkerPool`.
	CreateWorkerPool(ctx context.Context, in *CreateWorkerPoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns details of a `WorkerPool`.
	GetWorkerPool(ctx context.Context, in *GetWorkerPoolRequest, opts ...grpc.CallOption) (*WorkerPool, error)
	// Deletes a `WorkerPool`.
	DeleteWorkerPool(ctx context.Context, in *DeleteWorkerPoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates a `WorkerPool`.
	UpdateWorkerPool(ctx context.Context, in *UpdateWorkerPoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists `WorkerPool`s.
	ListWorkerPools(ctx context.Context, in *ListWorkerPoolsRequest, opts ...grpc.CallOption) (*ListWorkerPoolsResponse, error)
}

type cloudBuildClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudBuildClient(cc grpc.ClientConnInterface) CloudBuildClient {
	return &cloudBuildClient{cc}
}

func (c *cloudBuildClient) CreateBuild(ctx context.Context, in *CreateBuildRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CreateBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error) {
	out := new(ListBuildsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ListBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) CancelBuild(ctx context.Context, in *CancelBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CancelBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) RetryBuild(ctx context.Context, in *RetryBuildRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/RetryBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) ApproveBuild(ctx context.Context, in *ApproveBuildRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ApproveBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) CreateBuildTrigger(ctx context.Context, in *CreateBuildTriggerRequest, opts ...grpc.CallOption) (*BuildTrigger, error) {
	out := new(BuildTrigger)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CreateBuildTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) GetBuildTrigger(ctx context.Context, in *GetBuildTriggerRequest, opts ...grpc.CallOption) (*BuildTrigger, error) {
	out := new(BuildTrigger)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/GetBuildTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) ListBuildTriggers(ctx context.Context, in *ListBuildTriggersRequest, opts ...grpc.CallOption) (*ListBuildTriggersResponse, error) {
	out := new(ListBuildTriggersResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ListBuildTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) DeleteBuildTrigger(ctx context.Context, in *DeleteBuildTriggerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/DeleteBuildTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) UpdateBuildTrigger(ctx context.Context, in *UpdateBuildTriggerRequest, opts ...grpc.CallOption) (*BuildTrigger, error) {
	out := new(BuildTrigger)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/UpdateBuildTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) RunBuildTrigger(ctx context.Context, in *RunBuildTriggerRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/RunBuildTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) ReceiveTriggerWebhook(ctx context.Context, in *ReceiveTriggerWebhookRequest, opts ...grpc.CallOption) (*ReceiveTriggerWebhookResponse, error) {
	out := new(ReceiveTriggerWebhookResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ReceiveTriggerWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) CreateWorkerPool(ctx context.Context, in *CreateWorkerPoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CreateWorkerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) GetWorkerPool(ctx context.Context, in *GetWorkerPoolRequest, opts ...grpc.CallOption) (*WorkerPool, error) {
	out := new(WorkerPool)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/GetWorkerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) DeleteWorkerPool(ctx context.Context, in *DeleteWorkerPoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/DeleteWorkerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) UpdateWorkerPool(ctx context.Context, in *UpdateWorkerPoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/UpdateWorkerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudBuildClient) ListWorkerPools(ctx context.Context, in *ListWorkerPoolsRequest, opts ...grpc.CallOption) (*ListWorkerPoolsResponse, error) {
	out := new(ListWorkerPoolsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ListWorkerPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudBuildServer is the server API for CloudBuild service.
// All implementations must embed UnimplementedCloudBuildServer
// for forward compatibility
type CloudBuildServer interface {
	// Starts a build with the specified configuration.
	//
	// This method returns a long-running `Operation`, which includes the build
	// ID. Pass the build ID to `GetBuild` to determine the build status (such as
	// `SUCCESS` or `FAILURE`).
	CreateBuild(context.Context, *CreateBuildRequest) (*longrunningpb.Operation, error)
	// Returns information about a previously requested build.
	//
	// The `Build` that is returned includes its status (such as `SUCCESS`,
	// `FAILURE`, or `WORKING`), and timing information.
	GetBuild(context.Context, *GetBuildRequest) (*Build, error)
	// Lists previously requested builds.
	//
	// Previously requested builds may still be in-progress, or may have finished
	// successfully or unsuccessfully.
	ListBuilds(context.Context, *ListBuildsRequest) (*ListBuildsResponse, error)
	// Cancels a build in progress.
	CancelBuild(context.Context, *CancelBuildRequest) (*Build, error)
	// Creates a new build based on the specified build.
	//
	// This method creates a new build using the original build request, which may
	// or may not result in an identical build.
	//
	// For triggered builds:
	//
	// * Triggered builds resolve to a precise revision; therefore a retry of a
	// triggered build will result in a build that uses the same revision.
	//
	// For non-triggered builds that specify `RepoSource`:
	//
	// * If the original build built from the tip of a branch, the retried build
	// will build from the tip of that branch, which may not be the same revision
	// as the original build.
	// * If the original build specified a commit sha or revision ID, the retried
	// build will use the identical source.
	//
	// For builds that specify `StorageSource`:
	//
	// * If the original build pulled source from Cloud Storage without
	// specifying the generation of the object, the new build will use the current
	// object, which may be different from the original build source.
	// * If the original build pulled source from Cloud Storage and specified the
	// generation of the object, the new build will attempt to use the same
	// object, which may or may not be available depending on the bucket's
	// lifecycle management settings.
	RetryBuild(context.Context, *RetryBuildRequest) (*longrunningpb.Operation, error)
	// Approves or rejects a pending build.
	//
	// If approved, the returned LRO will be analogous to the LRO returned from
	// a CreateBuild call.
	//
	// If rejected, the returned LRO will be immediately done.
	ApproveBuild(context.Context, *ApproveBuildRequest) (*longrunningpb.Operation, error)
	// Creates a new `BuildTrigger`.
	//
	// This API is experimental.
	CreateBuildTrigger(context.Context, *CreateBuildTriggerRequest) (*BuildTrigger, error)
	// Returns information about a `BuildTrigger`.
	//
	// This API is experimental.
	GetBuildTrigger(context.Context, *GetBuildTriggerRequest) (*BuildTrigger, error)
	// Lists existing `BuildTrigger`s.
	//
	// This API is experimental.
	ListBuildTriggers(context.Context, *ListBuildTriggersRequest) (*ListBuildTriggersResponse, error)
	// Deletes a `BuildTrigger` by its project ID and trigger ID.
	//
	// This API is experimental.
	DeleteBuildTrigger(context.Context, *DeleteBuildTriggerRequest) (*empty.Empty, error)
	// Updates a `BuildTrigger` by its project ID and trigger ID.
	//
	// This API is experimental.
	UpdateBuildTrigger(context.Context, *UpdateBuildTriggerRequest) (*BuildTrigger, error)
	// Runs a `BuildTrigger` at a particular source revision.
	//
	// To run a regional or global trigger, use the POST request
	// that includes the location endpoint in the path (ex.
	// v1/projects/{projectId}/locations/{region}/triggers/{triggerId}:run). The
	// POST request that does not include the location endpoint in the path can
	// only be used when running global triggers.
	RunBuildTrigger(context.Context, *RunBuildTriggerRequest) (*longrunningpb.Operation, error)
	// ReceiveTriggerWebhook [Experimental] is called when the API receives a
	// webhook request targeted at a specific trigger.
	ReceiveTriggerWebhook(context.Context, *ReceiveTriggerWebhookRequest) (*ReceiveTriggerWebhookResponse, error)
	// Creates a `WorkerPool`.
	CreateWorkerPool(context.Context, *CreateWorkerPoolRequest) (*longrunningpb.Operation, error)
	// Returns details of a `WorkerPool`.
	GetWorkerPool(context.Context, *GetWorkerPoolRequest) (*WorkerPool, error)
	// Deletes a `WorkerPool`.
	DeleteWorkerPool(context.Context, *DeleteWorkerPoolRequest) (*longrunningpb.Operation, error)
	// Updates a `WorkerPool`.
	UpdateWorkerPool(context.Context, *UpdateWorkerPoolRequest) (*longrunningpb.Operation, error)
	// Lists `WorkerPool`s.
	ListWorkerPools(context.Context, *ListWorkerPoolsRequest) (*ListWorkerPoolsResponse, error)
	mustEmbedUnimplementedCloudBuildServer()
}

// UnimplementedCloudBuildServer must be embedded to have forward compatible implementations.
type UnimplementedCloudBuildServer struct {
}

func (UnimplementedCloudBuildServer) CreateBuild(context.Context, *CreateBuildRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBuild not implemented")
}
func (UnimplementedCloudBuildServer) GetBuild(context.Context, *GetBuildRequest) (*Build, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuild not implemented")
}
func (UnimplementedCloudBuildServer) ListBuilds(context.Context, *ListBuildsRequest) (*ListBuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuilds not implemented")
}
func (UnimplementedCloudBuildServer) CancelBuild(context.Context, *CancelBuildRequest) (*Build, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBuild not implemented")
}
func (UnimplementedCloudBuildServer) RetryBuild(context.Context, *RetryBuildRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryBuild not implemented")
}
func (UnimplementedCloudBuildServer) ApproveBuild(context.Context, *ApproveBuildRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveBuild not implemented")
}
func (UnimplementedCloudBuildServer) CreateBuildTrigger(context.Context, *CreateBuildTriggerRequest) (*BuildTrigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBuildTrigger not implemented")
}
func (UnimplementedCloudBuildServer) GetBuildTrigger(context.Context, *GetBuildTriggerRequest) (*BuildTrigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildTrigger not implemented")
}
func (UnimplementedCloudBuildServer) ListBuildTriggers(context.Context, *ListBuildTriggersRequest) (*ListBuildTriggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuildTriggers not implemented")
}
func (UnimplementedCloudBuildServer) DeleteBuildTrigger(context.Context, *DeleteBuildTriggerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBuildTrigger not implemented")
}
func (UnimplementedCloudBuildServer) UpdateBuildTrigger(context.Context, *UpdateBuildTriggerRequest) (*BuildTrigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBuildTrigger not implemented")
}
func (UnimplementedCloudBuildServer) RunBuildTrigger(context.Context, *RunBuildTriggerRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunBuildTrigger not implemented")
}
func (UnimplementedCloudBuildServer) ReceiveTriggerWebhook(context.Context, *ReceiveTriggerWebhookRequest) (*ReceiveTriggerWebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveTriggerWebhook not implemented")
}
func (UnimplementedCloudBuildServer) CreateWorkerPool(context.Context, *CreateWorkerPoolRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkerPool not implemented")
}
func (UnimplementedCloudBuildServer) GetWorkerPool(context.Context, *GetWorkerPoolRequest) (*WorkerPool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerPool not implemented")
}
func (UnimplementedCloudBuildServer) DeleteWorkerPool(context.Context, *DeleteWorkerPoolRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkerPool not implemented")
}
func (UnimplementedCloudBuildServer) UpdateWorkerPool(context.Context, *UpdateWorkerPoolRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkerPool not implemented")
}
func (UnimplementedCloudBuildServer) ListWorkerPools(context.Context, *ListWorkerPoolsRequest) (*ListWorkerPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkerPools not implemented")
}
func (UnimplementedCloudBuildServer) mustEmbedUnimplementedCloudBuildServer() {}

// UnsafeCloudBuildServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudBuildServer will
// result in compilation errors.
type UnsafeCloudBuildServer interface {
	mustEmbedUnimplementedCloudBuildServer()
}

func RegisterCloudBuildServer(s grpc.ServiceRegistrar, srv CloudBuildServer) {
	s.RegisterService(&CloudBuild_ServiceDesc, srv)
}

func _CloudBuild_CreateBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).CreateBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CreateBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).CreateBuild(ctx, req.(*CreateBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_GetBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).GetBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/GetBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).GetBuild(ctx, req.(*GetBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_ListBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).ListBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ListBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).ListBuilds(ctx, req.(*ListBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_CancelBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).CancelBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CancelBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).CancelBuild(ctx, req.(*CancelBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_RetryBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).RetryBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/RetryBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).RetryBuild(ctx, req.(*RetryBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_ApproveBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).ApproveBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ApproveBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).ApproveBuild(ctx, req.(*ApproveBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_CreateBuildTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBuildTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).CreateBuildTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CreateBuildTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).CreateBuildTrigger(ctx, req.(*CreateBuildTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_GetBuildTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).GetBuildTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/GetBuildTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).GetBuildTrigger(ctx, req.(*GetBuildTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_ListBuildTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).ListBuildTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ListBuildTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).ListBuildTriggers(ctx, req.(*ListBuildTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_DeleteBuildTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBuildTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).DeleteBuildTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/DeleteBuildTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).DeleteBuildTrigger(ctx, req.(*DeleteBuildTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_UpdateBuildTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBuildTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).UpdateBuildTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/UpdateBuildTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).UpdateBuildTrigger(ctx, req.(*UpdateBuildTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_RunBuildTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunBuildTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).RunBuildTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/RunBuildTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).RunBuildTrigger(ctx, req.(*RunBuildTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_ReceiveTriggerWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveTriggerWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).ReceiveTriggerWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ReceiveTriggerWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).ReceiveTriggerWebhook(ctx, req.(*ReceiveTriggerWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_CreateWorkerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).CreateWorkerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/CreateWorkerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).CreateWorkerPool(ctx, req.(*CreateWorkerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_GetWorkerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).GetWorkerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/GetWorkerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).GetWorkerPool(ctx, req.(*GetWorkerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_DeleteWorkerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).DeleteWorkerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/DeleteWorkerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).DeleteWorkerPool(ctx, req.(*DeleteWorkerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_UpdateWorkerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).UpdateWorkerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/UpdateWorkerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).UpdateWorkerPool(ctx, req.(*UpdateWorkerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudBuild_ListWorkerPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkerPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudBuildServer).ListWorkerPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.devtools.cloudbuild.v1.CloudBuild/ListWorkerPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudBuildServer).ListWorkerPools(ctx, req.(*ListWorkerPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudBuild_ServiceDesc is the grpc.ServiceDesc for CloudBuild service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudBuild_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.devtools.cloudbuild.v1.CloudBuild",
	HandlerType: (*CloudBuildServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBuild",
			Handler:    _CloudBuild_CreateBuild_Handler,
		},
		{
			MethodName: "GetBuild",
			Handler:    _CloudBuild_GetBuild_Handler,
		},
		{
			MethodName: "ListBuilds",
			Handler:    _CloudBuild_ListBuilds_Handler,
		},
		{
			MethodName: "CancelBuild",
			Handler:    _CloudBuild_CancelBuild_Handler,
		},
		{
			MethodName: "RetryBuild",
			Handler:    _CloudBuild_RetryBuild_Handler,
		},
		{
			MethodName: "ApproveBuild",
			Handler:    _CloudBuild_ApproveBuild_Handler,
		},
		{
			MethodName: "CreateBuildTrigger",
			Handler:    _CloudBuild_CreateBuildTrigger_Handler,
		},
		{
			MethodName: "GetBuildTrigger",
			Handler:    _CloudBuild_GetBuildTrigger_Handler,
		},
		{
			MethodName: "ListBuildTriggers",
			Handler:    _CloudBuild_ListBuildTriggers_Handler,
		},
		{
			MethodName: "DeleteBuildTrigger",
			Handler:    _CloudBuild_DeleteBuildTrigger_Handler,
		},
		{
			MethodName: "UpdateBuildTrigger",
			Handler:    _CloudBuild_UpdateBuildTrigger_Handler,
		},
		{
			MethodName: "RunBuildTrigger",
			Handler:    _CloudBuild_RunBuildTrigger_Handler,
		},
		{
			MethodName: "ReceiveTriggerWebhook",
			Handler:    _CloudBuild_ReceiveTriggerWebhook_Handler,
		},
		{
			MethodName: "CreateWorkerPool",
			Handler:    _CloudBuild_CreateWorkerPool_Handler,
		},
		{
			MethodName: "GetWorkerPool",
			Handler:    _CloudBuild_GetWorkerPool_Handler,
		},
		{
			MethodName: "DeleteWorkerPool",
			Handler:    _CloudBuild_DeleteWorkerPool_Handler,
		},
		{
			MethodName: "UpdateWorkerPool",
			Handler:    _CloudBuild_UpdateWorkerPool_Handler,
		},
		{
			MethodName: "ListWorkerPools",
			Handler:    _CloudBuild_ListWorkerPools_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/devtools/cloudbuild/v1/cloudbuild.proto",
}
