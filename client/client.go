package main

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"todo-app/todo"

	"google.golang.org/grpc"
)

func addTodo(ctx context.Context, todoService todo.TodoServiceClient, userID int, todoItem string) {
	message := todo.AddTodoRequest{Item: &todo.TodoItem{UserID: int32(userID), TodoID: -1, Todo: todoItem}}
	response, err := todoService.AddTodo(ctx, &message)

	if err != nil {
		log.Printf("Error when calling add todo %s", err)
	}

	log.Printf("Response from server: %s", response)
}

func getAllTodos(ctx context.Context, todoService todo.TodoServiceClient) {
	todos, err := todoService.GetAllTodos(ctx, &todo.NoParams{})

	if err != nil {
		log.Printf("Error when calling get all todos %s", err)
	}

	log.Printf("Response From server: %s", todos)
}

func getAllTodosStreaming(ctx context.Context, todoService todo.TodoServiceClient) {
	stream, err := todoService.GetAllTodosStreaming(ctx, &todo.NoParams{}, grpc.EmptyCallOption{})
	if err != nil {
		log.Printf("Error couldn't init stream %s", err)
	}
	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error in get all todos streaming %s", err)
		}
		log.Println("Received ", item)
	}
}

func testingCounter(ctx context.Context, todoService todo.TodoServiceClient) {
	stream, err := todoService.CountingTest(ctx, grpc.EmptyCallOption{})
	if err != nil {
		log.Printf("Error couldn't init stream %s", err)
	}
	var counter int32 = 0
	for i := 0; i < 10; i++ {
		select {
		case <-time.NewTicker(time.Second).C:
			log.Println("Sending", counter+1)
			stream.Send(&todo.Counter{Counter: counter + 1})
			response, err := stream.Recv()
			if err != nil {
				log.Printf("Error in receiving")
				return
			}
			counter = response.Counter
			log.Println("Received", counter)
		}
	}
	log.Println("Closing client")
	if err := stream.CloseSend(); err != nil {
		log.Printf("Failed to close")
	}
	log.Println("Closed")
}

func getUserTodos(ctx context.Context, todoService todo.TodoServiceClient, userIDS []int) []todo.TodoItem {
	todos := make([]todo.TodoItem, 0)
	stream, err := todoService.GetUserTodos(ctx, grpc.EmptyCallOption{})
	if err != nil {
		log.Printf("Error couldn't init stream %s", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Error receiving %s", err)
				close(waitc)
				return
			}
			userTodoItems := message.Items
			log.Println("Received ", userTodoItems)
			for _, todoItem := range userTodoItems {
				todos = append(todos, *todoItem)
			}
		}
	}()
	for i := 0; i < len(userIDS); i++ {
		log.Println("Sending ", userIDS[i])
		stream.Send(&todo.GetUserTodosRequest{UserID: int32(userIDS[i])})
	}
	log.Println("Closing client")
	if err := stream.CloseSend(); err != nil {
		log.Printf("Failed to close")
	}
	log.Println("Closed")
	<-waitc
	return todos
}

func main() {

	//get arguments
	if len(os.Args) == 1 {
		log.Printf("You must specify arguments")
		return
	}

	//init connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("Could not connect %s", err)
	}
	defer conn.Close()

	//get todo service
	todoService := todo.NewTodoServiceClient(conn)

	//creating context
	ctx := context.Background()

	//add todo
	//command : !add userID todoItem
	if os.Args[1] == "!add" {
		if len(os.Args) <= 3 {
			log.Println("Invalid arguments")
			return
		}
		userID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println("User id must be a number")
			return
		}
		todoItem := strings.Join(os.Args[3:], " ")
		addTodo(ctx, todoService, userID, todoItem)
	}

	//get all todos
	if os.Args[1] == "!get_all" {
		getAllTodos(ctx, todoService)
	}

	//get all todos streaming
	if os.Args[1] == "!get_all_streaming" {
		getAllTodosStreaming(ctx, todoService)
	}

	//testing counter
	if os.Args[1] == "!counter" {
		testingCounter(ctx, todoService)
	}

	if os.Args[1] == "!get_user_todos" {
		if len(os.Args) <= 2 {
			log.Println("Invalid arguments")
			return
		}
		userIDS := make([]int, 0)
		for i := 2; i < len(os.Args); i++ {
			if val, err := strconv.Atoi(os.Args[i]); err != nil {
				log.Println("Invalid arguments")
				return
			} else {
				userIDS = append(userIDS, val)
			}
		}
		todos := getUserTodos(ctx, todoService, userIDS)
		log.Println("All user todos")
		for _, todo := range todos {
			log.Println(todo)
		}
	}
}
