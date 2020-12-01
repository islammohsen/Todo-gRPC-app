package todo

import (
	"context"
	"log"
	"net"
	"os"
	"testing"
	"todo-app/models"

	"github.com/google/go-cmp/cmp"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/testing/protocmp"
)

type testingDB struct {
}

func (this *testingDB) InsertTodoItem(item *models.TodoItem) (int32, error)   { return 1, nil }
func (this *testingDB) GetAllTodos() ([]*models.TodoItem, error)              { return nil, nil }
func (this *testingDB) GetUserTodos(userID int32) ([]*models.TodoItem, error) { return nil, nil }
func (this *testingDB) DeleteUserTodos(userID int32) error                    { return nil }
func (this *testingDB) Truncate() error                                       { return nil }

var backgroundContext context.Context
var todoService TodoServiceClient

func TestMain(m *testing.M) {

	backgroundContext = context.Background()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen to port 9009 : %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	RegisterTodoServiceServer(grpcServer, &Server{DS: &testingDB{}})

	go grpcServer.Serve(lis)

	//init connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %s", err)
	}

	//get todo service
	todoService = NewTodoServiceClient(conn)

	os.Exit(m.Run())
}

func TestAddTodo(t *testing.T) {
	testData := []struct {
		desc    string
		input   *AddTodoRequest
		wantRes *AddTodoResponse
		wantErr bool
	}{
		{
			desc:    "",
			input:   &AddTodoRequest{Item: &TodoItem{UserID: 1, TodoID: -1, Todo: "Task 1"}},
			wantRes: &AddTodoResponse{Item: &TodoItem{UserID: 1, TodoID: 1, Todo: "Task 1"}},
			wantErr: false,
		},
	}

	for _, tc := range testData {

		got, err := todoService.AddTodo(backgroundContext, tc.input)

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: GetUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: GetUserTodos() got error %v, want success", tc.desc, err)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, got, protocmp.Transform()); diff != "" {
			t.Errorf("[%q]: GetUserTodos() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
			continue
		}
	}
}
