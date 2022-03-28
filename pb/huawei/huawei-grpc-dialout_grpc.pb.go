// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: huawei-grpc-dialout.proto

package huawei

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

// GRPCDataserviceClient is the client API for GRPCDataservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCDataserviceClient interface {
	DataPublish(ctx context.Context, opts ...grpc.CallOption) (GRPCDataservice_DataPublishClient, error)
}

type gRPCDataserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCDataserviceClient(cc grpc.ClientConnInterface) GRPCDataserviceClient {
	return &gRPCDataserviceClient{cc}
}

func (c *gRPCDataserviceClient) DataPublish(ctx context.Context, opts ...grpc.CallOption) (GRPCDataservice_DataPublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &GRPCDataservice_ServiceDesc.Streams[0], "/huawei_dialout.gRPCDataservice/dataPublish", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCDataserviceDataPublishClient{stream}
	return x, nil
}

type GRPCDataservice_DataPublishClient interface {
	Send(*ServiceArgs) error
	Recv() (*ServiceArgs, error)
	grpc.ClientStream
}

type gRPCDataserviceDataPublishClient struct {
	grpc.ClientStream
}

func (x *gRPCDataserviceDataPublishClient) Send(m *ServiceArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCDataserviceDataPublishClient) Recv() (*ServiceArgs, error) {
	m := new(ServiceArgs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GRPCDataserviceServer is the server API for GRPCDataservice service.
// All implementations must embed UnimplementedGRPCDataserviceServer
// for forward compatibility
type GRPCDataserviceServer interface {
	DataPublish(GRPCDataservice_DataPublishServer) error
	mustEmbedUnimplementedGRPCDataserviceServer()
}

// UnimplementedGRPCDataserviceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCDataserviceServer struct {
}

func (UnimplementedGRPCDataserviceServer) DataPublish(GRPCDataservice_DataPublishServer) error {
	return status.Errorf(codes.Unimplemented, "method DataPublish not implemented")
}
func (UnimplementedGRPCDataserviceServer) mustEmbedUnimplementedGRPCDataserviceServer() {}

// UnsafeGRPCDataserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCDataserviceServer will
// result in compilation errors.
type UnsafeGRPCDataserviceServer interface {
	mustEmbedUnimplementedGRPCDataserviceServer()
}

func RegisterGRPCDataserviceServer(s grpc.ServiceRegistrar, srv GRPCDataserviceServer) {
	s.RegisterService(&GRPCDataservice_ServiceDesc, srv)
}

func _GRPCDataservice_DataPublish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCDataserviceServer).DataPublish(&gRPCDataserviceDataPublishServer{stream})
}

type GRPCDataservice_DataPublishServer interface {
	Send(*ServiceArgs) error
	Recv() (*ServiceArgs, error)
	grpc.ServerStream
}

type gRPCDataserviceDataPublishServer struct {
	grpc.ServerStream
}

func (x *gRPCDataserviceDataPublishServer) Send(m *ServiceArgs) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCDataserviceDataPublishServer) Recv() (*ServiceArgs, error) {
	m := new(ServiceArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GRPCDataservice_ServiceDesc is the grpc.ServiceDesc for GRPCDataservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCDataservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "huawei_dialout.gRPCDataservice",
	HandlerType: (*GRPCDataserviceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "dataPublish",
			Handler:       _GRPCDataservice_DataPublish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "huawei-grpc-dialout.proto",
}
