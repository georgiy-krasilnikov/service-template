// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: proto/grpc-service-template.proto

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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	GetListOfUsersByIds(ctx context.Context, in *GetListOfUsersByIdsRequest, opts ...grpc.CallOption) (*GetListOfUsersByIdsResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetPostsOfUser(ctx context.Context, in *GetPostsOfUserRequest, opts ...grpc.CallOption) (*GetPostsOfUserResponse, error)
	CreateCommentForPost(ctx context.Context, in *CreateCommentForPostRequest, opts ...grpc.CallOption) (*CreateCommentForPostResponse, error)
	DeleteCommentFromPost(ctx context.Context, in *DeleteCommentFromPostRequest, opts ...grpc.CallOption) (*DeleteCommentFromPostResponse, error)
	GetPostWithComments(ctx context.Context, in *GetPostWithCommentsRequest, opts ...grpc.CallOption) (*GetPostWithCommentsResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, "/Users/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetListOfUsersByIds(ctx context.Context, in *GetListOfUsersByIdsRequest, opts ...grpc.CallOption) (*GetListOfUsersByIdsResponse, error) {
	out := new(GetListOfUsersByIdsResponse)
	err := c.cc.Invoke(ctx, "/Users/GetListOfUsersByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/Users/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetPostsOfUser(ctx context.Context, in *GetPostsOfUserRequest, opts ...grpc.CallOption) (*GetPostsOfUserResponse, error) {
	out := new(GetPostsOfUserResponse)
	err := c.cc.Invoke(ctx, "/Users/GetPostsOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateCommentForPost(ctx context.Context, in *CreateCommentForPostRequest, opts ...grpc.CallOption) (*CreateCommentForPostResponse, error) {
	out := new(CreateCommentForPostResponse)
	err := c.cc.Invoke(ctx, "/Users/CreateCommentForPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteCommentFromPost(ctx context.Context, in *DeleteCommentFromPostRequest, opts ...grpc.CallOption) (*DeleteCommentFromPostResponse, error) {
	out := new(DeleteCommentFromPostResponse)
	err := c.cc.Invoke(ctx, "/Users/DeleteCommentFromPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetPostWithComments(ctx context.Context, in *GetPostWithCommentsRequest, opts ...grpc.CallOption) (*GetPostWithCommentsResponse, error) {
	out := new(GetPostWithCommentsResponse)
	err := c.cc.Invoke(ctx, "/Users/GetPostWithComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations should embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	GetListOfUsersByIds(context.Context, *GetListOfUsersByIdsRequest) (*GetListOfUsersByIdsResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetPostsOfUser(context.Context, *GetPostsOfUserRequest) (*GetPostsOfUserResponse, error)
	CreateCommentForPost(context.Context, *CreateCommentForPostRequest) (*CreateCommentForPostResponse, error)
	DeleteCommentFromPost(context.Context, *DeleteCommentFromPostRequest) (*DeleteCommentFromPostResponse, error)
	GetPostWithComments(context.Context, *GetPostWithCommentsRequest) (*GetPostWithCommentsResponse, error)
}

// UnimplementedUsersServer should be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsersServer) GetListOfUsersByIds(context.Context, *GetListOfUsersByIdsRequest) (*GetListOfUsersByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfUsersByIds not implemented")
}
func (UnimplementedUsersServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedUsersServer) GetPostsOfUser(context.Context, *GetPostsOfUserRequest) (*GetPostsOfUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsOfUser not implemented")
}
func (UnimplementedUsersServer) CreateCommentForPost(context.Context, *CreateCommentForPostRequest) (*CreateCommentForPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommentForPost not implemented")
}
func (UnimplementedUsersServer) DeleteCommentFromPost(context.Context, *DeleteCommentFromPostRequest) (*DeleteCommentFromPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommentFromPost not implemented")
}
func (UnimplementedUsersServer) GetPostWithComments(context.Context, *GetPostWithCommentsRequest) (*GetPostWithCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostWithComments not implemented")
}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetListOfUsersByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOfUsersByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetListOfUsersByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/GetListOfUsersByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetListOfUsersByIds(ctx, req.(*GetListOfUsersByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetPostsOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsOfUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetPostsOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/GetPostsOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetPostsOfUser(ctx, req.(*GetPostsOfUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreateCommentForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentForPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateCommentForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/CreateCommentForPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateCommentForPost(ctx, req.(*CreateCommentForPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteCommentFromPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentFromPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteCommentFromPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/DeleteCommentFromPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteCommentFromPost(ctx, req.(*DeleteCommentFromPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetPostWithComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostWithCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetPostWithComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/GetPostWithComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetPostWithComments(ctx, req.(*GetPostWithCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Users_GetUserById_Handler,
		},
		{
			MethodName: "GetListOfUsersByIds",
			Handler:    _Users_GetListOfUsersByIds_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Users_CreatePost_Handler,
		},
		{
			MethodName: "GetPostsOfUser",
			Handler:    _Users_GetPostsOfUser_Handler,
		},
		{
			MethodName: "CreateCommentForPost",
			Handler:    _Users_CreateCommentForPost_Handler,
		},
		{
			MethodName: "DeleteCommentFromPost",
			Handler:    _Users_DeleteCommentFromPost_Handler,
		},
		{
			MethodName: "GetPostWithComments",
			Handler:    _Users_GetPostWithComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc-service-template.proto",
}
