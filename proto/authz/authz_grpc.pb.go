// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/authz/authz.proto

package authz

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

// AuthzClient is the client API for Authz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzClient interface {
	Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error)
	GetRolesForUser(ctx context.Context, in *RolesForUserRequest, opts ...grpc.CallOption) (*RolesForUserResponse, error)
}

type authzClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzClient(cc grpc.ClientConnInterface) AuthzClient {
	return &authzClient{cc}
}

func (c *authzClient) Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error) {
	out := new(VerificationResponse)
	err := c.cc.Invoke(ctx, "/authz.Authz/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzClient) GetRolesForUser(ctx context.Context, in *RolesForUserRequest, opts ...grpc.CallOption) (*RolesForUserResponse, error) {
	out := new(RolesForUserResponse)
	err := c.cc.Invoke(ctx, "/authz.Authz/GetRolesForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServer is the server API for Authz service.
// All implementations must embed UnimplementedAuthzServer
// for forward compatibility
type AuthzServer interface {
	Verify(context.Context, *VerificationRequest) (*VerificationResponse, error)
	GetRolesForUser(context.Context, *RolesForUserRequest) (*RolesForUserResponse, error)
	mustEmbedUnimplementedAuthzServer()
}

// UnimplementedAuthzServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServer struct {
}

func (UnimplementedAuthzServer) Verify(context.Context, *VerificationRequest) (*VerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthzServer) GetRolesForUser(context.Context, *RolesForUserRequest) (*RolesForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesForUser not implemented")
}
func (UnimplementedAuthzServer) mustEmbedUnimplementedAuthzServer() {}

// UnsafeAuthzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServer will
// result in compilation errors.
type UnsafeAuthzServer interface {
	mustEmbedUnimplementedAuthzServer()
}

func RegisterAuthzServer(s grpc.ServiceRegistrar, srv AuthzServer) {
	s.RegisterService(&Authz_ServiceDesc, srv)
}

func _Authz_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authz.Authz/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).Verify(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authz_GetRolesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).GetRolesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authz.Authz/GetRolesForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).GetRolesForUser(ctx, req.(*RolesForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authz_ServiceDesc is the grpc.ServiceDesc for Authz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authz.Authz",
	HandlerType: (*AuthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _Authz_Verify_Handler,
		},
		{
			MethodName: "GetRolesForUser",
			Handler:    _Authz_GetRolesForUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/authz/authz.proto",
}