// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: master.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MasterService_ConnectToStreamNode_FullMethodName    = "/MasterService/ConnectToStreamNode"
	MasterService_AssignAndGetStreamNode_FullMethodName = "/MasterService/AssignAndGetStreamNode"
	MasterService_GetStreamNode_FullMethodName          = "/MasterService/GetStreamNode"
	MasterService_TerminateSession_FullMethodName       = "/MasterService/TerminateSession"
)

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	ConnectToStreamNode(ctx context.Context, in *ConnectToStreamNodeRequest, opts ...grpc.CallOption) (*ConnectToStreamNodeResponse, error)
	AssignAndGetStreamNode(ctx context.Context, in *AssignAndGetStreamNodeRequest, opts ...grpc.CallOption) (*AssignAndGetStreamNodeResponse, error)
	GetStreamNode(ctx context.Context, in *GetStreamNodeRequest, opts ...grpc.CallOption) (*GetStreamNodeResponse, error)
	TerminateSession(ctx context.Context, in *TerminateSessionRequest, opts ...grpc.CallOption) (*TerminateSessionResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) ConnectToStreamNode(ctx context.Context, in *ConnectToStreamNodeRequest, opts ...grpc.CallOption) (*ConnectToStreamNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectToStreamNodeResponse)
	err := c.cc.Invoke(ctx, MasterService_ConnectToStreamNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) AssignAndGetStreamNode(ctx context.Context, in *AssignAndGetStreamNodeRequest, opts ...grpc.CallOption) (*AssignAndGetStreamNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignAndGetStreamNodeResponse)
	err := c.cc.Invoke(ctx, MasterService_AssignAndGetStreamNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetStreamNode(ctx context.Context, in *GetStreamNodeRequest, opts ...grpc.CallOption) (*GetStreamNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStreamNodeResponse)
	err := c.cc.Invoke(ctx, MasterService_GetStreamNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) TerminateSession(ctx context.Context, in *TerminateSessionRequest, opts ...grpc.CallOption) (*TerminateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminateSessionResponse)
	err := c.cc.Invoke(ctx, MasterService_TerminateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility.
type MasterServiceServer interface {
	ConnectToStreamNode(context.Context, *ConnectToStreamNodeRequest) (*ConnectToStreamNodeResponse, error)
	AssignAndGetStreamNode(context.Context, *AssignAndGetStreamNodeRequest) (*AssignAndGetStreamNodeResponse, error)
	GetStreamNode(context.Context, *GetStreamNodeRequest) (*GetStreamNodeResponse, error)
	TerminateSession(context.Context, *TerminateSessionRequest) (*TerminateSessionResponse, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMasterServiceServer struct{}

func (UnimplementedMasterServiceServer) ConnectToStreamNode(context.Context, *ConnectToStreamNodeRequest) (*ConnectToStreamNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectToStreamNode not implemented")
}
func (UnimplementedMasterServiceServer) AssignAndGetStreamNode(context.Context, *AssignAndGetStreamNodeRequest) (*AssignAndGetStreamNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignAndGetStreamNode not implemented")
}
func (UnimplementedMasterServiceServer) GetStreamNode(context.Context, *GetStreamNodeRequest) (*GetStreamNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamNode not implemented")
}
func (UnimplementedMasterServiceServer) TerminateSession(context.Context, *TerminateSessionRequest) (*TerminateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateSession not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}
func (UnimplementedMasterServiceServer) testEmbeddedByValue()                       {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	// If the following call pancis, it indicates UnimplementedMasterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_ConnectToStreamNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectToStreamNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ConnectToStreamNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_ConnectToStreamNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ConnectToStreamNode(ctx, req.(*ConnectToStreamNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_AssignAndGetStreamNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignAndGetStreamNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).AssignAndGetStreamNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_AssignAndGetStreamNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).AssignAndGetStreamNode(ctx, req.(*AssignAndGetStreamNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetStreamNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetStreamNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_GetStreamNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetStreamNode(ctx, req.(*GetStreamNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_TerminateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).TerminateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_TerminateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).TerminateSession(ctx, req.(*TerminateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectToStreamNode",
			Handler:    _MasterService_ConnectToStreamNode_Handler,
		},
		{
			MethodName: "AssignAndGetStreamNode",
			Handler:    _MasterService_AssignAndGetStreamNode_Handler,
		},
		{
			MethodName: "GetStreamNode",
			Handler:    _MasterService_GetStreamNode_Handler,
		},
		{
			MethodName: "TerminateSession",
			Handler:    _MasterService_TerminateSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
