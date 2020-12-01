package todo

import (
	"context"
	"errors"
	"testing"
	"todo-app/models"

	"github.com/google/go-cmp/cmp"
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

	for _, tc := range testData {

		ctx := context.Background()
		fakeDS := testingDB{}
		server := Server{DS: &fakeDS}

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

	for _, tc := range testData {

		ctx := context.Background()
		fakeDS := testingDB{}
		server := Server{DS: &fakeDS}

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

	for _, tc := range testData {

		ctx := context.Background()
		fakeDS := testingDB{}
		server := Server{DS: &fakeDS}

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
