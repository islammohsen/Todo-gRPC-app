package todo

import (
	"context"
	"errors"
	"io"
	"log"
	sync "sync"
	"time"
	"todo-app/models"
)

const mod = 291391

//DataStore defining functions to be implemented to store user todos
type DataStore interface {
	InsertTodoItem(item *models.TodoItem) (int32, error)
	GetAllTodos() ([]*models.TodoItem, error)
	GetUserTodos(userID int32) ([]*models.TodoItem, error)
	DeleteUserTodos(userID int32) error
	Truncate() error
}

//Server implementing TodoSeviceServer
type Server struct {
	DS          DataStore
	WaitingTime time.Duration
}

func (s *Server) mustEmbedUnimplementedTodoServiceServer() {}

//AddTodo function to add todoitem to database
func (s *Server) AddTodo(ctx context.Context, message *AddTodoRequest) (*AddTodoResponse, error) {
	log.Printf("Received : %v", message)
	item := message.GetItem()
	id, err := s.DS.InsertTodoItem(toModelsTodoItem(item))
	if err != nil {
		return nil, err
	}
	item.TodoID = id
	return &AddTodoResponse{Item: item}, nil
}

//GetAllTodos function to get all todos from database
func (s *Server) GetAllTodos(ctx context.Context, message *NoParams) (*GetAllTodosResponse, error) {
	log.Printf("Received Get all todos request")
	response := GetAllTodosResponse{Items: make([]*TodoItem, 0)}
	todos, err := s.DS.GetAllTodos()
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		response.Items = append(response.Items, toProtoTodoItem(todo))
	}
	return &response, nil
}

//GetAllTodosStreaming function to get all todos from database
//server side streaming
func (s *Server) GetAllTodosStreaming(message *NoParams, stream TodoService_GetAllTodosStreamingServer) error {
	log.Printf("Received Get all todos streaming request")
	todos, err := s.DS.GetAllTodos()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(s.WaitingTime)
	defer ticker.Stop()
	for _, todo := range todos {
		select {
		case <-ticker.C:
			err := stream.Send(toProtoTodoItem(todo))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//GetUserTodos function to get a stream of user ids and return a stream of todoitems
func (s *Server) GetUserTodos(stream TodoService_GetUserTodosServer) error {
	log.Println("Received get user todos request")
	ticker := time.NewTicker(s.WaitingTime)
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

			dbTodos, err := s.DS.GetUserTodos(userID)
			if err != nil {
				return err
			}
			var todos []*TodoItem
			for _, todo := range dbTodos {
				todos = append(todos, toProtoTodoItem(todo))
			}
			response := &GetUserTodosResponse{Items: todos}
			log.Println("Sending", response)
			stream.Send(response)
		}
	}
}

//DeleteUserTodos input user id, delete user todos from datastore
func (s *Server) DeleteUserTodos(ctx context.Context, message *DeleteUserTodosRequest) (*DeleteUserTodosResponse, error) {
	userID := message.UserID
	err := s.DS.DeleteUserTodos(userID)
	if err != nil {
		return nil, err
	}
	return &DeleteUserTodosResponse{}, nil
}

func (s *Server) computeTodoHash(ctx context.Context, item *TodoItem) (int32, error) {
	waitingTime := s.WaitingTime / 2
	select {
	case <-time.After(waitingTime):
		return (item.TodoID + item.UserID) % mod, nil
	//context timed out or canceld
	case <-ctx.Done():
		return 0, errors.New("Canceld or Timed out")
	}
}

func parallel(ctx context.Context, list []func(context.Context) error, ch chan error) {

	childContext, cancel := context.WithCancel(ctx)
	defer cancel()

	errorChannel := make(chan error, len(list))

	for _, f := range list {
		go func(f func(context.Context) error) {
			errorChannel <- f(childContext)
		}(f)
	}

	for i := 0; i < len(list); i++ {
		err := <-errorChannel
		if err != nil {
			ch <- err
			return
		}
	}
	ch <- nil
}

func (s *Server) transformTodos(ctx context.Context, todos []*models.TodoItem, process func(context.Context, *TodoItem) (int32, error)) ([]*TodoItemWithHash, error) {

	response := make([]*TodoItemWithHash, len(todos))

	f := func(ctx context.Context, idx int, item *TodoItem) error {
		hash, err := process(ctx, item)
		response[idx] = &TodoItemWithHash{Item: item, Hash: hash}
		return err
	}

	var list []func(context.Context) error
	for idx, todo := range todos {
		idx := idx
		todo := todo
		list = append(list, func(ctx context.Context) error {
			return f(ctx, idx, toProtoTodoItem(todo))
		})
	}

	parallelChannel := make(chan error)
	go parallel(ctx, list, parallelChannel)
	err := <-parallelChannel

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserTodoItemsWithHash(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.DS.GetUserTodos(userID)

	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}

	items, err := s.transformTodos(ctx, todos, s.computeTodoHash)
	if err != nil {
		return nil, err
	}
	response.Items = items

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	default:
		return response, nil
	}
}

func (s *Server) transformTodosPointer(ctx context.Context, todos []*models.TodoItem, process func(context.Context, *TodoItem) (int32, error)) ([]*TodoItemWithHash, error) {

	response := make([]*TodoItemWithHash, len(todos))

	f := func(ctx context.Context, item *TodoItemWithHash) error {
		hash, err := process(ctx, item.Item)
		item.Hash = hash
		return err
	}

	var list []func(context.Context) error
	for idx, todo := range todos {
		idx := idx
		response[idx] = &TodoItemWithHash{Item: toProtoTodoItem(todo)}
		list = append(list, func(ctx context.Context) error {
			return f(ctx, response[idx])
		})
	}

	parallelChannel := make(chan error)
	go parallel(ctx, list, parallelChannel)
	err := <-parallelChannel

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserTodoItemsWithHashPointer(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.DS.GetUserTodos(userID)

	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}

	items, err := s.transformTodosPointer(ctx, todos, s.computeTodoHash)
	if err != nil {
		return nil, err
	}
	response.Items = items

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	default:
		return response, nil
	}
}

