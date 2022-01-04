// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: featureflag/v1/featureflag.proto

package featureflagv1

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

// FeatureFlagAPIClient is the client API for FeatureFlagAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureFlagAPIClient interface {
	GetFlags(ctx context.Context, in *GetFlagsRequest, opts ...grpc.CallOption) (*GetFlagsResponse, error)
}

type featureFlagAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureFlagAPIClient(cc grpc.ClientConnInterface) FeatureFlagAPIClient {
	return &featureFlagAPIClient{cc}
}

func (c *featureFlagAPIClient) GetFlags(ctx context.Context, in *GetFlagsRequest, opts ...grpc.CallOption) (*GetFlagsResponse, error) {
	out := new(GetFlagsResponse)
	err := c.cc.Invoke(ctx, "/clutch.featureflag.v1.FeatureFlagAPI/GetFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureFlagAPIServer is the server API for FeatureFlagAPI service.
// All implementations should embed UnimplementedFeatureFlagAPIServer
// for forward compatibility
type FeatureFlagAPIServer interface {
	GetFlags(context.Context, *GetFlagsRequest) (*GetFlagsResponse, error)
}

// UnimplementedFeatureFlagAPIServer should be embedded to have forward compatible implementations.
type UnimplementedFeatureFlagAPIServer struct {
}

func (UnimplementedFeatureFlagAPIServer) GetFlags(context.Context, *GetFlagsRequest) (*GetFlagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlags not implemented")
}

// UnsafeFeatureFlagAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureFlagAPIServer will
// result in compilation errors.
type UnsafeFeatureFlagAPIServer interface {
	mustEmbedUnimplementedFeatureFlagAPIServer()
}

func RegisterFeatureFlagAPIServer(s grpc.ServiceRegistrar, srv FeatureFlagAPIServer) {
	s.RegisterService(&FeatureFlagAPI_ServiceDesc, srv)
}

func _FeatureFlagAPI_GetFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagAPIServer).GetFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.featureflag.v1.FeatureFlagAPI/GetFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagAPIServer).GetFlags(ctx, req.(*GetFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureFlagAPI_ServiceDesc is the grpc.ServiceDesc for FeatureFlagAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureFlagAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.featureflag.v1.FeatureFlagAPI",
	HandlerType: (*FeatureFlagAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlags",
			Handler:    _FeatureFlagAPI_GetFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "featureflag/v1/featureflag.proto",
}
