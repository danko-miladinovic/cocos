// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: manager/manager.proto

package manager

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

const (
	ManagerService_Process_FullMethodName = "/manager.ManagerService/Process"
)

// ManagerServiceClient is the client API for ManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerServiceClient interface {
	Process(ctx context.Context, opts ...grpc.CallOption) (ManagerService_ProcessClient, error)
}

type managerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerServiceClient(cc grpc.ClientConnInterface) ManagerServiceClient {
	return &managerServiceClient{cc}
}

func (c *managerServiceClient) Process(ctx context.Context, opts ...grpc.CallOption) (ManagerService_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManagerService_ServiceDesc.Streams[0], ManagerService_Process_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &managerServiceProcessClient{stream}
	return x, nil
}

type ManagerService_ProcessClient interface {
	Send(*ClientStreamMessage) error
	Recv() (*ComputationRunReq, error)
	grpc.ClientStream
}

type managerServiceProcessClient struct {
	grpc.ClientStream
}

func (x *managerServiceProcessClient) Send(m *ClientStreamMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managerServiceProcessClient) Recv() (*ComputationRunReq, error) {
	m := new(ComputationRunReq)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManagerServiceServer is the server API for ManagerService service.
// All implementations must embed UnimplementedManagerServiceServer
// for forward compatibility
type ManagerServiceServer interface {
	Process(ManagerService_ProcessServer) error
	mustEmbedUnimplementedManagerServiceServer()
}

// UnimplementedManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServiceServer struct {
}

func (UnimplementedManagerServiceServer) Process(ManagerService_ProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedManagerServiceServer) mustEmbedUnimplementedManagerServiceServer() {}

// UnsafeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServiceServer will
// result in compilation errors.
type UnsafeManagerServiceServer interface {
	mustEmbedUnimplementedManagerServiceServer()
}

func RegisterManagerServiceServer(s grpc.ServiceRegistrar, srv ManagerServiceServer) {
	s.RegisterService(&ManagerService_ServiceDesc, srv)
}

func _ManagerService_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagerServiceServer).Process(&managerServiceProcessServer{stream})
}

type ManagerService_ProcessServer interface {
	Send(*ComputationRunReq) error
	Recv() (*ClientStreamMessage, error)
	grpc.ServerStream
}

type managerServiceProcessServer struct {
	grpc.ServerStream
}

func (x *managerServiceProcessServer) Send(m *ComputationRunReq) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managerServiceProcessServer) Recv() (*ClientStreamMessage, error) {
	m := new(ClientStreamMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManagerService_ServiceDesc is the grpc.ServiceDesc for ManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.ManagerService",
	HandlerType: (*ManagerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _ManagerService_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "manager/manager.proto",
}