func (s *Server) transformTodosAppend(ctx context.Context, todos []*models.TodoItem, process func(context.Context, *TodoItem) (int32, error)) ([]*TodoItemWithHash, error) {

	mu := sync.Mutex{}
	var response []*TodoItemWithHash

	f := func(ctx context.Context, item *TodoItem) error {
		hash, err := process(ctx, item)
		mu.Lock()
		response = append(response, &TodoItemWithHash{Item: item, Hash: hash})
		mu.Unlock()
		return err
	}

	var list []func(context.Context) error
	for _, todo := range todos {
		todo := todo
		list = append(list, func(ctx context.Context) error {
			return f(ctx, toProtoTodoItem(todo))
		})
	}

	parallelChannel := make(chan error)
	go parallel(ctx, list, parallelChannel)
	err := <-parallelChannel

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserTodoItemsWithHashAppend(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.DS.GetUserTodos(userID)

	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}

	items, err := s.transformTodosAppend(ctx, todos, s.computeTodoHash)
	if err != nil {
		return nil, err
	}
	response.Items = items

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	default:
		return response, nil
	}
}

func (s *Server) transformTodosAppendPreAllocation(ctx context.Context, todos []*models.TodoItem, process func(context.Context, *TodoItem) (int32, error)) ([]*TodoItemWithHash, error) {

	mu := sync.Mutex{}
	response := make([]*TodoItemWithHash, 0, len(todos))

	f := func(ctx context.Context, item *TodoItem) error {
		hash, err := process(ctx, item)
		mu.Lock()
		response = append(response, &TodoItemWithHash{Item: item, Hash: hash})
		mu.Unlock()
		return err
	}

	var list []func(context.Context) error
	for _, todo := range todos {
		todo := todo
		list = append(list, func(ctx context.Context) error {
			return f(ctx, toProtoTodoItem(todo))
		})
	}

	parallelChannel := make(chan error)
	go parallel(ctx, list, parallelChannel)
	err := <-parallelChannel

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserTodoItemsWithHashAppendPreAllocation(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.DS.GetUserTodos(userID)

	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}

	items, err := s.transformTodosAppendPreAllocation(ctx, todos, s.computeTodoHash)
	if err != nil {
		return nil, err
	}
	response.Items = items

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	default:
		return response, nil
	}
}

func (s *Server) transformTodosAppendChannels(ctx context.Context, todos []*models.TodoItem, process func(context.Context, *TodoItem) (int32, error)) ([]*TodoItemWithHash, error) {

	todoWithHashChannel := make(chan *TodoItemWithHash, len(todos))
	var response []*TodoItemWithHash

	f := func(ctx context.Context, item *TodoItem, ch chan *TodoItemWithHash) error {
		hash, err := process(ctx, item)
		ch <- &TodoItemWithHash{Item: item, Hash: hash}
		return err
	}

	var list []func(context.Context) error
	for _, todo := range todos {
		todo := todo
		list = append(list, func(ctx context.Context) error {
			return f(ctx, toProtoTodoItem(todo), todoWithHashChannel)
		})
	}

	parallelChannel := make(chan error)
	go parallel(ctx, list, parallelChannel)

	for i := 0; i < len(todos); i++ {
		response = append(response, <-todoWithHashChannel)
	}

	err := <-parallelChannel

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserTodoItemsWithHashAppendChannels(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.DS.GetUserTodos(userID)

	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}

	items, err := s.transformTodosAppendChannels(ctx, todos, s.computeTodoHash)
	if err != nil {
		return nil, err
	}
	response.Items = items

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	default:
		return response, nil
	}
}

func (s *Server) transformTodosIndexingChannels(ctx context.Context, todos []*models.TodoItem, process func(context.Context, *TodoItem) (int32, error)) ([]*TodoItemWithHash, error) {

	length := len(todos)
	todoWithHashChannel := make(chan *TodoItemWithHash, length)
	response := make([]*TodoItemWithHash, length)

	f := func(ctx context.Context, item *TodoItem, ch chan *TodoItemWithHash) error {
		hash, err := process(ctx, item)
		ch <- &TodoItemWithHash{Item: item, Hash: hash}
		return err
	}

	var list []func(context.Context) error
	for _, todo := range todos {
		todo := todo
		list = append(list, func(ctx context.Context) error {
			return f(ctx, toProtoTodoItem(todo), todoWithHashChannel)
		})
	}

	parallelChannel := make(chan error)
	go parallel(ctx, list, parallelChannel)

	for i := 0; i < length; i++ {
		response[i] = <-todoWithHashChannel
	}

	err := <-parallelChannel

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserTodoItemsWithHashIndexingChannels(ctx context.Context, message *GetUserTodoItemsWithHashRequest) (*GetUserTodoItemsWithHashResponse, error) {

	log.Println("Received get user todo items with hash request", message)

	userID := message.UserID
	todos, err := s.DS.GetUserTodos(userID)

	if err != nil {
		return nil, err
	}

	response := &GetUserTodoItemsWithHashResponse{}

	items, err := s.transformTodosIndexingChannels(ctx, todos, s.computeTodoHash)
	if err != nil {
		return nil, err
	}
	response.Items = items

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	default:
		return response, nil
	}
}
