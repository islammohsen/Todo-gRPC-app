package main

import (
	"log"
	"net"
	"todo-app/todo"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Printf("Failed to listen to port 9009 : %v", err)
	}
	s := todo.Server{Todos: make([]*todo.TodoItem, 0)}
	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to serve gRPC server over port 9000 : %v", err)
	}
}
