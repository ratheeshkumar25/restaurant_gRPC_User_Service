// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: user.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserServices_Signup_FullMethodName         = "/pbU.UserServices/Signup"
	UserServices_VerifyOTP_FullMethodName      = "/pbU.UserServices/VerifyOTP"
	UserServices_Login_FullMethodName          = "/pbU.UserServices/Login"
	UserServices_UserMenuList_FullMethodName   = "/pbU.UserServices/UserMenuList"
	UserServices_UserFoodByName_FullMethodName = "/pbU.UserServices/UserFoodByName"
	UserServices_UserMenuByID_FullMethodName   = "/pbU.UserServices/UserMenuByID"
)

// UserServicesClient is the client API for UserServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServicesClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupRespnse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPRespnse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UserMenuList(ctx context.Context, in *RNoParam, opts ...grpc.CallOption) (*MenuList, error)
	UserFoodByName(ctx context.Context, in *FoodByName, opts ...grpc.CallOption) (*MenuItem, error)
	UserMenuByID(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*MenuItem, error)
}

type userServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServicesClient(cc grpc.ClientConnInterface) UserServicesClient {
	return &userServicesClient{cc}
}

func (c *userServicesClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupRespnse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignupRespnse)
	err := c.cc.Invoke(ctx, UserServices_Signup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPRespnse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyOTPRespnse)
	err := c.cc.Invoke(ctx, UserServices_VerifyOTP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserServices_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UserMenuList(ctx context.Context, in *RNoParam, opts ...grpc.CallOption) (*MenuList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuList)
	err := c.cc.Invoke(ctx, UserServices_UserMenuList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UserFoodByName(ctx context.Context, in *FoodByName, opts ...grpc.CallOption) (*MenuItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuItem)
	err := c.cc.Invoke(ctx, UserServices_UserFoodByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UserMenuByID(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*MenuItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuItem)
	err := c.cc.Invoke(ctx, UserServices_UserMenuByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServicesServer is the server API for UserServices service.
// All implementations must embed UnimplementedUserServicesServer
// for forward compatibility
type UserServicesServer interface {
	Signup(context.Context, *SignupRequest) (*SignupRespnse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPRespnse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UserMenuList(context.Context, *RNoParam) (*MenuList, error)
	UserFoodByName(context.Context, *FoodByName) (*MenuItem, error)
	UserMenuByID(context.Context, *MenuID) (*MenuItem, error)
	mustEmbedUnimplementedUserServicesServer()
}

// UnimplementedUserServicesServer must be embedded to have forward compatible implementations.
type UnimplementedUserServicesServer struct {
}

func (UnimplementedUserServicesServer) Signup(context.Context, *SignupRequest) (*SignupRespnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServicesServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPRespnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedUserServicesServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServicesServer) UserMenuList(context.Context, *RNoParam) (*MenuList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserMenuList not implemented")
}
func (UnimplementedUserServicesServer) UserFoodByName(context.Context, *FoodByName) (*MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFoodByName not implemented")
}
func (UnimplementedUserServicesServer) UserMenuByID(context.Context, *MenuID) (*MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserMenuByID not implemented")
}
func (UnimplementedUserServicesServer) mustEmbedUnimplementedUserServicesServer() {}

// UnsafeUserServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServicesServer will
// result in compilation errors.
type UnsafeUserServicesServer interface {
	mustEmbedUnimplementedUserServicesServer()
}

func RegisterUserServicesServer(s grpc.ServiceRegistrar, srv UserServicesServer) {
	s.RegisterService(&UserServices_ServiceDesc, srv)
}

func _UserServices_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UserMenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UserMenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UserMenuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UserMenuList(ctx, req.(*RNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UserFoodByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoodByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UserFoodByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UserFoodByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UserFoodByName(ctx, req.(*FoodByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UserMenuByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UserMenuByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServices_UserMenuByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UserMenuByID(ctx, req.(*MenuID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServices_ServiceDesc is the grpc.ServiceDesc for UserServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbU.UserServices",
	HandlerType: (*UserServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserServices_Signup_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _UserServices_VerifyOTP_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserServices_Login_Handler,
		},
		{
			MethodName: "UserMenuList",
			Handler:    _UserServices_UserMenuList_Handler,
		},
		{
			MethodName: "UserFoodByName",
			Handler:    _UserServices_UserFoodByName_Handler,
		},
		{
			MethodName: "UserMenuByID",
			Handler:    _UserServices_UserMenuByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
