// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: endpoints.v1.proto

package service

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

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherClient interface {
	GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherReply, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherReply, error) {
	out := new(WeatherReply)
	err := c.cc.Invoke(ctx, "/endpoints.v1.Weather/GetWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServer is the server API for Weather service.
// All implementations should embed UnimplementedWeatherServer
// for forward compatibility
type WeatherServer interface {
	GetWeather(context.Context, *WeatherRequest) (*WeatherReply, error)
}

// UnimplementedWeatherServer should be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (UnimplementedWeatherServer) GetWeather(context.Context, *WeatherRequest) (*WeatherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}

// UnsafeWeatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServer will
// result in compilation errors.
type UnsafeWeatherServer interface {
	mustEmbedUnimplementedWeatherServer()
}

func RegisterWeatherServer(s grpc.ServiceRegistrar, srv WeatherServer) {
	s.RegisterService(&Weather_ServiceDesc, srv)
}

func _Weather_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.v1.Weather/GetWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).GetWeather(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Weather_ServiceDesc is the grpc.ServiceDesc for Weather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoints.v1.Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _Weather_GetWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "endpoints.v1.proto",
}
