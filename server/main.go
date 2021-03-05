package main

import (
	"context"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"github.com/naoyakurokawa/app-grpc-web/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":9090"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Recieved : %s", r.GetName())
	return &pb.HelloResponse{Message: "Hello " + r.GetName() + "!"}, nil
}

// GET Users
func (s *server) GetUsers(ctx context.Context, r *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var users, err = models.GetUsers(*r)
	return &pb.GetUsersResponse{Users: users}, err
}

// CreateUser
func (s *server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var _, err = models.CreateUser(*r)
	if err == nil {
		return &pb.CreateUserResponse{}, nil
	} else {
		return &pb.CreateUserResponse{}, err
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %c", err)
	}
}
