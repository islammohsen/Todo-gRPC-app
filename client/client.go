package main

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"todo-app/todo"

	"google.golang.org/grpc"
)

func addTodo(todoService todo.TodoServiceClient, todoString string) {
	message := todo.AddTodoRequest{Item: &todo.TodoItem{Id: -1, Todo: todoString}}
	response, err := todoService.AddTodo(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling add todo %s", err)
	}

	log.Printf("Response from server: %s", response)
}

func getAllTodos(todoService todo.TodoServiceClient) {
	todos, err := todoService.GetAllTodos(context.Background(), &todo.NoParams{})

	if err != nil {
		log.Fatalf("Error when calling get all todos %s", err)
	}

	log.Printf("Response From server: %s", todos)
}

func getAllTodosStreaming(todoService todo.TodoServiceClient) {
	stream, err := todoService.GetAllTodosStreaming(context.Background(), &todo.NoParams{}, grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("Error couldn't init stream %s", err)
	}
	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error in get all todos streaming %s", err)
		}
		log.Println("Received ", item)
	}
}

func main() {

	//get arguments
	if len(os.Args) == 1 {
		log.Fatalf("You must specify arguments")
		return
	}
	s := strings.Join(os.Args[1:], " ")

	//init connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %s", err)
	}
	defer conn.Close()

	//get todo service
	todoService := todo.NewTodoServiceClient(conn)

	//add todo
	addTodo(todoService, s)

	//get all todos
	getAllTodos(todoService)

	//get all todos streaming
	getAllTodosStreaming(todoService)
}
