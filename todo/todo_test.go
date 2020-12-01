package todo

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"
	"time"
	"todo-app/models"

	"github.com/google/go-cmp/cmp"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/testing/protocmp"
)

type testingDB struct {
	intResp   int32
	todosResp []*models.TodoItem
	err       error
	data      []*models.TodoItem
}

func (this *testingDB) InsertTodoItem(item *models.TodoItem) (int32, error) {
	return this.intResp, this.err
}
func (this *testingDB) GetAllTodos() ([]*models.TodoItem, error) {
	return this.todosResp, this.err
}
func (this *testingDB) GetUserTodos(userID int32) ([]*models.TodoItem, error) {
	return this.todosResp, this.err
}
func (this *testingDB) DeleteUserTodos(userID int32) error {
	if this.err != nil {
		return this.err
	}
	for i := 0; i < len(this.data); i++ {
		if this.data[i].UserID == userID {
			//delete index i
			copy(this.data[i:], this.data[i+1:])     // Shift a[i+1:] left one index.
			this.data = this.data[:len(this.data)-1] // Truncate slice
			i--
		}
	}
	return this.err
}
func (this *testingDB) Truncate() error {
	return this.err
}

const testingWaitingTime = 10 * time.Millisecond

func TestAddTodo(t *testing.T) {
	testData := []struct {
		desc    string
		input   *AddTodoRequest
		dsResp  int32
		dsErr   error
		wantRes *AddTodoResponse
		wantErr bool
	}{
		{
			desc:    "Test AddTodo returning todo item",
			input:   &AddTodoRequest{Item: &TodoItem{UserID: 1, TodoID: -1, Todo: "Task 1"}},
			dsResp:  1,
			dsErr:   nil,
			wantRes: &AddTodoResponse{Item: &TodoItem{UserID: 1, TodoID: 1, Todo: "Task 1"}},
			wantErr: false,
		},
		{
			desc:    "Test AddTodo returning error",
			input:   &AddTodoRequest{Item: &TodoItem{UserID: 1, TodoID: -1, Todo: "Task 1"}},
			dsResp:  0,
			dsErr:   errors.New("Invalid"),
			wantRes: nil,
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tc := range testData {

		fakeDS := testingDB{}
		server := Server{DS: &fakeDS, WaitingTime: testingWaitingTime}

		fakeDS.intResp = tc.dsResp
		fakeDS.err = tc.dsErr

		got, err := server.AddTodo(ctx, tc.input)

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: AddTodo() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: AddTodo() got error %v, want success", tc.desc, err)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, got, protocmp.Transform()); diff != "" {
			t.Errorf("[%q]: AddTodo() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}
	}
}

func TestGetAllTodos(t *testing.T) {
	testData := []struct {
		desc    string
		input   *NoParams
		dsResp  []*models.TodoItem
		dsErr   error
		wantRes *GetAllTodosResponse
		wantErr bool
	}{
		{
			desc:   "Empty response",
			input:  &NoParams{},
			dsResp: []*models.TodoItem{},
			dsErr:  nil,
			wantRes: &GetAllTodosResponse{
				Items: []*TodoItem{},
			},
			wantErr: false,
		},
		{
			desc:  "one todo item",
			input: &NoParams{},
			dsResp: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			dsErr: nil,
			wantRes: &GetAllTodosResponse{
				Items: []*TodoItem{
					&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				},
			},
			wantErr: false,
		},
		{
			desc:  "multiple todo items",
			input: &NoParams{},
			dsResp: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 3, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
			},
			dsErr: nil,
			wantRes: &GetAllTodosResponse{
				Items: []*TodoItem{
					&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
					&TodoItem{TodoID: 2, UserID: 1, Todo: "Task 2"},
					&TodoItem{TodoID: 3, UserID: 2, Todo: "Task 1"},
					&TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				},
			},
			wantErr: false,
		},
		{
			desc:    "Error response",
			input:   &NoParams{},
			dsResp:  nil,
			dsErr:   errors.New("Invalid"),
			wantRes: nil,
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tc := range testData {

		fakeDS := testingDB{}
		server := Server{DS: &fakeDS, WaitingTime: testingWaitingTime}

		fakeDS.todosResp = tc.dsResp
		fakeDS.err = tc.dsErr

		got, err := server.GetAllTodos(ctx, tc.input)

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: GetAllTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: GetAllTodos() got error %v, want success", tc.desc, err)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, got, protocmp.Transform()); diff != "" {
			t.Errorf("[%q]: GetAllTodos() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}
	}
}

func TestDeleteUserTodos(t *testing.T) {
	testData := []struct {
		desc       string
		input      *DeleteUserTodosRequest
		dsData     []*models.TodoItem
		dsErr      error
		wantRes    *DeleteUserTodosResponse
		wantDsData []*models.TodoItem
		wantErr    bool
	}{
		{
			desc:  "Delete no items",
			input: &DeleteUserTodosRequest{UserID: 4},
			dsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			dsErr:   nil,
			wantRes: &DeleteUserTodosResponse{},
			wantDsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			wantErr: false,
		},
		{
			desc:  "Delete one items",
			input: &DeleteUserTodosRequest{UserID: 3},
			dsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			dsErr:   nil,
			wantRes: &DeleteUserTodosResponse{},
			wantDsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			wantErr: false,
		},
		{
			desc:  "Delete multiple items",
			input: &DeleteUserTodosRequest{UserID: 1},
			dsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			dsErr:   nil,
			wantRes: &DeleteUserTodosResponse{},
			wantDsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
			},
			wantErr: false,
		},
		{
			desc:  "no delete error",
			input: &DeleteUserTodosRequest{UserID: 4},
			dsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			dsErr:   errors.New("Invalid"),
			wantRes: &DeleteUserTodosResponse{},
			wantDsData: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 5, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tc := range testData {

		fakeDS := testingDB{}
		server := Server{DS: &fakeDS, WaitingTime: testingWaitingTime}

		fakeDS.data = tc.dsData
		fakeDS.err = tc.dsErr

		got, err := server.DeleteUserTodos(ctx, tc.input)

		if diff := cmp.Diff(tc.wantDsData, fakeDS.data); diff != "" {
			t.Errorf("[%q]: DeleteUserTodos() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: DeleteUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: DeleteUserTodos() got error %v, want success", tc.desc, err)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, got, protocmp.Transform()); diff != "" {
			t.Errorf("[%q]: DeleteUserTodos() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}
	}
}

func dialer(fakeServer *Server) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterTodoServiceServer(server, fakeServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestGetAllTodosStreaming(t *testing.T) {
	testData := []struct {
		desc    string
		input   *NoParams
		dsResp  []*models.TodoItem
		dsErr   error
		wantRes []*TodoItem
		wantErr bool
	}{
		{
			desc:    "Empty response",
			input:   &NoParams{},
			dsResp:  []*models.TodoItem{},
			dsErr:   nil,
			wantRes: nil,
			wantErr: false,
		},
		{
			desc:  "one todo item",
			input: &NoParams{},
			dsResp: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			dsErr: nil,
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc:  "multiple todo items",
			input: &NoParams{},
			dsResp: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 3, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
			},
			dsErr: nil,
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: 2, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: 3, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc:    "Error response",
			input:   &NoParams{},
			dsResp:  nil,
			dsErr:   errors.New("Invalid"),
			wantRes: nil,
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tc := range testData {

		fakeDS := testingDB{}
		server := Server{DS: &fakeDS, WaitingTime: testingWaitingTime}

		fakeDS.todosResp = tc.dsResp
		fakeDS.err = tc.dsErr

		//init server
		conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(&server)))
		if err != nil {
			t.Errorf("[%q]: GetAllTodosStreaming() got error %v", tc.desc, err)
			return
		}
		defer conn.Close()

		//init client
		client := NewTodoServiceClient(conn)

		stream, err := client.GetAllTodosStreaming(ctx, tc.input)
		if err != nil {
			t.Errorf("[%q]: GetAllTodosStreaming() got error %v", tc.desc, err)
			continue
		}

		err = nil
		var todos []*TodoItem
		for {
			item, curErr := stream.Recv()
			if curErr == io.EOF {
				break
			}
			if curErr != nil {
				err = curErr
				break
			}
			todos = append(todos, item)
		}

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: DeleteUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: DeleteUserTodos() got error %v, want success", tc.desc, err)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, todos, protocmp.Transform()); diff != "" {
			t.Errorf("[%q]: GetAllTodosStreaming() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}
	}

}

type testing_TodoService_GetAllTodosStreamingServer struct {
	grpc.ServerStream
	Results []*TodoItem
}

func (this *testing_TodoService_GetAllTodosStreamingServer) Send(item *TodoItem) error {
	this.Results = append(this.Results, item)
	return nil
}

func TestGetAllTodosStreaming2(t *testing.T) {
	testData := []struct {
		desc    string
		input   *NoParams
		dsResp  []*models.TodoItem
		dsErr   error
		wantRes []*TodoItem
		wantErr bool
	}{
		{
			desc:    "Empty response",
			input:   &NoParams{},
			dsResp:  []*models.TodoItem{},
			dsErr:   nil,
			wantRes: nil,
			wantErr: false,
		},
		{
			desc:  "one todo item",
			input: &NoParams{},
			dsResp: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			dsErr: nil,
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc:  "multiple todo items",
			input: &NoParams{},
			dsResp: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 3, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
			},
			dsErr: nil,
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: 2, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: 3, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: 4, UserID: 3, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc:    "Error response",
			input:   &NoParams{},
			dsResp:  nil,
			dsErr:   errors.New("Invalid"),
			wantRes: nil,
			wantErr: true,
		},
	}

	for _, tc := range testData {

		fakeDS := testingDB{}
		server := Server{DS: &fakeDS, WaitingTime: testingWaitingTime}

		fakeDS.todosResp = tc.dsResp
		fakeDS.err = tc.dsErr

		stream := &testing_TodoService_GetAllTodosStreamingServer{}
		err := server.GetAllTodosStreaming(tc.input, stream)

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: DeleteUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: DeleteUserTodos() got error %v, want success", tc.desc, err)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, stream.Results, protocmp.Transform()); diff != "" {
			t.Errorf("[%q]: GetAllTodosStreaming() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}
	}

}
