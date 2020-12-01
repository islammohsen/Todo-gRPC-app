package db

import (
	"log"
	"os"
	"testing"
	"todo-app/models"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var database *Database

const dbName = "TodoTestDB"

func TestMain(m *testing.M) {
	db, err := GetDB(dbName)
	defer db.db.Close()
	if err != nil {
		log.Fatalf("cannot connect to database")
	}
	database = db
	os.Exit(m.Run())
}

func setup(t *testing.T, initialTodos []*models.TodoItem) {
	err := database.Truncate()
	if err != nil {
		t.Errorf("Error in setup database %v", err)
	}
	for _, todo := range initialTodos {
		_, err := database.InsertTodoItem(todo)
		if err != nil {
			t.Errorf("Error in setup database %v", err)
		}
	}
}

//TestInsertTodoItem test inserting todo to database
//Uses multiple inputs per test case to test the sequential behaviour of insertion(generated ids should also be sequential)
func TestInsertTodoItem(t *testing.T) {
	testData := []struct {
		desc    string
		env     []*models.TodoItem
		input   []*models.TodoItem
		wantRes []int32
		wantErr bool
	}{
		{
			desc: "empty initial env - one input",
			env:  []*models.TodoItem{},
			input: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
			},
			wantRes: []int32{
				1,
			},
			wantErr: false,
		},
		{
			desc: "empty initial env - multiple input",
			env:  []*models.TodoItem{},
			input: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
			},
			wantRes: []int32{
				1,
				2,
			},
			wantErr: false,
		},
		{
			desc: "populated initial env - one input",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
			},
			wantRes: []int32{
				4,
			},
			wantErr: false,
		},
		{
			desc: "populated initial env - multiple input",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
			},
			wantRes: []int32{
				4,
				5,
			},
			wantErr: false,
		},
	}
	for _, tc := range testData {

		//setting up enviroment
		setup(t, tc.env)

		//represent the Results
		var got []int32

		//boolean should we skip test or not
		skipTest := false

		gotError := false

		for _, todo := range tc.input {

			id, err := database.InsertTodoItem(todo)

			if err != nil {
				gotError = true
				if tc.wantErr == false {
					t.Errorf("[%q]: InsertTodoItem() got error %v, want success", tc.desc, err)
					skipTest = true
					break
				}
			}

			got = append(got, id)
		}

		//Error occured so we should skip test
		if skipTest == true {
			continue
		}

		//want error but no input produced error
		if tc.wantErr && gotError == false {
			t.Errorf("[%q]: MyFInsertTodoItemunc() got success, want an error", tc.desc)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, got); diff != "" {
			t.Errorf("[%q]: InsertTodoItem() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
		}
	}
}

//TestDeleteUserTodos test deleting user todos
func TestDeleteUserTodos(t *testing.T) {
	testData := []struct {
		desc    string
		env     []*models.TodoItem
		input   int32
		wantErr bool
	}{
		{
			desc: "no todos for user",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantErr: false,
		},
		{
			desc: "one todo for user",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantErr: false,
		},
		{
			desc: "multiple todo for user",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantErr: false,
		},
	}

	for _, tc := range testData {

		setup(t, tc.env)

		err := database.DeleteUserTodos(tc.input)

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: DeleteUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: DeleteUserTodos() got error %v, want success", tc.desc, err)
		}

		todos, err := database.GetUserTodos(tc.input)

		//Error in get user todos
		if err != nil {
			t.Errorf("[%q]: error in getting user todos (external function)", tc.desc)
			continue
		}

		if len(todos) != 0 {
			t.Errorf("[%q]: Expected all user todos to be deleted", tc.desc)
		}
	}
}

//TestGetUserTodos checks get user todos
func TestGetUserTodos(t *testing.T) {

	testData := []struct {
		desc    string
		env     []*models.TodoItem
		input   int32
		wantRes []*models.TodoItem
		wantErr bool
	}{
		{
			desc: "No user todos",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantRes: []*models.TodoItem{},
			wantErr: false,
		},
		{
			desc: "one user todo",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input: 1,
			wantRes: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc: "multiple user todos",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 3"},
			},
			input: 1,
			wantRes: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 4, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			wantErr: false,
		},
	}

	for _, tc := range testData {

		setup(t, tc.env)

		got, err := database.GetUserTodos(tc.input)

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: GetUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: GetUserTodos() got error %v, want success", tc.desc, err)
		}

		if diff := cmp.Diff(tc.wantRes, got, cmpopts.IgnoreUnexported(models.TodoItem{})); diff != "" {
			t.Errorf("[%q]: InsertTodoItem() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
		}
	}
}

//TestGetAllTodos checks geting all todos
func TestGetAllTodos(t *testing.T) {

	testData := []struct {
		desc    string
		env     []*models.TodoItem
		wantRes []*models.TodoItem
		wantErr bool
	}{
		{
			desc:    "No Todos",
			env:     []*models.TodoItem{},
			wantRes: []*models.TodoItem{},
			wantErr: false,
		},
		{
			desc: "one user todo",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
			},
			wantRes: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc: "multiple users todos",
			env: []*models.TodoItem{
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: -1, UserID: 1, Todo: "Task 3"},
			},
			wantRes: []*models.TodoItem{
				&models.TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&models.TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&models.TodoItem{TodoID: 3, UserID: 2, Todo: "Task 2"},
				&models.TodoItem{TodoID: 4, UserID: 1, Todo: "Task 2"},
				&models.TodoItem{TodoID: 5, UserID: 3, Todo: "Task 1"},
				&models.TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
			},
			wantErr: false,
		},
	}

	for _, tc := range testData {

		setup(t, tc.env)

		got, err := database.GetAllTodos()

		if tc.wantErr {
			if err == nil {
				t.Errorf("[%q]: GetUserTodos() got success, want an error", tc.desc)
			}
			continue
		}

		if err != nil {
			t.Errorf("[%q]: GetUserTodos() got error %v, want success", tc.desc, err)
		}

		if diff := cmp.Diff(tc.wantRes, got, cmpopts.IgnoreUnexported(models.TodoItem{})); diff != "" {
			t.Errorf("[%q]: GetUserTodos() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
		}
	}
}
