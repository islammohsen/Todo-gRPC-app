package todo

import (
	context "context"
	"log"
	"time"
)

var todos = make([]TodoItem, 0)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedTodoServiceServer() {}

func (s *Server) AddTodo(ctx context.Context, message *AddTodoRequest) (*AddTodoResponse, error) {
	log.Printf("Received : %v", message)
	item := *message.GetItem()
	item.Id = int32(len(todos) + 1)
	todos = append(todos, item)
	return &AddTodoResponse{Item: &item}, nil
}

func (s *Server) GetAllTodos(ctx context.Context, message *NoParams) (*GetAllTodosResponse, error) {
	log.Printf("Received Get all todos request")
	response := GetAllTodosResponse{Items: make([]*TodoItem, 0)}
	for i := 0; i < len(todos); i++ {
		response.Items = append(response.Items, &todos[i])
	}
	return &response, nil
}

func (s *Server) GetAllTodosStreaming(message *NoParams, stream TodoService_GetAllTodosStreamingServer) error {
	log.Printf("Received Get all todos streaming request")
	for i := 0; i < len(todos); i++ {
		select {
		case <-time.NewTicker(time.Second).C:
			err := stream.Send(&todos[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
