package grpc

import (
	"context"
	protoUser "github.com/4nar1k/project-protos/proto/user" // Тег protoUser для gRPC-пакета
	svcUser "github.com/4nar1k/users-service/internal/user" // Тег svcUser для локального сервиса
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

func (h *UserHandler) GetUser(ctx context.Context, req *protoUser.User) (*protoUser.User, error) {
	return h.service.GetUser(req)
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
