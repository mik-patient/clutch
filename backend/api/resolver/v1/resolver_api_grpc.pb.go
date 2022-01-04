// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: resolver/v1/resolver_api.proto

package resolverv1

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

// ResolverAPIClient is the client API for ResolverAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResolverAPIClient interface {
	GetObjectSchemas(ctx context.Context, in *GetObjectSchemasRequest, opts ...grpc.CallOption) (*GetObjectSchemasResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
	Autocomplete(ctx context.Context, in *AutocompleteRequest, opts ...grpc.CallOption) (*AutocompleteResponse, error)
}

type resolverAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewResolverAPIClient(cc grpc.ClientConnInterface) ResolverAPIClient {
	return &resolverAPIClient{cc}
}

func (c *resolverAPIClient) GetObjectSchemas(ctx context.Context, in *GetObjectSchemasRequest, opts ...grpc.CallOption) (*GetObjectSchemasResponse, error) {
	out := new(GetObjectSchemasResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/GetObjectSchemas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolverAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolverAPIClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolverAPIClient) Autocomplete(ctx context.Context, in *AutocompleteRequest, opts ...grpc.CallOption) (*AutocompleteResponse, error) {
	out := new(AutocompleteResponse)
	err := c.cc.Invoke(ctx, "/clutch.resolver.v1.ResolverAPI/Autocomplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolverAPIServer is the server API for ResolverAPI service.
// All implementations should embed UnimplementedResolverAPIServer
// for forward compatibility
type ResolverAPIServer interface {
	GetObjectSchemas(context.Context, *GetObjectSchemasRequest) (*GetObjectSchemasResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
	Autocomplete(context.Context, *AutocompleteRequest) (*AutocompleteResponse, error)
}

// UnimplementedResolverAPIServer should be embedded to have forward compatible implementations.
type UnimplementedResolverAPIServer struct {
}

func (UnimplementedResolverAPIServer) GetObjectSchemas(context.Context, *GetObjectSchemasRequest) (*GetObjectSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectSchemas not implemented")
}
func (UnimplementedResolverAPIServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedResolverAPIServer) Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedResolverAPIServer) Autocomplete(context.Context, *AutocompleteRequest) (*AutocompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Autocomplete not implemented")
}

// UnsafeResolverAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResolverAPIServer will
// result in compilation errors.
type UnsafeResolverAPIServer interface {
	mustEmbedUnimplementedResolverAPIServer()
}

func RegisterResolverAPIServer(s grpc.ServiceRegistrar, srv ResolverAPIServer) {
	s.RegisterService(&ResolverAPI_ServiceDesc, srv)
}

func _ResolverAPI_GetObjectSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).GetObjectSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/GetObjectSchemas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).GetObjectSchemas(ctx, req.(*GetObjectSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolverAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolverAPI_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolverAPI_Autocomplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutocompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverAPIServer).Autocomplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.resolver.v1.ResolverAPI/Autocomplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverAPIServer).Autocomplete(ctx, req.(*AutocompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResolverAPI_ServiceDesc is the grpc.ServiceDesc for ResolverAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResolverAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.resolver.v1.ResolverAPI",
	HandlerType: (*ResolverAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObjectSchemas",
			Handler:    _ResolverAPI_GetObjectSchemas_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ResolverAPI_Search_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _ResolverAPI_Resolve_Handler,
		},
		{
			MethodName: "Autocomplete",
			Handler:    _ResolverAPI_Autocomplete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resolver/v1/resolver_api.proto",
}
