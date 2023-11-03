package main

import (
	"context"
	"log"
	"net"

	"github.com/bete7512/go-grpc/todo"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening at %v", lis.Addr())
	service := grpc.NewServer()
	todo.RegisterTodoServiceServer(service, &TodoServiceServer{})
	if err := service.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type TodoServiceServer struct {
	todo.UnimplementedTodoServiceServer
}

func (s *TodoServiceServer) CreateTodo(ctx context.Context, req *todo.NewTodo) (*todo.Todo, error) {
	log.Printf("Received: %v", req.GetName())
	todo := todo.Todo{
		Name:        "Congratulations Brothers",
		Description: req.Description,
		Done:        req.Done,
	}
	return &todo, nil
}
