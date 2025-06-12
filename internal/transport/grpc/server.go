package grpc

import (
	protoUser "github.com/4nar1k/project-protos/proto/user"
	svcUser "github.com/4nar1k/users-service/internal/user"
	"google.golang.org/grpc"
	"net"
)

func RunGRPC(svc *svcUser.UserService) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	protoUser.RegisterUserServiceServer(grpcServer, NewUserHandler(svc))

	return grpcServer.Serve(listener)
}
