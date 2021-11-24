// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authsv1

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

// AuthsServiceClient is the client API for AuthsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthsServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*OtpDto, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	ValidateOtp(ctx context.Context, in *ValidateOtpRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthsServiceClient(cc grpc.ClientConnInterface) AuthsServiceClient {
	return &authsServiceClient{cc}
}

func (c *authsServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*OtpDto, error) {
	out := new(OtpDto)
	err := c.cc.Invoke(ctx, "/auths.v1.AuthsService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authsServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/auths.v1.AuthsService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authsServiceClient) ValidateOtp(ctx context.Context, in *ValidateOtpRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auths.v1.AuthsService/ValidateOtp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthsServiceServer is the server API for AuthsService service.
// All implementations must embed UnimplementedAuthsServiceServer
// for forward compatibility
type AuthsServiceServer interface {
	Login(context.Context, *LoginRequest) (*OtpDto, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	ValidateOtp(context.Context, *ValidateOtpRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthsServiceServer()
}

// UnimplementedAuthsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthsServiceServer struct {
}

func (UnimplementedAuthsServiceServer) Login(context.Context, *LoginRequest) (*OtpDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthsServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthsServiceServer) ValidateOtp(context.Context, *ValidateOtpRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateOtp not implemented")
}
func (UnimplementedAuthsServiceServer) mustEmbedUnimplementedAuthsServiceServer() {}

// UnsafeAuthsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthsServiceServer will
// result in compilation errors.
type UnsafeAuthsServiceServer interface {
	mustEmbedUnimplementedAuthsServiceServer()
}

func RegisterAuthsServiceServer(s grpc.ServiceRegistrar, srv AuthsServiceServer) {
	s.RegisterService(&AuthsService_ServiceDesc, srv)
}

func _AuthsService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auths.v1.AuthsService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthsService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auths.v1.AuthsService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthsService_ValidateOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServiceServer).ValidateOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auths.v1.AuthsService/ValidateOtp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServiceServer).ValidateOtp(ctx, req.(*ValidateOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthsService_ServiceDesc is the grpc.ServiceDesc for AuthsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auths.v1.AuthsService",
	HandlerType: (*AuthsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthsService_Login_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthsService_RefreshToken_Handler,
		},
		{
			MethodName: "ValidateOtp",
			Handler:    _AuthsService_ValidateOtp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auths/v1/auths.proto",
}
