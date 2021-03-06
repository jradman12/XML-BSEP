// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: connection_service.proto

package connection_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	GetConnections(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	GetConnectionRequests(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	GetRecommendedNewConnections(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	GetRecommendedJobOffers(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Offers, error)
	CreateConnection(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error)
	AcceptConnection(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error)
	ConnectionStatusForUsers(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error)
	BlockUser(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) GetConnections(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnectionRequests(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetConnectionRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetRecommendedNewConnections(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetRecommendedNewConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetRecommendedJobOffers(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Offers, error) {
	out := new(Offers)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetRecommendedJobOffers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) CreateConnection(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) AcceptConnection(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/AcceptConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ConnectionStatusForUsers(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/ConnectionStatusForUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) BlockUser(ctx context.Context, in *NewConnection, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	GetConnections(context.Context, *GetRequest) (*Users, error)
	GetConnectionRequests(context.Context, *GetRequest) (*Users, error)
	GetRecommendedNewConnections(context.Context, *GetRequest) (*Users, error)
	GetRecommendedJobOffers(context.Context, *GetRequest) (*Offers, error)
	CreateConnection(context.Context, *NewConnection) (*ConnectionResponse, error)
	AcceptConnection(context.Context, *NewConnection) (*ConnectionResponse, error)
	ConnectionStatusForUsers(context.Context, *NewConnection) (*ConnectionResponse, error)
	BlockUser(context.Context, *NewConnection) (*ConnectionResponse, error)
	MustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) GetConnections(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnections not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnectionRequests(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionRequests not implemented")
}
func (UnimplementedConnectionServiceServer) GetRecommendedNewConnections(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendedNewConnections not implemented")
}
func (UnimplementedConnectionServiceServer) GetRecommendedJobOffers(context.Context, *GetRequest) (*Offers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendedJobOffers not implemented")
}
func (UnimplementedConnectionServiceServer) CreateConnection(context.Context, *NewConnection) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (UnimplementedConnectionServiceServer) AcceptConnection(context.Context, *NewConnection) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptConnection not implemented")
}
func (UnimplementedConnectionServiceServer) ConnectionStatusForUsers(context.Context, *NewConnection) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionStatusForUsers not implemented")
}
func (UnimplementedConnectionServiceServer) BlockUser(context.Context, *NewConnection) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedConnectionServiceServer) MustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	MustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_GetConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnections(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnectionRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnectionRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetConnectionRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnectionRequests(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetRecommendedNewConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetRecommendedNewConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetRecommendedNewConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetRecommendedNewConnections(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetRecommendedJobOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetRecommendedJobOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetRecommendedJobOffers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetRecommendedJobOffers(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).CreateConnection(ctx, req.(*NewConnection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_AcceptConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).AcceptConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/AcceptConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).AcceptConnection(ctx, req.(*NewConnection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ConnectionStatusForUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ConnectionStatusForUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/ConnectionStatusForUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ConnectionStatusForUsers(ctx, req.(*NewConnection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).BlockUser(ctx, req.(*NewConnection))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connection_service.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConnections",
			Handler:    _ConnectionService_GetConnections_Handler,
		},
		{
			MethodName: "GetConnectionRequests",
			Handler:    _ConnectionService_GetConnectionRequests_Handler,
		},
		{
			MethodName: "GetRecommendedNewConnections",
			Handler:    _ConnectionService_GetRecommendedNewConnections_Handler,
		},
		{
			MethodName: "GetRecommendedJobOffers",
			Handler:    _ConnectionService_GetRecommendedJobOffers_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _ConnectionService_CreateConnection_Handler,
		},
		{
			MethodName: "AcceptConnection",
			Handler:    _ConnectionService_AcceptConnection_Handler,
		},
		{
			MethodName: "ConnectionStatusForUsers",
			Handler:    _ConnectionService_ConnectionStatusForUsers_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _ConnectionService_BlockUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection_service.proto",
}
