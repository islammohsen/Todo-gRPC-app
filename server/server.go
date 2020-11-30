package main

import (
	"log"
	"net"
	"todo-app/db"
	"todo-app/todo"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Printf("Failed to listen to port 9009 : %v", err)
		return
	}

	database, err := db.GetDB("testdb")
	if err != nil {
		log.Printf("Error when connecting to database : %v", err)
		return
	}
	s := todo.Server{Database: database}

	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to serve gRPC server over port 9000 : %v", err)
	}
}
