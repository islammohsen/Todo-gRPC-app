package todo

import (
	"testing"

	"todo-app/util"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

const maxUserID = 100
const maxTodoLength = 20

//creatRandomTodo insert new Random todo with random user id and todo
func createRandomTodo(t *testing.T) *TodoItem {

	Todo := util.RandomString(util.RandomInt(1, maxTodoLength))
	item := &TodoItem{TodoID: -1, UserID: int32(util.RandomInt(1, maxUserID)), Todo: Todo}

	id, err := database.InsertTodoItem(item)
	item.TodoID = int32(id)

	require.NoError(t, err)
	require.NotZero(t, id)

	return item
}

//creatUserRandomTodo insert new Random todo with todo
func createUserRanomTodo(t *testing.T, userID int) *TodoItem {

	Todo := util.RandomString(util.RandomInt(1, maxTodoLength))
	item := &TodoItem{TodoID: -1, UserID: int32(userID), Todo: Todo}

	id, err := database.InsertTodoItem(item)
	item.TodoID = int32(id)

	require.NoError(t, err)
	require.NotZero(t, id)

	return item
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
					t.Errorf("[%q]: MyFunc() got error %v, want success", tc.desc, err)
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
			t.Errorf("[%q]: MyFunc() got success, want an error", tc.desc)
			continue
		}

		if diff := cmp.Diff(tc.wantRes, got); diff != "" {
			t.Errorf("[%q]: MyFunc() returned unexpected diff (-want, +got):\n%s", tc.desc, diff)
		}
	}
}

//TestDeleteUserTodos test deleting user todos
func TestDeleteUserTodos(t *testing.T) {

	//Truncate database
	err := database.Truncate()
	require.NoError(t, err)

	//create new todos
	todo := createRandomTodo(t)
	for i := 0; i < 9; i++ {
		createUserRanomTodo(t, int(todo.UserID))
	}

	//delete user created todo
	err = database.DeleteUserTodos(int(todo.UserID))
	require.NoError(t, err)

	//check user todos is empty
	userTodos, err := database.GetUserTodos(int(todo.UserID))
	require.NoError(t, err)
	require.Equal(t, len(userTodos), 0)
}

//TestGetUserTodos checks get user todos
func TestGetUserTodos(t *testing.T) {

	//Truncate database
	err := database.Truncate()
	require.NoError(t, err)

	//Create mark map to check all todo ids are unique
	mark := make(map[int32]*TodoItem)

	//create new todos
	todo := createRandomTodo(t)
	mark[todo.TodoID] = todo
	for i := 0; i < 9; i++ {
		currentTodo := createUserRanomTodo(t, int(todo.UserID))
		require.Empty(t, mark[currentTodo.TodoID])
		mark[currentTodo.TodoID] = currentTodo
	}

	//check user todos
	userTodos, err := database.GetUserTodos(int(todo.UserID))
	require.NoError(t, err)
	require.Equal(t, len(userTodos), 10)
	for _, currentTodo := range userTodos {
		require.NotEmpty(t, mark[currentTodo.TodoID])
		require.Equal(t, mark[currentTodo.TodoID].UserID, currentTodo.UserID)
		require.Equal(t, mark[currentTodo.TodoID].Todo, currentTodo.Todo)
		delete(mark, currentTodo.TodoID)
	}
}

//TestGetAllTodos checks geting all todos
func TestGetAllTodos(t *testing.T) {

	//Truncate database
	err := database.Truncate()
	require.NoError(t, err)

	//Create mark map to check all todo ids are unique
	mark := make(map[int32]*TodoItem)

	//create new todos
	for i := 0; i < 10; i++ {
		currentTodo := createRandomTodo(t)
		require.Empty(t, mark[currentTodo.TodoID])
		mark[currentTodo.TodoID] = currentTodo
	}

	//check todos
	userTodos, err := database.GetAllTodos()
	require.NoError(t, err)
	require.Equal(t, len(userTodos), 10)
	for _, currentTodo := range userTodos {
		require.NotEmpty(t, mark[currentTodo.TodoID])
		require.Equal(t, mark[currentTodo.TodoID].UserID, currentTodo.UserID)
		require.Equal(t, mark[currentTodo.TodoID].Todo, currentTodo.Todo)
		delete(mark, currentTodo.TodoID)
	}
}
