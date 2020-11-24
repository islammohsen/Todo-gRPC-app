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
		log.Fatalf("Failed to listen to port 9009 : %v", err)
	}
	s := todo.Server{}
	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 : %v", err)
	}
}
