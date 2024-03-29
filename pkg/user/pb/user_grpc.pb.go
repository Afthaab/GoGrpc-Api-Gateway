// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/user/pb/user.proto

package pb

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

// ProfileManagementClient is the client API for ProfileManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileManagementClient interface {
	ViewProfile(ctx context.Context, in *ViewProfileRequest, opts ...grpc.CallOption) (*ViewProfileResponse, error)
	EditProfile(ctx context.Context, in *EditProfileRequest, opts ...grpc.CallOption) (*EditProfileResponse, error)
	ChangePassword(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*ChangeResponse, error)
	AddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*AddAddressResponse, error)
	ViewAddress(ctx context.Context, in *ViewAddressRequest, opts ...grpc.CallOption) (*ViewAddressResponse, error)
	ViewAddressById(ctx context.Context, in *ViewAddressByIdRequest, opts ...grpc.CallOption) (*ViewAddressByIdResponse, error)
	EditAddress(ctx context.Context, in *EditAddressRequest, opts ...grpc.CallOption) (*EditAddressResponse, error)
	ViewAllUsers(ctx context.Context, in *ViewAllUsersRequest, opts ...grpc.CallOption) (*ViewAllUsersResponse, error)
	BlockOrUnblockUser(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
}

type profileManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileManagementClient(cc grpc.ClientConnInterface) ProfileManagementClient {
	return &profileManagementClient{cc}
}

