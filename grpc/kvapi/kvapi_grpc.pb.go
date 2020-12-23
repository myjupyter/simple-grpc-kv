// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kvapi

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// KVStorageClient is the client API for KVStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVStorageClient interface {
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ValueResponse, error)
	Set(ctx context.Context, in *KeyValueRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *KeyValueRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type kVStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewKVStorageClient(cc grpc.ClientConnInterface) KVStorageClient {
	return &kVStorageClient{cc}
}

func (c *kVStorageClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ValueResponse, error) {
	out := new(ValueResponse)
	err := c.cc.Invoke(ctx, "/demo.kvapi.v1.KVStorage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStorageClient) Set(ctx context.Context, in *KeyValueRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/demo.kvapi.v1.KVStorage/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStorageClient) Update(ctx context.Context, in *KeyValueRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/demo.kvapi.v1.KVStorage/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStorageClient) Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/demo.kvapi.v1.KVStorage/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVStorageServer is the server API for KVStorage service.
// All implementations must embed UnimplementedKVStorageServer
// for forward compatibility
type KVStorageServer interface {
	Get(context.Context, *KeyRequest) (*ValueResponse, error)
	Set(context.Context, *KeyValueRequest) (*empty.Empty, error)
	Update(context.Context, *KeyValueRequest) (*empty.Empty, error)
	Delete(context.Context, *KeyRequest) (*empty.Empty, error)
	mustEmbedUnimplementedKVStorageServer()
}

// UnimplementedKVStorageServer must be embedded to have forward compatible implementations.
type UnimplementedKVStorageServer struct {
}

func (UnimplementedKVStorageServer) Get(context.Context, *KeyRequest) (*ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKVStorageServer) Set(context.Context, *KeyValueRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKVStorageServer) Update(context.Context, *KeyValueRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKVStorageServer) Delete(context.Context, *KeyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKVStorageServer) mustEmbedUnimplementedKVStorageServer() {}

// UnsafeKVStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVStorageServer will
// result in compilation errors.
type UnsafeKVStorageServer interface {
	mustEmbedUnimplementedKVStorageServer()
}

func RegisterKVStorageServer(s grpc.ServiceRegistrar, srv KVStorageServer) {
	s.RegisterService(&_KVStorage_serviceDesc, srv)
}

func _KVStorage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.kvapi.v1.KVStorage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStorageServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStorage_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStorageServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.kvapi.v1.KVStorage/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStorageServer).Set(ctx, req.(*KeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStorage_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStorageServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.kvapi.v1.KVStorage/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStorageServer).Update(ctx, req.(*KeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStorage_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStorageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.kvapi.v1.KVStorage/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStorageServer).Delete(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.kvapi.v1.KVStorage",
	HandlerType: (*KVStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVStorage_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _KVStorage_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _KVStorage_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVStorage_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvapi.proto",
}
