package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-ecosystem-template/api"
	"log"
)

type UserService struct {
	api.UnimplementedUserServiceServer
}

func (UserService) Status(ctx context.Context, req *api.Request) (*api.Response, error) {
	// get request header
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println(md)

	// send response header. can only with SetHeader/SendHeader SetTrailer
	// 客户端可以接收的 metadata 只有 header 和 trailer，因此 server 也只能发送 header 和 trailer
	grpc.SetHeader(ctx, metadata.New(
		map[string]string{
			"request_id":     "request id",
			"user_id": "user id",
		},
	))
	return &api.Response{Message: "system is running"}, nil
	//return nil, status.Errorf(codes.ResourceExhausted, "internal server error")
}

func (UserService) UserCreate(context.Context, *api.UserCreateRequest) (*api.UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UserService) UserDelete(context.Context, *api.UserDeleteRequest) (*api.UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UserService) UserGet(context.Context, *api.UserGetRequest) (*api.UserGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}
func (UserService) UserList(context.Context, *api.UserListRequest) (*api.UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UserService) Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}