// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: car.proto

package types

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
	Car_CreateCar_FullMethodName = "/car.Car/CreateCar"
)

// CarClient is the client API for Car service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarClient interface {
	CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarResponse, error)
}

type carClient struct {
	cc grpc.ClientConnInterface
}

func NewCarClient(cc grpc.ClientConnInterface) CarClient {
	return &carClient{cc}
}

func (c *carClient) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarResponse, error) {
	out := new(CreateCarResponse)
	err := c.cc.Invoke(ctx, Car_CreateCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServer is the server API for Car service.
// All implementations must embed UnimplementedCarServer
// for forward compatibility
type CarServer interface {
	CreateCar(context.Context, *CreateCarRequest) (*CreateCarResponse, error)
	mustEmbedUnimplementedCarServer()
}

// UnimplementedCarServer must be embedded to have forward compatible implementations.
type UnimplementedCarServer struct {
}

func (UnimplementedCarServer) CreateCar(context.Context, *CreateCarRequest) (*CreateCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}
func (UnimplementedCarServer) mustEmbedUnimplementedCarServer() {}

// UnsafeCarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServer will
// result in compilation errors.
type UnsafeCarServer interface {
	mustEmbedUnimplementedCarServer()
}

func RegisterCarServer(s grpc.ServiceRegistrar, srv CarServer) {
	s.RegisterService(&Car_ServiceDesc, srv)
}

func _Car_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_CreateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).CreateCar(ctx, req.(*CreateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Car_ServiceDesc is the grpc.ServiceDesc for Car service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Car_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "car.Car",
	HandlerType: (*CarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCar",
			Handler:    _Car_CreateCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "car.proto",
}