package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

type server struct {
	db *sqlx.DB
}

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Recieved : %s", r.GetName())
	return &pb.HelloResponse{Message: "Hello " + r.GetName() + "!"}, nil
}

// GET Users
func (s *server) GetUsers(ctx context.Context, r *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var users, err = models.GetUsers(ctx, s.db, *r)
	return &pb.GetUsersResponse{Users: users}, err
}

// CreateUser
func (s *server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var _, err = models.CreateUser(ctx, s.db, *r)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{}, nil
}

// GetUserById
func (s *server) GetUserById(ctx context.Context, r *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	var id = r.Id
	var user, err = models.GetUserById(ctx, s.db, id)
	return &pb.GetUserByIdResponse{User: user}, err
}

// DeleteUser
func (s *server) DeleteUser(ctx context.Context, r *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	var id = r.Id
	err := models.DeleteUser(ctx, s.db, id)
	if err != nil {
		log.Println(err)
		return &pb.DeleteUserResponse{IsDelete: false}, nil
	}
	return &pb.DeleteUserResponse{IsDelete: true}, nil
}

func main() {
	se := &server{}
	var err error
	se.db, err = sqlx.Open("mysql", "root:test@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, se)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %c", err)
	}
}
