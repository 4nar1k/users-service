package grpc

import (
	"context"
	protoUser "github.com/4nar1k/project-protos/proto/user" // Тег protoUser для gRPC-пакета
	svcUser "github.com/4nar1k/users-service/internal/user" // Тег svcUser для локального сервиса
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	protoUser.UnimplementedUserServiceServer
	service *svcUser.UserService
}

func NewUserHandler(service *svcUser.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *protoUser.CreateUserRequest) (*protoUser.CreateUserResponse, error) {
	return h.service.CreateUser(req)
}

func (h *UserHandler) GetUser(ctx context.Context, req *protoUser.GetUserRequest) (*protoUser.User, error) {
	if req.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user_id is required")
	}
	user, err := h.service.GetUserByID(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return &protoUser.User{
		Id:    user.ID,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *protoUser.UpdateUserRequest) (*protoUser.UpdateUserResponse, error) {
	return h.service.UpdateUser(req)
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *protoUser.DeleteUserRequest) (*protoUser.DeleteUserResponse, error) {
	return h.service.DeleteUser(req)
}

func (h *UserHandler) ListUsers(ctx context.Context, req *protoUser.ListUsersRequest) (*protoUser.ListUsersResponse, error) {
	return h.service.ListUsers(req)
}
