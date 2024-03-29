// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package communicationsv1

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

// CommunicationServiceClient is the client API for CommunicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationServiceClient interface {
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	CreateChannelTemplate(ctx context.Context, in *CreateChannelTemplateRequest, opts ...grpc.CallOption) (*CreateChannelTemplateResponse, error)
	UpdateChannelTemplate(ctx context.Context, in *UpdateChannelTemplateRequest, opts ...grpc.CallOption) (*UpdateChannelTemplateResponse, error)
	CreateChannelProvider(ctx context.Context, in *CreateChannelProviderRequest, opts ...grpc.CallOption) (*CreateChannelProviderResponse, error)
	CreateCommunication(ctx context.Context, in *CreateCommunicationRequest, opts ...grpc.CallOption) (*CreateCommunicationResponse, error)
}

type communicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationServiceClient(cc grpc.ClientConnInterface) CommunicationServiceClient {
	return &communicationServiceClient{cc}
}

func (c *communicationServiceClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, "/communications.v1.CommunicationService/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) CreateChannelTemplate(ctx context.Context, in *CreateChannelTemplateRequest, opts ...grpc.CallOption) (*CreateChannelTemplateResponse, error) {
	out := new(CreateChannelTemplateResponse)
	err := c.cc.Invoke(ctx, "/communications.v1.CommunicationService/CreateChannelTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) UpdateChannelTemplate(ctx context.Context, in *UpdateChannelTemplateRequest, opts ...grpc.CallOption) (*UpdateChannelTemplateResponse, error) {
	out := new(UpdateChannelTemplateResponse)
	err := c.cc.Invoke(ctx, "/communications.v1.CommunicationService/UpdateChannelTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) CreateChannelProvider(ctx context.Context, in *CreateChannelProviderRequest, opts ...grpc.CallOption) (*CreateChannelProviderResponse, error) {
	out := new(CreateChannelProviderResponse)
	err := c.cc.Invoke(ctx, "/communications.v1.CommunicationService/CreateChannelProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationServiceClient) CreateCommunication(ctx context.Context, in *CreateCommunicationRequest, opts ...grpc.CallOption) (*CreateCommunicationResponse, error) {
	out := new(CreateCommunicationResponse)
	err := c.cc.Invoke(ctx, "/communications.v1.CommunicationService/CreateCommunication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServiceServer is the server API for CommunicationService service.
// All implementations must embed UnimplementedCommunicationServiceServer
// for forward compatibility
type CommunicationServiceServer interface {
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	CreateChannelTemplate(context.Context, *CreateChannelTemplateRequest) (*CreateChannelTemplateResponse, error)
	UpdateChannelTemplate(context.Context, *UpdateChannelTemplateRequest) (*UpdateChannelTemplateResponse, error)
	CreateChannelProvider(context.Context, *CreateChannelProviderRequest) (*CreateChannelProviderResponse, error)
	CreateCommunication(context.Context, *CreateCommunicationRequest) (*CreateCommunicationResponse, error)
	mustEmbedUnimplementedCommunicationServiceServer()
}

// UnimplementedCommunicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicationServiceServer struct {
}

func (UnimplementedCommunicationServiceServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedCommunicationServiceServer) CreateChannelTemplate(context.Context, *CreateChannelTemplateRequest) (*CreateChannelTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannelTemplate not implemented")
}
func (UnimplementedCommunicationServiceServer) UpdateChannelTemplate(context.Context, *UpdateChannelTemplateRequest) (*UpdateChannelTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannelTemplate not implemented")
}
func (UnimplementedCommunicationServiceServer) CreateChannelProvider(context.Context, *CreateChannelProviderRequest) (*CreateChannelProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannelProvider not implemented")
}
func (UnimplementedCommunicationServiceServer) CreateCommunication(context.Context, *CreateCommunicationRequest) (*CreateCommunicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunication not implemented")
}
func (UnimplementedCommunicationServiceServer) mustEmbedUnimplementedCommunicationServiceServer() {}

// UnsafeCommunicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationServiceServer will
// result in compilation errors.
type UnsafeCommunicationServiceServer interface {
	mustEmbedUnimplementedCommunicationServiceServer()
}

func RegisterCommunicationServiceServer(s grpc.ServiceRegistrar, srv CommunicationServiceServer) {
	s.RegisterService(&CommunicationService_ServiceDesc, srv)
}

func _CommunicationService_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.v1.CommunicationService/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_CreateChannelTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).CreateChannelTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.v1.CommunicationService/CreateChannelTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).CreateChannelTemplate(ctx, req.(*CreateChannelTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_UpdateChannelTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).UpdateChannelTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.v1.CommunicationService/UpdateChannelTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).UpdateChannelTemplate(ctx, req.(*UpdateChannelTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_CreateChannelProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).CreateChannelProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.v1.CommunicationService/CreateChannelProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).CreateChannelProvider(ctx, req.(*CreateChannelProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationService_CreateCommunication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).CreateCommunication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.v1.CommunicationService/CreateCommunication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).CreateCommunication(ctx, req.(*CreateCommunicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunicationService_ServiceDesc is the grpc.ServiceDesc for CommunicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communications.v1.CommunicationService",
	HandlerType: (*CommunicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _CommunicationService_CreateChannel_Handler,
		},
		{
			MethodName: "CreateChannelTemplate",
			Handler:    _CommunicationService_CreateChannelTemplate_Handler,
		},
		{
			MethodName: "UpdateChannelTemplate",
			Handler:    _CommunicationService_UpdateChannelTemplate_Handler,
		},
		{
			MethodName: "CreateChannelProvider",
			Handler:    _CommunicationService_CreateChannelProvider_Handler,
		},
		{
			MethodName: "CreateCommunication",
			Handler:    _CommunicationService_CreateCommunication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communications/v1/communications.proto",
}
