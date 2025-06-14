package user

import (
	"context"
	"fmt"
	"github.com/4nar1k/project-protos/proto/user"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	u := User{Email: req.Email}
	created, err := s.repo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return &user.CreateUserResponse{
		User: &user.User{Id: created.ID, Email: created.Email},
	}, nil
}
func (s *UserService) GetUserByID(ctx context.Context, id uint32) (*User, error) {
	if id == 0 {
		return nil, fmt.Errorf("user_id is required")
	}
	u, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &User{ID: u.ID, Email: u.Email}, nil
}

func (s *UserService) UpdateUser(req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	u := User{ID: req.Id, Email: req.Email}
	updated, err := s.repo.UpdateUserByID(req.Id, u)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResponse{
		User: &user.User{Id: updated.ID, Email: updated.Email},
	}, nil
}

func (s *UserService) DeleteUser(req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := s.repo.DeleteUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &user.DeleteUserResponse{Success: true}, nil
}

func (s *UserService) ListUsers(req *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}
	protoUsers := make([]*user.User, len(users))
	for i, u := range users {
		protoUsers[i] = &user.User{Id: u.ID, Email: u.Email}
	}
	return &user.ListUsersResponse{Users: protoUsers}, nil
}
