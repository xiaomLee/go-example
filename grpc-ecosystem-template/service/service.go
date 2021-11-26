package service

import (
	"context"
	"grpc-ecosystem-template/api"
	"grpc-ecosystem-template/internal/logic"
	"log"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
			"request_id": "request id",
			"user_id":    "user id",
		},
	))
	return &api.Response{Message: "system is running"}, nil
	//return nil, status.Errorf(codes.ResourceExhausted, "internal server error")
}

func (UserService) UserCreate(ctx context.Context, req *api.UserCreateRequest) (*api.UserCreateResponse, error) {
	if err := logic.CreateUser(ctx, req.User); err != nil {
		logrus.Errorf("service create user err:%s", err.Error())
		return nil, status.Errorf(codes.Internal, "service create user err:%s", err.Error())
	}
	return &api.UserCreateResponse{
		Code: 0,
		User: req.User,
	}, nil
}

func (UserService) UserDelete(ctx context.Context, req *api.UserDeleteRequest) (*api.UserDeleteResponse, error) {
	if err := logic.DeleteUser(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.UserDeleteResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (UserService) UserGet(ctx context.Context, req *api.UserGetRequest) (*api.UserGetResponse, error) {
	user, err := logic.GetUser(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.UserGetResponse{
		Code: 0,
		User: user,
	}, nil
}

func (UserService) UserList(ctx context.Context, req *api.UserListRequest) (*api.UserListResponse, error) {
	if req.StartTime > 0 && req.EndTime > 0 && req.StartTime > req.EndTime {
		return nil, status.Errorf(codes.InvalidArgument, "startTime:%d > endTime:%d", req.StartTime, req.EndTime)
	}
	params := make(map[string]interface{})
	if req.Id > 0 {
		params["id"] = req.Id
	}
	if req.Name != "" {
		params["name"] = req.Name
	}
	if len(req.Gender) > 0 {
		gender := make([]int32, 0)
		for _, g := range req.Gender {
			gender = append(gender, int32(g))
		}
		params["gender"] = gender
	}
	if len(req.Status) > 0 {
		status := make([]int32, 0)
		for _, g := range req.Gender {
			status = append(status, int32(g))
		}
		params["status"] = status
	}
	if req.Email != "" {
		params["email"] = req.Email
	}
	if req.StartTime > 0 {
		params["start_time"] = req.StartTime
	}
	if req.EndTime > 0 {
		params["end_time"] = req.EndTime
	}
	list, err := logic.ListUser(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &api.UserListResponse{
		Code:  0,
		Items: list,
	}, nil
}
func (UserService) Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
