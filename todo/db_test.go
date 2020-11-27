package todo

import (
	"log"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func setup(t *testing.T, initialTodos []*TodoItem) {
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
		env     []*TodoItem
		input   []*TodoItem
		wantRes []int
		wantErr bool
	}{
		{
			desc: "empty initial env - one input",
			env:  []*TodoItem{},
			input: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
			},
			wantRes: []int{
				1,
			},
			wantErr: false,
		},
		{
			desc: "empty initial env - multiple input",
			env:  []*TodoItem{},
			input: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
			},
			wantRes: []int{
				1,
				2,
			},
			wantErr: false,
		},
		{
			desc: "populated initial env - one input",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
			},
			wantRes: []int{
				4,
			},
			wantErr: false,
		},
		{
			desc: "populated initial env - multiple input",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
			},
			wantRes: []int{
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
		var got []int

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
		env     []*TodoItem
		input   int
		wantErr bool
	}{
		{
			desc: "no todos for user",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantErr: false,
		},
		{
			desc: "one todo for user",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantErr: false,
		},
		{
			desc: "multiple todo for user",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
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
		env     []*TodoItem
		input   int
		wantRes []*TodoItem
		wantErr bool
	}{
		{
			desc: "No user todos",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input:   1,
			wantRes: []*TodoItem{},
			wantErr: false,
		},
		{
			desc: "one user todo",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
			},
			input: 1,
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc: "multiple user todos",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 3"},
			},
			input: 1,
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: 4, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
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

		if cmp.Equal(len(tc.wantRes), len(got)) == false {
			t.Errorf("[%q]: GetUserTodos() returned unexpected diff: len wantRes %d, len got %d", tc.desc, len(tc.wantRes), len(got))
			continue
		}

		for idx := range tc.wantRes {
			if cmp.Equal(tc.wantRes[idx].TodoID, got[idx].TodoID) == false {
				t.Errorf("[%q]: GetUserTodos() returned unexpected diff todoID: want (%d), got (%d)",
					tc.desc, tc.wantRes[idx].TodoID, got[idx].TodoID)
				break
			}
			if cmp.Equal(tc.wantRes[idx].UserID, got[idx].UserID) == false {
				t.Errorf("[%q]: GetUserTodos() returned unexpected diff userID: want (%d), got (%d)",
					tc.desc, tc.wantRes[idx].UserID, got[idx].UserID)
				break
			}
			if cmp.Equal(tc.wantRes[idx].Todo, got[idx].Todo) == false {
				t.Errorf("[%q]: GetUserTodos() returned unexpected diff Todo: want (%v), got (%v)",
					tc.desc, tc.wantRes[idx].Todo, got[idx].Todo)
				break
			}
		}
	}
}

//TestGetAllTodos checks geting all todos
func TestGetAllTodos(t *testing.T) {

	testData := []struct {
		desc    string
		env     []*TodoItem
		wantRes []*TodoItem
		wantErr bool
	}{
		{
			desc:    "No Todos",
			env:     []*TodoItem{},
			wantRes: []*TodoItem{},
			wantErr: false,
		},
		{
			desc: "one user todo",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
			},
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
			},
			wantErr: false,
		},
		{
			desc: "multiple users todos",
			env: []*TodoItem{
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: -1, UserID: 3, Todo: "Task 1"},
				&TodoItem{TodoID: -1, UserID: 1, Todo: "Task 3"},
			},
			wantRes: []*TodoItem{
				&TodoItem{TodoID: 1, UserID: 1, Todo: "Task 1"},
				&TodoItem{TodoID: 2, UserID: 2, Todo: "Task 1"},
				&TodoItem{TodoID: 3, UserID: 2, Todo: "Task 2"},
				&TodoItem{TodoID: 4, UserID: 1, Todo: "Task 2"},
				&TodoItem{TodoID: 5, UserID: 3, Todo: "Task 1"},
				&TodoItem{TodoID: 6, UserID: 1, Todo: "Task 3"},
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

		if cmp.Equal(len(tc.wantRes), len(got)) == false {
			t.Errorf("[%q]: GetUserTodos() returned unexpected diff: len wantRes %d, len got %d", tc.desc, len(tc.wantRes), len(got))
			continue
		}

		for idx := range tc.wantRes {
			if cmp.Equal(tc.wantRes[idx].TodoID, got[idx].TodoID) == false {
				t.Errorf("[%q]: GetUserTodos() returned unexpected diff todoID: want (%d), got (%d)",
					tc.desc, tc.wantRes[idx].TodoID, got[idx].TodoID)
				break
			}
			if cmp.Equal(tc.wantRes[idx].UserID, got[idx].UserID) == false {
				t.Errorf("[%q]: GetUserTodos() returned unexpected diff userID: want (%d), got (%d)",
					tc.desc, tc.wantRes[idx].UserID, got[idx].UserID)
				break
			}
			if cmp.Equal(tc.wantRes[idx].Todo, got[idx].Todo) == false {
				t.Errorf("[%q]: GetUserTodos() returned unexpected diff Todo: want (%v), got (%v)",
					tc.desc, tc.wantRes[idx].Todo, got[idx].Todo)
				break
			}
		}
	}
}
