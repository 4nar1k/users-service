package main

import (
	"github.com/4nar1k/users-service/internal/database"
	"github.com/4nar1k/users-service/internal/transport/grpc"
	"github.com/4nar1k/users-service/internal/user"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	repo := user.NewUserRepository(db)
	service := user.NewUserService(repo)

	log.Println("Starting gRPC server on :50051")
	if err := grpc.RunGRPC(service); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
