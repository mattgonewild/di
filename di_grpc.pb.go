// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: dat.proto

package di

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

// DataLayerClient is the client API for DataLayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataLayerClient interface {
	CreateImg(ctx context.Context, in *CreateImgRequest, opts ...grpc.CallOption) (*CreateImgResponse, error)
	ReadImg(ctx context.Context, in *ReadImgRequest, opts ...grpc.CallOption) (*ReadImgResponse, error)
	UpdateImg(ctx context.Context, in *UpdateImgRequest, opts ...grpc.CallOption) (*UpdateImgResponse, error)
	DeleteImg(ctx context.Context, in *DeleteImgRequest, opts ...grpc.CallOption) (*DeleteImgResponse, error)
}

type dataLayerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataLayerClient(cc grpc.ClientConnInterface) DataLayerClient {
	return &dataLayerClient{cc}
}

func (c *dataLayerClient) CreateImg(ctx context.Context, in *CreateImgRequest, opts ...grpc.CallOption) (*CreateImgResponse, error) {
	out := new(CreateImgResponse)
	err := c.cc.Invoke(ctx, "/DataLayer/CreateImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) ReadImg(ctx context.Context, in *ReadImgRequest, opts ...grpc.CallOption) (*ReadImgResponse, error) {
	out := new(ReadImgResponse)
	err := c.cc.Invoke(ctx, "/DataLayer/ReadImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) UpdateImg(ctx context.Context, in *UpdateImgRequest, opts ...grpc.CallOption) (*UpdateImgResponse, error) {
	out := new(UpdateImgResponse)
	err := c.cc.Invoke(ctx, "/DataLayer/UpdateImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLayerClient) DeleteImg(ctx context.Context, in *DeleteImgRequest, opts ...grpc.CallOption) (*DeleteImgResponse, error) {
	out := new(DeleteImgResponse)
	err := c.cc.Invoke(ctx, "/DataLayer/DeleteImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataLayerServer is the server API for DataLayer service.
// All implementations must embed UnimplementedDataLayerServer
// for forward compatibility
type DataLayerServer interface {
	CreateImg(context.Context, *CreateImgRequest) (*CreateImgResponse, error)
	ReadImg(context.Context, *ReadImgRequest) (*ReadImgResponse, error)
	UpdateImg(context.Context, *UpdateImgRequest) (*UpdateImgResponse, error)
	DeleteImg(context.Context, *DeleteImgRequest) (*DeleteImgResponse, error)
	mustEmbedUnimplementedDataLayerServer()
}

// UnimplementedDataLayerServer must be embedded to have forward compatible implementations.
type UnimplementedDataLayerServer struct {
}

func (UnimplementedDataLayerServer) CreateImg(context.Context, *CreateImgRequest) (*CreateImgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImg not implemented")
}
func (UnimplementedDataLayerServer) ReadImg(context.Context, *ReadImgRequest) (*ReadImgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadImg not implemented")
}
func (UnimplementedDataLayerServer) UpdateImg(context.Context, *UpdateImgRequest) (*UpdateImgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImg not implemented")
}
func (UnimplementedDataLayerServer) DeleteImg(context.Context, *DeleteImgRequest) (*DeleteImgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImg not implemented")
}
func (UnimplementedDataLayerServer) mustEmbedUnimplementedDataLayerServer() {}

// UnsafeDataLayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataLayerServer will
// result in compilation errors.
type UnsafeDataLayerServer interface {
	mustEmbedUnimplementedDataLayerServer()
}

func RegisterDataLayerServer(s grpc.ServiceRegistrar, srv DataLayerServer) {
	s.RegisterService(&DataLayer_ServiceDesc, srv)
}

func _DataLayer_CreateImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).CreateImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLayer/CreateImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).CreateImg(ctx, req.(*CreateImgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_ReadImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadImgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).ReadImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLayer/ReadImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).ReadImg(ctx, req.(*ReadImgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_UpdateImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).UpdateImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLayer/UpdateImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).UpdateImg(ctx, req.(*UpdateImgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLayer_DeleteImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLayerServer).DeleteImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLayer/DeleteImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLayerServer).DeleteImg(ctx, req.(*DeleteImgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataLayer_ServiceDesc is the grpc.ServiceDesc for DataLayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataLayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataLayer",
	HandlerType: (*DataLayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImg",
			Handler:    _DataLayer_CreateImg_Handler,
		},
		{
			MethodName: "ReadImg",
			Handler:    _DataLayer_ReadImg_Handler,
		},
		{
			MethodName: "UpdateImg",
			Handler:    _DataLayer_UpdateImg_Handler,
		},
		{
			MethodName: "DeleteImg",
			Handler:    _DataLayer_DeleteImg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dat.proto",
}
