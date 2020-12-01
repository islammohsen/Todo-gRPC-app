package todo

import (
	"context"
	"os"
	"testing"
	"todo-app/models"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

type testingDB struct {
	intResp   int32
	todosResp []*models.TodoItem
	err       error
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
	return this.err
}
func (this *testingDB) Truncate() error {
	return this.err
}

var backgroundContext context.Context
var server Server
var fakeDS testingDB

func TestMain(m *testing.M) {

	backgroundContext = context.Background()

	fakeDS = testingDB{}

	server = Server{DS: &fakeDS}

	os.Exit(m.Run())
}

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
			desc:    "",
			input:   &AddTodoRequest{Item: &TodoItem{UserID: 1, TodoID: -1, Todo: "Task 1"}},
			dsResp:  1,
			dsErr:   nil,
			wantRes: &AddTodoResponse{Item: &TodoItem{UserID: 1, TodoID: 1, Todo: "Task 1"}},
			wantErr: false,
		},
		{
			desc:    "",
			input:   &AddTodoRequest{Item: &TodoItem{UserID: 2, TodoID: -1, Todo: "Task 1"}},
			dsResp:  2,
			dsErr:   nil,
			wantRes: &AddTodoResponse{Item: &TodoItem{UserID: 2, TodoID: 2, Todo: "Task 1"}},
			wantErr: false,
		},
	}

	for _, tc := range testData {

		fakeDS.intResp = tc.dsResp
		fakeDS.err = tc.dsErr

		got, err := server.AddTodo(backgroundContext, tc.input)

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
