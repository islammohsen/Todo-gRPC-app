package todo

import (
	"context"
	"errors"
	"io"
	"log"
	sync "sync"
	"time"
)

//Server implementing TodoSeviceServer
type Server struct {
	Database *Database
}

func (s *Server) mustEmbedUnimplementedTodoServiceServer() {}

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
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for _, todo := range todos {
		select {
		case <-ticker.C:
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
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
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
		case <-ticker.C:
			log.Println("Sending ", (*val).Counter+1)
			stream.Send(&Counter{Counter: (*val).Counter + 1})
		}
	}
}

//GetUserTodos function to get a stream of user ids and return a stream of todoitems
func (s *Server) GetUserTodos(stream TodoService_GetUserTodosServer) error {
	log.Println("Received get user todos request")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
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
		case <-ticker.C:

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

func (s *Server) DeleteUserTodos(ctx context.Context, message *DeleteUserTodosRequest) (*DeleteUserTodosResponse, error) {
	userID := message.UserID
	err := s.Database.DeleteUserTodos(int(userID))
	if err != nil {
		return nil, err
	}
	return &DeleteUserTodosResponse{}, nil
}

func computeTodoHash(ctx context.Context, item *TodoItem) (*TodoItemWithHash, error) {
	const waitingTime = 500 * time.Millisecond
	select {
	case <-time.After(waitingTime):
		hashedItem := &TodoItemWithHash{Item: item}
		hashedItem.Hash = (item.TodoID + item.UserID) % 291391
		return hashedItem, nil
	//context timed out or canceld
	case <-ctx.Done():
		return nil, errors.New("Canceld or Timed out")
	}
}

func (s *Server) GetUserTodoItemsWithHash(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.Database.GetUserTodos(int(userID))
	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}
	childContext, cancel := context.WithCancel(ctx)
	defer cancel()
	var hashingError error

	for _, todo := range todos {
		wg.Add(1)
		go func(item *TodoItem) {
			defer wg.Done()
			hashedTodo, err := computeTodoHash(childContext, item)
			if err != nil {
				mu.Lock()
				if hashingError == nil {
					hashingError = err
				}
				mu.Unlock()
				cancel()
			} else {
				mu.Lock()
				response.Items = append(response.Items, hashedTodo)
				mu.Unlock()
			}
		}(todo)
	}

	wg.Wait()

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case <-childContext.Done():
		return nil, hashingError
	default:
		return response, nil
	}
}
