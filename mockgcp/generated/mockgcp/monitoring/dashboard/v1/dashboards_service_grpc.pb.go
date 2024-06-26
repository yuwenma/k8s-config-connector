// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/monitoring/dashboard/v1/dashboards_service.proto

package dashboardpb

import (
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

// DashboardsServiceClient is the client API for DashboardsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardsServiceClient interface {
	// Creates a new custom dashboard. For examples on how you can use this API to
	// create dashboards, see [Managing dashboards by
	// API](https://cloud.google.com/monitoring/dashboards/api-dashboard). This
	// method requires the `monitoring.dashboards.create` permission on the
	// specified project. For more information about permissions, see [Cloud
	// Identity and Access Management](https://cloud.google.com/iam).
	CreateDashboard(ctx context.Context, in *CreateDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error)
	// Lists the existing dashboards.
	//
	// This method requires the `monitoring.dashboards.list` permission
	// on the specified project. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	ListDashboards(ctx context.Context, in *ListDashboardsRequest, opts ...grpc.CallOption) (*ListDashboardsResponse, error)
	// Fetches a specific dashboard.
	//
	// This method requires the `monitoring.dashboards.get` permission
	// on the specified dashboard. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error)
	// Deletes an existing custom dashboard.
	//
	// This method requires the `monitoring.dashboards.delete` permission
	// on the specified dashboard. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	DeleteDashboard(ctx context.Context, in *DeleteDashboardRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Replaces an existing custom dashboard with a new definition.
	//
	// This method requires the `monitoring.dashboards.update` permission
	// on the specified dashboard. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	UpdateDashboard(ctx context.Context, in *UpdateDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error)
}

type dashboardsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardsServiceClient(cc grpc.ClientConnInterface) DashboardsServiceClient {
	return &dashboardsServiceClient{cc}
}

func (c *dashboardsServiceClient) CreateDashboard(ctx context.Context, in *CreateDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error) {
	out := new(Dashboard)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.dashboard.v1.DashboardsService/CreateDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) ListDashboards(ctx context.Context, in *ListDashboardsRequest, opts ...grpc.CallOption) (*ListDashboardsResponse, error) {
	out := new(ListDashboardsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.dashboard.v1.DashboardsService/ListDashboards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error) {
	out := new(Dashboard)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.dashboard.v1.DashboardsService/GetDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) DeleteDashboard(ctx context.Context, in *DeleteDashboardRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.dashboard.v1.DashboardsService/DeleteDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) UpdateDashboard(ctx context.Context, in *UpdateDashboardRequest, opts ...grpc.CallOption) (*Dashboard, error) {
	out := new(Dashboard)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.dashboard.v1.DashboardsService/UpdateDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardsServiceServer is the server API for DashboardsService service.
// All implementations must embed UnimplementedDashboardsServiceServer
// for forward compatibility
type DashboardsServiceServer interface {
	// Creates a new custom dashboard. For examples on how you can use this API to
	// create dashboards, see [Managing dashboards by
	// API](https://cloud.google.com/monitoring/dashboards/api-dashboard). This
	// method requires the `monitoring.dashboards.create` permission on the
	// specified project. For more information about permissions, see [Cloud
	// Identity and Access Management](https://cloud.google.com/iam).
	CreateDashboard(context.Context, *CreateDashboardRequest) (*Dashboard, error)
	// Lists the existing dashboards.
	//
	// This method requires the `monitoring.dashboards.list` permission
	// on the specified project. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	ListDashboards(context.Context, *ListDashboardsRequest) (*ListDashboardsResponse, error)
	// Fetches a specific dashboard.
	//
	// This method requires the `monitoring.dashboards.get` permission
	// on the specified dashboard. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	GetDashboard(context.Context, *GetDashboardRequest) (*Dashboard, error)
	// Deletes an existing custom dashboard.
	//
	// This method requires the `monitoring.dashboards.delete` permission
	// on the specified dashboard. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	DeleteDashboard(context.Context, *DeleteDashboardRequest) (*empty.Empty, error)
	// Replaces an existing custom dashboard with a new definition.
	//
	// This method requires the `monitoring.dashboards.update` permission
	// on the specified dashboard. For more information, see
	// [Cloud Identity and Access Management](https://cloud.google.com/iam).
	UpdateDashboard(context.Context, *UpdateDashboardRequest) (*Dashboard, error)
	mustEmbedUnimplementedDashboardsServiceServer()
}

// UnimplementedDashboardsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardsServiceServer struct {
}

func (UnimplementedDashboardsServiceServer) CreateDashboard(context.Context, *CreateDashboardRequest) (*Dashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) ListDashboards(context.Context, *ListDashboardsRequest) (*ListDashboardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDashboards not implemented")
}
func (UnimplementedDashboardsServiceServer) GetDashboard(context.Context, *GetDashboardRequest) (*Dashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) DeleteDashboard(context.Context, *DeleteDashboardRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) UpdateDashboard(context.Context, *UpdateDashboardRequest) (*Dashboard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) mustEmbedUnimplementedDashboardsServiceServer() {}

// UnsafeDashboardsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardsServiceServer will
// result in compilation errors.
type UnsafeDashboardsServiceServer interface {
	mustEmbedUnimplementedDashboardsServiceServer()
}

func RegisterDashboardsServiceServer(s grpc.ServiceRegistrar, srv DashboardsServiceServer) {
	s.RegisterService(&DashboardsService_ServiceDesc, srv)
}

func _DashboardsService_CreateDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).CreateDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.dashboard.v1.DashboardsService/CreateDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).CreateDashboard(ctx, req.(*CreateDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_ListDashboards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDashboardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).ListDashboards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.dashboard.v1.DashboardsService/ListDashboards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).ListDashboards(ctx, req.(*ListDashboardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.dashboard.v1.DashboardsService/GetDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).GetDashboard(ctx, req.(*GetDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_DeleteDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).DeleteDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.dashboard.v1.DashboardsService/DeleteDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).DeleteDashboard(ctx, req.(*DeleteDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_UpdateDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).UpdateDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.dashboard.v1.DashboardsService/UpdateDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).UpdateDashboard(ctx, req.(*UpdateDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DashboardsService_ServiceDesc is the grpc.ServiceDesc for DashboardsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.monitoring.dashboard.v1.DashboardsService",
	HandlerType: (*DashboardsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDashboard",
			Handler:    _DashboardsService_CreateDashboard_Handler,
		},
		{
			MethodName: "ListDashboards",
			Handler:    _DashboardsService_ListDashboards_Handler,
		},
		{
			MethodName: "GetDashboard",
			Handler:    _DashboardsService_GetDashboard_Handler,
		},
		{
			MethodName: "DeleteDashboard",
			Handler:    _DashboardsService_DeleteDashboard_Handler,
		},
		{
			MethodName: "UpdateDashboard",
			Handler:    _DashboardsService_UpdateDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/monitoring/dashboard/v1/dashboards_service.proto",
}
