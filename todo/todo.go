package todo

import (
	context "context"
	"io"
	"log"
	"time"
)

type Server struct {
	Todos []*TodoItem
}

func (s *Server) mustEmbedUnimplementedTodoServiceServer() {}

func (s *Server) AddTodo(ctx context.Context, message *AddTodoRequest) (*AddTodoResponse, error) {
	log.Printf("Received : %v", message)
	item := message.GetItem()
	item.TodoID = int32(len(s.Todos) + 1)
	s.Todos = append(s.Todos, item)
	return &AddTodoResponse{Item: item}, nil
}

func (s *Server) GetAllTodos(ctx context.Context, message *NoParams) (*GetAllTodosResponse, error) {
	log.Printf("Received Get all todos request")
	response := GetAllTodosResponse{Items: make([]*TodoItem, 0)}
	for _, todo := range s.Todos {
		response.Items = append(response.Items, todo)
	}
	return &response, nil
}

func (s *Server) GetAllTodosStreaming(message *NoParams, stream TodoService_GetAllTodosStreamingServer) error {
	log.Printf("Received Get all todos streaming request")
	for _, todo := range s.Todos {
		select {
		case <-time.NewTicker(time.Second).C:
			err := stream.Send(todo)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Server) CountingTest(stream TodoService_CountingTestServer) error {
	log.Println("Start counting")
	for {
		val, err := stream.Recv()
		if err == io.EOF {
			log.Println("Ended counting")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving from client %s", err)
			return err
		}
		log.Println("Received ", val)
		select {
		case <-time.NewTicker(time.Second).C:
			log.Println("Sending ", (*val).Counter+1)
			stream.Send(&Counter{Counter: (*val).Counter + 1})
		}
	}
}

func (s *Server) GetUserTodos(stream TodoService_GetUserTodosServer) error {
	log.Println("Received get user todos request")
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			log.Println("Finished get user todos request")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving from client %s", err)
			return err
		}
		log.Println("Received", message)
		userID := message.UserID
		select {
		case <-time.NewTicker(time.Second).C:
			response := &GetUserTodosResponse{Items: make([]*TodoItem, 0)}
			for _, todo := range s.Todos {
				if todo.UserID == userID {
					response.Items = append(response.Items, todo)
				}
			}
			log.Println("Sending", response)
			stream.Send(response)
		}
	}
}
