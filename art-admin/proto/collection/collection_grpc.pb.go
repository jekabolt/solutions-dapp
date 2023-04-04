// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: collection/collection.proto

package metadata

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CollectionsClient is the client API for Collections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionsClient interface {
	CreateNewCollection(ctx context.Context, in *CreateNewCollectionRequest, opts ...grpc.CallOption) (*CreateNewCollectionResponse, error)
	DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error)
	UpdateCollectionCapacity(ctx context.Context, in *UpdateCollectionCapacityRequest, opts ...grpc.CallOption) (*UpdateCollectionCapacityResponse, error)
	UpdateCollectionName(ctx context.Context, in *UpdateCollectionNameRequest, opts ...grpc.CallOption) (*UpdateCollectionNameResponse, error)
	GetAllCollections(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCollectionsResponse, error)
	GetCollectionByKey(ctx context.Context, in *GetCollectionByKeyRequest, opts ...grpc.CallOption) (*Collection, error)
}

type collectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionsClient(cc grpc.ClientConnInterface) CollectionsClient {
	return &collectionsClient{cc}
}

func (c *collectionsClient) CreateNewCollection(ctx context.Context, in *CreateNewCollectionRequest, opts ...grpc.CallOption) (*CreateNewCollectionResponse, error) {
	out := new(CreateNewCollectionResponse)
	err := c.cc.Invoke(ctx, "/collection.Collections/CreateNewCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionsClient) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error) {
	out := new(DeleteCollectionResponse)
	err := c.cc.Invoke(ctx, "/collection.Collections/DeleteCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionsClient) UpdateCollectionCapacity(ctx context.Context, in *UpdateCollectionCapacityRequest, opts ...grpc.CallOption) (*UpdateCollectionCapacityResponse, error) {
	out := new(UpdateCollectionCapacityResponse)
	err := c.cc.Invoke(ctx, "/collection.Collections/UpdateCollectionCapacity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionsClient) UpdateCollectionName(ctx context.Context, in *UpdateCollectionNameRequest, opts ...grpc.CallOption) (*UpdateCollectionNameResponse, error) {
	out := new(UpdateCollectionNameResponse)
	err := c.cc.Invoke(ctx, "/collection.Collections/UpdateCollectionName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionsClient) GetAllCollections(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCollectionsResponse, error) {
	out := new(GetAllCollectionsResponse)
	err := c.cc.Invoke(ctx, "/collection.Collections/GetAllCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionsClient) GetCollectionByKey(ctx context.Context, in *GetCollectionByKeyRequest, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/collection.Collections/GetCollectionByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionsServer is the server API for Collections service.
// All implementations should embed UnimplementedCollectionsServer
// for forward compatibility
type CollectionsServer interface {
	CreateNewCollection(context.Context, *CreateNewCollectionRequest) (*CreateNewCollectionResponse, error)
	DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error)
	UpdateCollectionCapacity(context.Context, *UpdateCollectionCapacityRequest) (*UpdateCollectionCapacityResponse, error)
	UpdateCollectionName(context.Context, *UpdateCollectionNameRequest) (*UpdateCollectionNameResponse, error)
	GetAllCollections(context.Context, *emptypb.Empty) (*GetAllCollectionsResponse, error)
	GetCollectionByKey(context.Context, *GetCollectionByKeyRequest) (*Collection, error)
}

// UnimplementedCollectionsServer should be embedded to have forward compatible implementations.
type UnimplementedCollectionsServer struct {
}

func (UnimplementedCollectionsServer) CreateNewCollection(context.Context, *CreateNewCollectionRequest) (*CreateNewCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewCollection not implemented")
}
func (UnimplementedCollectionsServer) DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedCollectionsServer) UpdateCollectionCapacity(context.Context, *UpdateCollectionCapacityRequest) (*UpdateCollectionCapacityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollectionCapacity not implemented")
}
func (UnimplementedCollectionsServer) UpdateCollectionName(context.Context, *UpdateCollectionNameRequest) (*UpdateCollectionNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollectionName not implemented")
}
func (UnimplementedCollectionsServer) GetAllCollections(context.Context, *emptypb.Empty) (*GetAllCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCollections not implemented")
}
func (UnimplementedCollectionsServer) GetCollectionByKey(context.Context, *GetCollectionByKeyRequest) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionByKey not implemented")
}

// UnsafeCollectionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionsServer will
// result in compilation errors.
type UnsafeCollectionsServer interface {
	mustEmbedUnimplementedCollectionsServer()
}

func RegisterCollectionsServer(s grpc.ServiceRegistrar, srv CollectionsServer) {
	s.RegisterService(&Collections_ServiceDesc, srv)
}

func _Collections_CreateNewCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServer).CreateNewCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.Collections/CreateNewCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServer).CreateNewCollection(ctx, req.(*CreateNewCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collections_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.Collections/DeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServer).DeleteCollection(ctx, req.(*DeleteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collections_UpdateCollectionCapacity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionCapacityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServer).UpdateCollectionCapacity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.Collections/UpdateCollectionCapacity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServer).UpdateCollectionCapacity(ctx, req.(*UpdateCollectionCapacityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collections_UpdateCollectionName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServer).UpdateCollectionName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.Collections/UpdateCollectionName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServer).UpdateCollectionName(ctx, req.(*UpdateCollectionNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collections_GetAllCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServer).GetAllCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.Collections/GetAllCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServer).GetAllCollections(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collections_GetCollectionByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServer).GetCollectionByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.Collections/GetCollectionByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServer).GetCollectionByKey(ctx, req.(*GetCollectionByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Collections_ServiceDesc is the grpc.ServiceDesc for Collections service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collections_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collection.Collections",
	HandlerType: (*CollectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewCollection",
			Handler:    _Collections_CreateNewCollection_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _Collections_DeleteCollection_Handler,
		},
		{
			MethodName: "UpdateCollectionCapacity",
			Handler:    _Collections_UpdateCollectionCapacity_Handler,
		},
		{
			MethodName: "UpdateCollectionName",
			Handler:    _Collections_UpdateCollectionName_Handler,
		},
		{
			MethodName: "GetAllCollections",
			Handler:    _Collections_GetAllCollections_Handler,
		},
		{
			MethodName: "GetCollectionByKey",
			Handler:    _Collections_GetCollectionByKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection/collection.proto",
}