// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package user

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

// B24UserServiceClient is the client API for B24UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type B24UserServiceClient interface {
	// auth
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*UserResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// company user
	GetUserByEmail(ctx context.Context, in *UserGetByEmailRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserById(ctx context.Context, in *UserGetByIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsers(ctx context.Context, in *UsersGetRequest, opts ...grpc.CallOption) (*UsersGetResponse, error)
	UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UpdateUserRole(ctx context.Context, in *UserRoleUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UpdateUserAvatar(ctx context.Context, in *UserAvatarUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ChangePassword(ctx context.Context, in *UserChangePasswordRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// customer user
	AddCustomerUser(ctx context.Context, in *CustomerUserAddRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	GetCustomerUserByEmail(ctx context.Context, in *CustomerUserGetByEmailRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	GetCustomerUserById(ctx context.Context, in *CustomerUserGetByIdRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	GetCustomerUsers(ctx context.Context, in *CustomerUsersGetRequest, opts ...grpc.CallOption) (*CustomerUsersGetResponse, error)
	UpdateCustomerUser(ctx context.Context, in *CustomerUserUpdateRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	UpdateCustomerUserRole(ctx context.Context, in *CustomerUserRoleUpdateRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	UpdateCustomerUserAvatar(ctx context.Context, in *CustomerUserAvatarUpdateRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	ChangeCustomerUserPassword(ctx context.Context, in *CustomerUserChangePasswordRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error)
	//role
	CreateRole(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	GetRole(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	GetRoleByName(ctx context.Context, in *RoleGetByNameRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	GetRoles(ctx context.Context, in *RolesGetRequest, opts ...grpc.CallOption) (*RolesGetResponse, error)
	UpdateRole(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleResponse, error)
}

type b24UserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewB24UserServiceClient(cc grpc.ClientConnInterface) B24UserServiceClient {
	return &b24UserServiceClient{cc}
}

func (c *b24UserServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetUserByEmail(ctx context.Context, in *UserGetByEmailRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetUserById(ctx context.Context, in *UserGetByIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetUsers(ctx context.Context, in *UsersGetRequest, opts ...grpc.CallOption) (*UsersGetResponse, error) {
	out := new(UsersGetResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateUserRole(ctx context.Context, in *UserRoleUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateUserAvatar(ctx context.Context, in *UserAvatarUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateUserAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) ChangePassword(ctx context.Context, in *UserChangePasswordRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) AddCustomerUser(ctx context.Context, in *CustomerUserAddRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/AddCustomerUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetCustomerUserByEmail(ctx context.Context, in *CustomerUserGetByEmailRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetCustomerUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetCustomerUserById(ctx context.Context, in *CustomerUserGetByIdRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetCustomerUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetCustomerUsers(ctx context.Context, in *CustomerUsersGetRequest, opts ...grpc.CallOption) (*CustomerUsersGetResponse, error) {
	out := new(CustomerUsersGetResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetCustomerUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateCustomerUser(ctx context.Context, in *CustomerUserUpdateRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateCustomerUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateCustomerUserRole(ctx context.Context, in *CustomerUserRoleUpdateRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateCustomerUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateCustomerUserAvatar(ctx context.Context, in *CustomerUserAvatarUpdateRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateCustomerUserAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) ChangeCustomerUserPassword(ctx context.Context, in *CustomerUserChangePasswordRequest, opts ...grpc.CallOption) (*CustomerUserResponse, error) {
	out := new(CustomerUserResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/ChangeCustomerUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) CreateRole(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetRole(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetRoleByName(ctx context.Context, in *RoleGetByNameRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetRoleByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) GetRoles(ctx context.Context, in *RolesGetRequest, opts ...grpc.CallOption) (*RolesGetResponse, error) {
	out := new(RolesGetResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *b24UserServiceClient) UpdateRole(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/user.B24UserService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// B24UserServiceServer is the server API for B24UserService service.
// All implementations must embed UnimplementedB24UserServiceServer
// for forward compatibility
type B24UserServiceServer interface {
	// auth
	SignUp(context.Context, *SignUpRequest) (*UserResponse, error)
	SignIn(context.Context, *SignInRequest) (*UserResponse, error)
	// company user
	GetUserByEmail(context.Context, *UserGetByEmailRequest) (*UserResponse, error)
	GetUserById(context.Context, *UserGetByIdRequest) (*UserResponse, error)
	GetUsers(context.Context, *UsersGetRequest) (*UsersGetResponse, error)
	UpdateUser(context.Context, *UserUpdateRequest) (*UserResponse, error)
	UpdateUserRole(context.Context, *UserRoleUpdateRequest) (*UserResponse, error)
	UpdateUserAvatar(context.Context, *UserAvatarUpdateRequest) (*UserResponse, error)
	ChangePassword(context.Context, *UserChangePasswordRequest) (*UserResponse, error)
	// customer user
	AddCustomerUser(context.Context, *CustomerUserAddRequest) (*CustomerUserResponse, error)
	GetCustomerUserByEmail(context.Context, *CustomerUserGetByEmailRequest) (*CustomerUserResponse, error)
	GetCustomerUserById(context.Context, *CustomerUserGetByIdRequest) (*CustomerUserResponse, error)
	GetCustomerUsers(context.Context, *CustomerUsersGetRequest) (*CustomerUsersGetResponse, error)
	UpdateCustomerUser(context.Context, *CustomerUserUpdateRequest) (*CustomerUserResponse, error)
	UpdateCustomerUserRole(context.Context, *CustomerUserRoleUpdateRequest) (*CustomerUserResponse, error)
	UpdateCustomerUserAvatar(context.Context, *CustomerUserAvatarUpdateRequest) (*CustomerUserResponse, error)
	ChangeCustomerUserPassword(context.Context, *CustomerUserChangePasswordRequest) (*CustomerUserResponse, error)
	//role
	CreateRole(context.Context, *RoleCreateRequest) (*RoleResponse, error)
	GetRole(context.Context, *RoleGetRequest) (*RoleResponse, error)
	GetRoleByName(context.Context, *RoleGetByNameRequest) (*RoleResponse, error)
	GetRoles(context.Context, *RolesGetRequest) (*RolesGetResponse, error)
	UpdateRole(context.Context, *RoleUpdateRequest) (*RoleResponse, error)
	mustEmbedUnimplementedB24UserServiceServer()
}

// UnimplementedB24UserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedB24UserServiceServer struct {
}

func (UnimplementedB24UserServiceServer) SignUp(context.Context, *SignUpRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedB24UserServiceServer) SignIn(context.Context, *SignInRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedB24UserServiceServer) GetUserByEmail(context.Context, *UserGetByEmailRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedB24UserServiceServer) GetUserById(context.Context, *UserGetByIdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedB24UserServiceServer) GetUsers(context.Context, *UsersGetRequest) (*UsersGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateUser(context.Context, *UserUpdateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateUserRole(context.Context, *UserRoleUpdateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateUserAvatar(context.Context, *UserAvatarUpdateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAvatar not implemented")
}
func (UnimplementedB24UserServiceServer) ChangePassword(context.Context, *UserChangePasswordRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedB24UserServiceServer) AddCustomerUser(context.Context, *CustomerUserAddRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomerUser not implemented")
}
func (UnimplementedB24UserServiceServer) GetCustomerUserByEmail(context.Context, *CustomerUserGetByEmailRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerUserByEmail not implemented")
}
func (UnimplementedB24UserServiceServer) GetCustomerUserById(context.Context, *CustomerUserGetByIdRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerUserById not implemented")
}
func (UnimplementedB24UserServiceServer) GetCustomerUsers(context.Context, *CustomerUsersGetRequest) (*CustomerUsersGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerUsers not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateCustomerUser(context.Context, *CustomerUserUpdateRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerUser not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateCustomerUserRole(context.Context, *CustomerUserRoleUpdateRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerUserRole not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateCustomerUserAvatar(context.Context, *CustomerUserAvatarUpdateRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerUserAvatar not implemented")
}
func (UnimplementedB24UserServiceServer) ChangeCustomerUserPassword(context.Context, *CustomerUserChangePasswordRequest) (*CustomerUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCustomerUserPassword not implemented")
}
func (UnimplementedB24UserServiceServer) CreateRole(context.Context, *RoleCreateRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedB24UserServiceServer) GetRole(context.Context, *RoleGetRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedB24UserServiceServer) GetRoleByName(context.Context, *RoleGetByNameRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleByName not implemented")
}
func (UnimplementedB24UserServiceServer) GetRoles(context.Context, *RolesGetRequest) (*RolesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedB24UserServiceServer) UpdateRole(context.Context, *RoleUpdateRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedB24UserServiceServer) mustEmbedUnimplementedB24UserServiceServer() {}

// UnsafeB24UserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to B24UserServiceServer will
// result in compilation errors.
type UnsafeB24UserServiceServer interface {
	mustEmbedUnimplementedB24UserServiceServer()
}

func RegisterB24UserServiceServer(s grpc.ServiceRegistrar, srv B24UserServiceServer) {
	s.RegisterService(&B24UserService_ServiceDesc, srv)
}

func _B24UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetUserByEmail(ctx, req.(*UserGetByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetUserById(ctx, req.(*UserGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetUsers(ctx, req.(*UsersGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateUser(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateUserRole(ctx, req.(*UserRoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAvatarUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateUserAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateUserAvatar(ctx, req.(*UserAvatarUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).ChangePassword(ctx, req.(*UserChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_AddCustomerUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).AddCustomerUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/AddCustomerUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).AddCustomerUser(ctx, req.(*CustomerUserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetCustomerUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserGetByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetCustomerUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetCustomerUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetCustomerUserByEmail(ctx, req.(*CustomerUserGetByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetCustomerUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetCustomerUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetCustomerUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetCustomerUserById(ctx, req.(*CustomerUserGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetCustomerUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUsersGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetCustomerUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetCustomerUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetCustomerUsers(ctx, req.(*CustomerUsersGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateCustomerUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateCustomerUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateCustomerUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateCustomerUser(ctx, req.(*CustomerUserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateCustomerUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserRoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateCustomerUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateCustomerUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateCustomerUserRole(ctx, req.(*CustomerUserRoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateCustomerUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserAvatarUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateCustomerUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateCustomerUserAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateCustomerUserAvatar(ctx, req.(*CustomerUserAvatarUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_ChangeCustomerUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUserChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).ChangeCustomerUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/ChangeCustomerUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).ChangeCustomerUserPassword(ctx, req.(*CustomerUserChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).CreateRole(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetRole(ctx, req.(*RoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetRoleByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetRoleByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetRoleByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetRoleByName(ctx, req.(*RoleGetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).GetRoles(ctx, req.(*RolesGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _B24UserService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(B24UserServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.B24UserService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(B24UserServiceServer).UpdateRole(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// B24UserService_ServiceDesc is the grpc.ServiceDesc for B24UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var B24UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.B24UserService",
	HandlerType: (*B24UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _B24UserService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _B24UserService_SignIn_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _B24UserService_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _B24UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _B24UserService_GetUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _B24UserService_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _B24UserService_UpdateUserRole_Handler,
		},
		{
			MethodName: "UpdateUserAvatar",
			Handler:    _B24UserService_UpdateUserAvatar_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _B24UserService_ChangePassword_Handler,
		},
		{
			MethodName: "AddCustomerUser",
			Handler:    _B24UserService_AddCustomerUser_Handler,
		},
		{
			MethodName: "GetCustomerUserByEmail",
			Handler:    _B24UserService_GetCustomerUserByEmail_Handler,
		},
		{
			MethodName: "GetCustomerUserById",
			Handler:    _B24UserService_GetCustomerUserById_Handler,
		},
		{
			MethodName: "GetCustomerUsers",
			Handler:    _B24UserService_GetCustomerUsers_Handler,
		},
		{
			MethodName: "UpdateCustomerUser",
			Handler:    _B24UserService_UpdateCustomerUser_Handler,
		},
		{
			MethodName: "UpdateCustomerUserRole",
			Handler:    _B24UserService_UpdateCustomerUserRole_Handler,
		},
		{
			MethodName: "UpdateCustomerUserAvatar",
			Handler:    _B24UserService_UpdateCustomerUserAvatar_Handler,
		},
		{
			MethodName: "ChangeCustomerUserPassword",
			Handler:    _B24UserService_ChangeCustomerUserPassword_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _B24UserService_CreateRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _B24UserService_GetRole_Handler,
		},
		{
			MethodName: "GetRoleByName",
			Handler:    _B24UserService_GetRoleByName_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _B24UserService_GetRoles_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _B24UserService_UpdateRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
