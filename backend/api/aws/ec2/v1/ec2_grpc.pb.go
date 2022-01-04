// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: aws/ec2/v1/ec2.proto

package ec2v1

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

// EC2APIClient is the client API for EC2API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EC2APIClient interface {
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error)
	TerminateInstance(ctx context.Context, in *TerminateInstanceRequest, opts ...grpc.CallOption) (*TerminateInstanceResponse, error)
	ResizeAutoscalingGroup(ctx context.Context, in *ResizeAutoscalingGroupRequest, opts ...grpc.CallOption) (*ResizeAutoscalingGroupResponse, error)
	RebootInstance(ctx context.Context, in *RebootInstanceRequest, opts ...grpc.CallOption) (*RebootInstanceResponse, error)
}

type eC2APIClient struct {
	cc grpc.ClientConnInterface
}

func NewEC2APIClient(cc grpc.ClientConnInterface) EC2APIClient {
	return &eC2APIClient{cc}
}

func (c *eC2APIClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error) {
	out := new(GetInstanceResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.ec2.v1.EC2API/GetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eC2APIClient) TerminateInstance(ctx context.Context, in *TerminateInstanceRequest, opts ...grpc.CallOption) (*TerminateInstanceResponse, error) {
	out := new(TerminateInstanceResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.ec2.v1.EC2API/TerminateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eC2APIClient) ResizeAutoscalingGroup(ctx context.Context, in *ResizeAutoscalingGroupRequest, opts ...grpc.CallOption) (*ResizeAutoscalingGroupResponse, error) {
	out := new(ResizeAutoscalingGroupResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.ec2.v1.EC2API/ResizeAutoscalingGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eC2APIClient) RebootInstance(ctx context.Context, in *RebootInstanceRequest, opts ...grpc.CallOption) (*RebootInstanceResponse, error) {
	out := new(RebootInstanceResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.ec2.v1.EC2API/RebootInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EC2APIServer is the server API for EC2API service.
// All implementations should embed UnimplementedEC2APIServer
// for forward compatibility
type EC2APIServer interface {
	GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error)
	TerminateInstance(context.Context, *TerminateInstanceRequest) (*TerminateInstanceResponse, error)
	ResizeAutoscalingGroup(context.Context, *ResizeAutoscalingGroupRequest) (*ResizeAutoscalingGroupResponse, error)
	RebootInstance(context.Context, *RebootInstanceRequest) (*RebootInstanceResponse, error)
}

// UnimplementedEC2APIServer should be embedded to have forward compatible implementations.
type UnimplementedEC2APIServer struct {
}

func (UnimplementedEC2APIServer) GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedEC2APIServer) TerminateInstance(context.Context, *TerminateInstanceRequest) (*TerminateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateInstance not implemented")
}
func (UnimplementedEC2APIServer) ResizeAutoscalingGroup(context.Context, *ResizeAutoscalingGroupRequest) (*ResizeAutoscalingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeAutoscalingGroup not implemented")
}
func (UnimplementedEC2APIServer) RebootInstance(context.Context, *RebootInstanceRequest) (*RebootInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebootInstance not implemented")
}

// UnsafeEC2APIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EC2APIServer will
// result in compilation errors.
type UnsafeEC2APIServer interface {
	mustEmbedUnimplementedEC2APIServer()
}

func RegisterEC2APIServer(s grpc.ServiceRegistrar, srv EC2APIServer) {
	s.RegisterService(&EC2API_ServiceDesc, srv)
}

func _EC2API_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EC2APIServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.ec2.v1.EC2API/GetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EC2APIServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EC2API_TerminateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EC2APIServer).TerminateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.ec2.v1.EC2API/TerminateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EC2APIServer).TerminateInstance(ctx, req.(*TerminateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EC2API_ResizeAutoscalingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeAutoscalingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EC2APIServer).ResizeAutoscalingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.ec2.v1.EC2API/ResizeAutoscalingGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EC2APIServer).ResizeAutoscalingGroup(ctx, req.(*ResizeAutoscalingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EC2API_RebootInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebootInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EC2APIServer).RebootInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.ec2.v1.EC2API/RebootInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EC2APIServer).RebootInstance(ctx, req.(*RebootInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EC2API_ServiceDesc is the grpc.ServiceDesc for EC2API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EC2API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.aws.ec2.v1.EC2API",
	HandlerType: (*EC2APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInstance",
			Handler:    _EC2API_GetInstance_Handler,
		},
		{
			MethodName: "TerminateInstance",
			Handler:    _EC2API_TerminateInstance_Handler,
		},
		{
			MethodName: "ResizeAutoscalingGroup",
			Handler:    _EC2API_ResizeAutoscalingGroup_Handler,
		},
		{
			MethodName: "RebootInstance",
			Handler:    _EC2API_RebootInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aws/ec2/v1/ec2.proto",
}
