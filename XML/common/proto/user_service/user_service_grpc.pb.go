// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user_service.proto

package user_service

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	ActivateUserAccount(ctx context.Context, in *ActivationRequest, opts ...grpc.CallOption) (*ActivationResponse, error)
	SendRequestForPasswordRecovery(ctx context.Context, in *PasswordRecoveryRequest, opts ...grpc.CallOption) (*PasswordRecoveryResponse, error)
	RecoverPassword(ctx context.Context, in *NewPasswordRequest, opts ...grpc.CallOption) (*NewPasswordResponse, error)
	PwnedPassword(ctx context.Context, in *PwnedRequest, opts ...grpc.CallOption) (*PwnedResponse, error)
	GenerateAPIToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*ApiToken, error)
	ShareJobOffer(ctx context.Context, in *ShareJobOfferRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*UserDetails, error)
	EditUserDetails(ctx context.Context, in *UserDetailsRequest, opts ...grpc.CallOption) (*UserDetails, error)
	EditUserPersonalDetails(ctx context.Context, in *UserPersonalDetailsRequest, opts ...grpc.CallOption) (*UserPersonalDetails, error)
	EditUserProfessionalDetails(ctx context.Context, in *UserProfessionalDetailsRequest, opts ...grpc.CallOption) (*UserProfessionalDetails, error)
	ChangeProfileStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatus, error)
	ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error)
	ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error)
	GetEmailUsername(ctx context.Context, in *EmailUsernameRequest, opts ...grpc.CallOption) (*EmailUsernameResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActivateUserAccount(ctx context.Context, in *ActivationRequest, opts ...grpc.CallOption) (*ActivationResponse, error) {
	out := new(ActivationResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/ActivateUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendRequestForPasswordRecovery(ctx context.Context, in *PasswordRecoveryRequest, opts ...grpc.CallOption) (*PasswordRecoveryResponse, error) {
	out := new(PasswordRecoveryResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/SendRequestForPasswordRecovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RecoverPassword(ctx context.Context, in *NewPasswordRequest, opts ...grpc.CallOption) (*NewPasswordResponse, error) {
	out := new(NewPasswordResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/RecoverPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PwnedPassword(ctx context.Context, in *PwnedRequest, opts ...grpc.CallOption) (*PwnedResponse, error) {
	out := new(PwnedResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/PwnedPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenerateAPIToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*ApiToken, error) {
	out := new(ApiToken)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GenerateAPIToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ShareJobOffer(ctx context.Context, in *ShareJobOfferRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/user_service.UserService/ShareJobOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GetUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUserDetails(ctx context.Context, in *UserDetailsRequest, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/user_service.UserService/EditUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUserPersonalDetails(ctx context.Context, in *UserPersonalDetailsRequest, opts ...grpc.CallOption) (*UserPersonalDetails, error) {
	out := new(UserPersonalDetails)
	err := c.cc.Invoke(ctx, "/user_service.UserService/EditUserPersonalDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUserProfessionalDetails(ctx context.Context, in *UserProfessionalDetailsRequest, opts ...grpc.CallOption) (*UserProfessionalDetails, error) {
	out := new(UserProfessionalDetails)
	err := c.cc.Invoke(ctx, "/user_service.UserService/EditUserProfessionalDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeProfileStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatus, error) {
	out := new(ChangeStatus)
	err := c.cc.Invoke(ctx, "/user_service.UserService/ChangeProfileStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error) {
	out := new(ChangeEmailResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/ChangeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error) {
	out := new(ChangeUsernameResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/ChangeUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetEmailUsername(ctx context.Context, in *EmailUsernameRequest, opts ...grpc.CallOption) (*EmailUsernameResponse, error) {
	out := new(EmailUsernameResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GetEmailUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetAll(context.Context, *EmptyRequest) (*GetAllResponse, error)
	UpdateUser(context.Context, *UpdateRequest) (*UpdateUserResponse, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	ActivateUserAccount(context.Context, *ActivationRequest) (*ActivationResponse, error)
	SendRequestForPasswordRecovery(context.Context, *PasswordRecoveryRequest) (*PasswordRecoveryResponse, error)
	RecoverPassword(context.Context, *NewPasswordRequest) (*NewPasswordResponse, error)
	PwnedPassword(context.Context, *PwnedRequest) (*PwnedResponse, error)
	GenerateAPIToken(context.Context, *GenerateTokenRequest) (*ApiToken, error)
	ShareJobOffer(context.Context, *ShareJobOfferRequest) (*EmptyRequest, error)
	GetUserDetails(context.Context, *GetUserDetailsRequest) (*UserDetails, error)
	EditUserDetails(context.Context, *UserDetailsRequest) (*UserDetails, error)
	EditUserPersonalDetails(context.Context, *UserPersonalDetailsRequest) (*UserPersonalDetails, error)
	EditUserProfessionalDetails(context.Context, *UserProfessionalDetailsRequest) (*UserProfessionalDetails, error)
	ChangeProfileStatus(context.Context, *ChangeStatusRequest) (*ChangeStatus, error)
	ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error)
	ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error)
	GetEmailUsername(context.Context, *EmailUsernameRequest) (*EmailUsernameResponse, error)
	MustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetAll(context.Context, *EmptyRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) ActivateUserAccount(context.Context, *ActivationRequest) (*ActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUserAccount not implemented")
}
func (UnimplementedUserServiceServer) SendRequestForPasswordRecovery(context.Context, *PasswordRecoveryRequest) (*PasswordRecoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequestForPasswordRecovery not implemented")
}
func (UnimplementedUserServiceServer) RecoverPassword(context.Context, *NewPasswordRequest) (*NewPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverPassword not implemented")
}
func (UnimplementedUserServiceServer) PwnedPassword(context.Context, *PwnedRequest) (*PwnedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PwnedPassword not implemented")
}
func (UnimplementedUserServiceServer) GenerateAPIToken(context.Context, *GenerateTokenRequest) (*ApiToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAPIToken not implemented")
}
func (UnimplementedUserServiceServer) ShareJobOffer(context.Context, *ShareJobOfferRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareJobOffer not implemented")
}
func (UnimplementedUserServiceServer) GetUserDetails(context.Context, *GetUserDetailsRequest) (*UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedUserServiceServer) EditUserDetails(context.Context, *UserDetailsRequest) (*UserDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserDetails not implemented")
}
func (UnimplementedUserServiceServer) EditUserPersonalDetails(context.Context, *UserPersonalDetailsRequest) (*UserPersonalDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserPersonalDetails not implemented")
}
func (UnimplementedUserServiceServer) EditUserProfessionalDetails(context.Context, *UserProfessionalDetailsRequest) (*UserProfessionalDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserProfessionalDetails not implemented")
}
func (UnimplementedUserServiceServer) ChangeProfileStatus(context.Context, *ChangeStatusRequest) (*ChangeStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeProfileStatus not implemented")
}
func (UnimplementedUserServiceServer) ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEmail not implemented")
}
func (UnimplementedUserServiceServer) ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUsername not implemented")
}
func (UnimplementedUserServiceServer) GetEmailUsername(context.Context, *EmailUsernameRequest) (*EmailUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailUsername not implemented")
}
func (UnimplementedUserServiceServer) MustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	MustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActivateUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivateUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/ActivateUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivateUserAccount(ctx, req.(*ActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendRequestForPasswordRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendRequestForPasswordRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/SendRequestForPasswordRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendRequestForPasswordRecovery(ctx, req.(*PasswordRecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RecoverPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RecoverPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/RecoverPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RecoverPassword(ctx, req.(*NewPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PwnedPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PwnedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PwnedPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/PwnedPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PwnedPassword(ctx, req.(*PwnedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenerateAPIToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenerateAPIToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GenerateAPIToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenerateAPIToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ShareJobOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareJobOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ShareJobOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/ShareJobOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ShareJobOffer(ctx, req.(*ShareJobOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/EditUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUserDetails(ctx, req.(*UserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUserPersonalDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPersonalDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUserPersonalDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/EditUserPersonalDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUserPersonalDetails(ctx, req.(*UserPersonalDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUserProfessionalDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfessionalDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUserProfessionalDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/EditUserProfessionalDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUserProfessionalDetails(ctx, req.(*UserProfessionalDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeProfileStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeProfileStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/ChangeProfileStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeProfileStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/ChangeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeEmail(ctx, req.(*ChangeEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/ChangeUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUsername(ctx, req.(*ChangeUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetEmailUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetEmailUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GetEmailUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetEmailUsername(ctx, req.(*EmailUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "ActivateUserAccount",
			Handler:    _UserService_ActivateUserAccount_Handler,
		},
		{
			MethodName: "SendRequestForPasswordRecovery",
			Handler:    _UserService_SendRequestForPasswordRecovery_Handler,
		},
		{
			MethodName: "RecoverPassword",
			Handler:    _UserService_RecoverPassword_Handler,
		},
		{
			MethodName: "PwnedPassword",
			Handler:    _UserService_PwnedPassword_Handler,
		},
		{
			MethodName: "GenerateAPIToken",
			Handler:    _UserService_GenerateAPIToken_Handler,
		},
		{
			MethodName: "ShareJobOffer",
			Handler:    _UserService_ShareJobOffer_Handler,
		},
		{
			MethodName: "GetUserDetails",
			Handler:    _UserService_GetUserDetails_Handler,
		},
		{
			MethodName: "EditUserDetails",
			Handler:    _UserService_EditUserDetails_Handler,
		},
		{
			MethodName: "EditUserPersonalDetails",
			Handler:    _UserService_EditUserPersonalDetails_Handler,
		},
		{
			MethodName: "EditUserProfessionalDetails",
			Handler:    _UserService_EditUserProfessionalDetails_Handler,
		},
		{
			MethodName: "ChangeProfileStatus",
			Handler:    _UserService_ChangeProfileStatus_Handler,
		},
		{
			MethodName: "ChangeEmail",
			Handler:    _UserService_ChangeEmail_Handler,
		},
		{
			MethodName: "ChangeUsername",
			Handler:    _UserService_ChangeUsername_Handler,
		},
		{
			MethodName: "GetEmailUsername",
			Handler:    _UserService_GetEmailUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
