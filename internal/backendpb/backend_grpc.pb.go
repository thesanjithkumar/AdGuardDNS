// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: backend.proto

package backendpb

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

const (
	DNSService_GetDNSProfiles_FullMethodName         = "/DNSService/getDNSProfiles"
	DNSService_SaveDevicesBillingStat_FullMethodName = "/DNSService/saveDevicesBillingStat"
)

// DNSServiceClient is the client API for DNSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSServiceClient interface {
	// Gets DNS profiles.
	//
	// Field "sync_time" in DNSProfilesRequest - pass to return the latest updates after this time moment.
	//
	// The trailers headers will include a "sync_time", given in milliseconds,
	// that should be used for subsequent incremental DNS profile synchronization requests.
	GetDNSProfiles(ctx context.Context, in *DNSProfilesRequest, opts ...grpc.CallOption) (DNSService_GetDNSProfilesClient, error)
	// Stores devices activity.
	SaveDevicesBillingStat(ctx context.Context, opts ...grpc.CallOption) (DNSService_SaveDevicesBillingStatClient, error)
}

type dNSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSServiceClient(cc grpc.ClientConnInterface) DNSServiceClient {
	return &dNSServiceClient{cc}
}

func (c *dNSServiceClient) GetDNSProfiles(ctx context.Context, in *DNSProfilesRequest, opts ...grpc.CallOption) (DNSService_GetDNSProfilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &DNSService_ServiceDesc.Streams[0], DNSService_GetDNSProfiles_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dNSServiceGetDNSProfilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DNSService_GetDNSProfilesClient interface {
	Recv() (*DNSProfile, error)
	grpc.ClientStream
}

type dNSServiceGetDNSProfilesClient struct {
	grpc.ClientStream
}

func (x *dNSServiceGetDNSProfilesClient) Recv() (*DNSProfile, error) {
	m := new(DNSProfile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dNSServiceClient) SaveDevicesBillingStat(ctx context.Context, opts ...grpc.CallOption) (DNSService_SaveDevicesBillingStatClient, error) {
	stream, err := c.cc.NewStream(ctx, &DNSService_ServiceDesc.Streams[1], DNSService_SaveDevicesBillingStat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dNSServiceSaveDevicesBillingStatClient{stream}
	return x, nil
}

type DNSService_SaveDevicesBillingStatClient interface {
	Send(*DeviceBillingStat) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type dNSServiceSaveDevicesBillingStatClient struct {
	grpc.ClientStream
}

func (x *dNSServiceSaveDevicesBillingStatClient) Send(m *DeviceBillingStat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dNSServiceSaveDevicesBillingStatClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DNSServiceServer is the server API for DNSService service.
// All implementations must embed UnimplementedDNSServiceServer
// for forward compatibility
type DNSServiceServer interface {
	// Gets DNS profiles.
	//
	// Field "sync_time" in DNSProfilesRequest - pass to return the latest updates after this time moment.
	//
	// The trailers headers will include a "sync_time", given in milliseconds,
	// that should be used for subsequent incremental DNS profile synchronization requests.
	GetDNSProfiles(*DNSProfilesRequest, DNSService_GetDNSProfilesServer) error
	// Stores devices activity.
	SaveDevicesBillingStat(DNSService_SaveDevicesBillingStatServer) error
	mustEmbedUnimplementedDNSServiceServer()
}

// UnimplementedDNSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDNSServiceServer struct {
}

func (UnimplementedDNSServiceServer) GetDNSProfiles(*DNSProfilesRequest, DNSService_GetDNSProfilesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDNSProfiles not implemented")
}
func (UnimplementedDNSServiceServer) SaveDevicesBillingStat(DNSService_SaveDevicesBillingStatServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveDevicesBillingStat not implemented")
}
func (UnimplementedDNSServiceServer) mustEmbedUnimplementedDNSServiceServer() {}

// UnsafeDNSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSServiceServer will
// result in compilation errors.
type UnsafeDNSServiceServer interface {
	mustEmbedUnimplementedDNSServiceServer()
}

func RegisterDNSServiceServer(s grpc.ServiceRegistrar, srv DNSServiceServer) {
	s.RegisterService(&DNSService_ServiceDesc, srv)
}

func _DNSService_GetDNSProfiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DNSProfilesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DNSServiceServer).GetDNSProfiles(m, &dNSServiceGetDNSProfilesServer{stream})
}

type DNSService_GetDNSProfilesServer interface {
	Send(*DNSProfile) error
	grpc.ServerStream
}

type dNSServiceGetDNSProfilesServer struct {
	grpc.ServerStream
}

func (x *dNSServiceGetDNSProfilesServer) Send(m *DNSProfile) error {
	return x.ServerStream.SendMsg(m)
}

func _DNSService_SaveDevicesBillingStat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DNSServiceServer).SaveDevicesBillingStat(&dNSServiceSaveDevicesBillingStatServer{stream})
}

type DNSService_SaveDevicesBillingStatServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*DeviceBillingStat, error)
	grpc.ServerStream
}

type dNSServiceSaveDevicesBillingStatServer struct {
	grpc.ServerStream
}

func (x *dNSServiceSaveDevicesBillingStatServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dNSServiceSaveDevicesBillingStatServer) Recv() (*DeviceBillingStat, error) {
	m := new(DeviceBillingStat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DNSService_ServiceDesc is the grpc.ServiceDesc for DNSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DNSService",
	HandlerType: (*DNSServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getDNSProfiles",
			Handler:       _DNSService_GetDNSProfiles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "saveDevicesBillingStat",
			Handler:       _DNSService_SaveDevicesBillingStat_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "backend.proto",
}
