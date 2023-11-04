package main

import (
	"context"
	"log"

	"github.com/bete7512/go-grpc/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(":5051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	client := todo.NewTodoServiceClient(conn)

	todo := todo.NewTodo{
		Name:        "Client 1",
		Description: "Well done Client 1",
		Done:        true,
	}

	todo.ProtoMessage()
	log.Println(todo.ProtoReflect())

	response, err := client.CreateTodo(context.Background(), &todo)

	if err != nil {
		log.Fatalf("could not create todo: %v", err)
	}

	log.Printf("Todo created: %v", response)

}
