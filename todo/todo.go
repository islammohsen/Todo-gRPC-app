package todo

import (
	"context"
	"io"
	"log"
	"time"
)

//Server implementing TodoSeviceServer
type Server struct {
	Database *Database
}

//AddTodo function to add todoitem to database
func (s *Server) AddTodo(ctx context.Context, message *AddTodoRequest) (*AddTodoResponse, error) {
	log.Printf("Received : %v", message)
	item := message.GetItem()
	id, err := s.Database.InsertTodoItem(item)
	if err != nil {
		return nil, err
	}
	item.TodoID = int32(id)
	return &AddTodoResponse{Item: item}, nil
}

//GetAllTodos function to get all todos from database
func (s *Server) GetAllTodos(ctx context.Context, message *NoParams) (*GetAllTodosResponse, error) {
	log.Printf("Received Get all todos request")
	response := GetAllTodosResponse{Items: make([]*TodoItem, 0)}
	todos, err := s.Database.GetAllTodos()
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		response.Items = append(response.Items, todo)
	}
	return &response, nil
}

//GetAllTodosStreaming function to get all todos from database
//server side streaming
func (s *Server) GetAllTodosStreaming(message *NoParams, stream TodoService_GetAllTodosStreamingServer) error {
	log.Printf("Received Get all todos streaming request")
	todos, err := s.Database.GetAllTodos()
	if err != nil {
		return err
	}
	for _, todo := range todos {
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

//CountingTest function to test bi-directional server client streaming
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

//GetUserTodos function to get a stream of user ids and return a stream of todoitems
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

			todos, err := s.Database.GetUserTodos(int(userID))
			if err != nil {
				return err
			}
			response := &GetUserTodosResponse{Items: todos}
			log.Println("Sending", response)
			stream.Send(response)
		}
	}
}

func (s *Server) mustEmbedUnimplementedTodoServiceServer() {}