func (c *profileManagementClient) ViewProfile(ctx context.Context, in *ViewProfileRequest, opts ...grpc.CallOption) (*ViewProfileResponse, error) {
	out := new(ViewProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/ViewProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) EditProfile(ctx context.Context, in *EditProfileRequest, opts ...grpc.CallOption) (*EditProfileResponse, error) {
	out := new(EditProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/EditProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) ChangePassword(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*ChangeResponse, error) {
	out := new(ChangeResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) AddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*AddAddressResponse, error) {
	out := new(AddAddressResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/AddAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) ViewAddress(ctx context.Context, in *ViewAddressRequest, opts ...grpc.CallOption) (*ViewAddressResponse, error) {
	out := new(ViewAddressResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/ViewAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) ViewAddressById(ctx context.Context, in *ViewAddressByIdRequest, opts ...grpc.CallOption) (*ViewAddressByIdResponse, error) {
	out := new(ViewAddressByIdResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/ViewAddressById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) EditAddress(ctx context.Context, in *EditAddressRequest, opts ...grpc.CallOption) (*EditAddressResponse, error) {
	out := new(EditAddressResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/EditAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) ViewAllUsers(ctx context.Context, in *ViewAllUsersRequest, opts ...grpc.CallOption) (*ViewAllUsersResponse, error) {
	out := new(ViewAllUsersResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/ViewAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileManagementClient) BlockOrUnblockUser(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileManagement/BlockOrUnblockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileManagementServer is the server API for ProfileManagement service.
// All implementations must embed UnimplementedProfileManagementServer
// for forward compatibility
type ProfileManagementServer interface {
	ViewProfile(context.Context, *ViewProfileRequest) (*ViewProfileResponse, error)
	EditProfile(context.Context, *EditProfileRequest) (*EditProfileResponse, error)
	ChangePassword(context.Context, *ChangeRequest) (*ChangeResponse, error)
	AddAddress(context.Context, *AddAddressRequest) (*AddAddressResponse, error)
	ViewAddress(context.Context, *ViewAddressRequest) (*ViewAddressResponse, error)
	ViewAddressById(context.Context, *ViewAddressByIdRequest) (*ViewAddressByIdResponse, error)
	EditAddress(context.Context, *EditAddressRequest) (*EditAddressResponse, error)
	ViewAllUsers(context.Context, *ViewAllUsersRequest) (*ViewAllUsersResponse, error)
	BlockOrUnblockUser(context.Context, *BlockRequest) (*BlockResponse, error)
	mustEmbedUnimplementedProfileManagementServer()
}

// UnimplementedProfileManagementServer must be embedded to have forward compatible implementations.
type UnimplementedProfileManagementServer struct {
}

func (UnimplementedProfileManagementServer) ViewProfile(context.Context, *ViewProfileRequest) (*ViewProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewProfile not implemented")
}
func (UnimplementedProfileManagementServer) EditProfile(context.Context, *EditProfileRequest) (*EditProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedProfileManagementServer) ChangePassword(context.Context, *ChangeRequest) (*ChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedProfileManagementServer) AddAddress(context.Context, *AddAddressRequest) (*AddAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddress not implemented")
}
func (UnimplementedProfileManagementServer) ViewAddress(context.Context, *ViewAddressRequest) (*ViewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAddress not implemented")
}
func (UnimplementedProfileManagementServer) ViewAddressById(context.Context, *ViewAddressByIdRequest) (*ViewAddressByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAddressById not implemented")
}
func (UnimplementedProfileManagementServer) EditAddress(context.Context, *EditAddressRequest) (*EditAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAddress not implemented")
}
func (UnimplementedProfileManagementServer) ViewAllUsers(context.Context, *ViewAllUsersRequest) (*ViewAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAllUsers not implemented")
}
func (UnimplementedProfileManagementServer) BlockOrUnblockUser(context.Context, *BlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockOrUnblockUser not implemented")
}
func (UnimplementedProfileManagementServer) mustEmbedUnimplementedProfileManagementServer() {}

// UnsafeProfileManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileManagementServer will
// result in compilation errors.
type UnsafeProfileManagementServer interface {
	mustEmbedUnimplementedProfileManagementServer()
}

func RegisterProfileManagementServer(s grpc.ServiceRegistrar, srv ProfileManagementServer) {
	s.RegisterService(&ProfileManagement_ServiceDesc, srv)
}

func _ProfileManagement_ViewProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).ViewProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/ViewProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).ViewProfile(ctx, req.(*ViewProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/EditProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).EditProfile(ctx, req.(*EditProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).ChangePassword(ctx, req.(*ChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_AddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).AddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/AddAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).AddAddress(ctx, req.(*AddAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_ViewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).ViewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/ViewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).ViewAddress(ctx, req.(*ViewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_ViewAddressById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewAddressByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).ViewAddressById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/ViewAddressById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).ViewAddressById(ctx, req.(*ViewAddressByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_EditAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).EditAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/EditAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).EditAddress(ctx, req.(*EditAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_ViewAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).ViewAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/ViewAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).ViewAllUsers(ctx, req.(*ViewAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileManagement_BlockOrUnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileManagementServer).BlockOrUnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileManagement/BlockOrUnblockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileManagementServer).BlockOrUnblockUser(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileManagement_ServiceDesc is the grpc.ServiceDesc for ProfileManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileManagement",
	HandlerType: (*ProfileManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViewProfile",
			Handler:    _ProfileManagement_ViewProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _ProfileManagement_EditProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _ProfileManagement_ChangePassword_Handler,
		},
		{
			MethodName: "AddAddress",
			Handler:    _ProfileManagement_AddAddress_Handler,
		},
		{
			MethodName: "ViewAddress",
			Handler:    _ProfileManagement_ViewAddress_Handler,
		},
		{
			MethodName: "ViewAddressById",
			Handler:    _ProfileManagement_ViewAddressById_Handler,
		},
		{
			MethodName: "EditAddress",
			Handler:    _ProfileManagement_EditAddress_Handler,
		},
		{
			MethodName: "ViewAllUsers",
			Handler:    _ProfileManagement_ViewAllUsers_Handler,
		},
		{
			MethodName: "BlockOrUnblockUser",
			Handler:    _ProfileManagement_BlockOrUnblockUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/user/pb/user.proto",
}
