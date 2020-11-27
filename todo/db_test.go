package todo

import (
	"testing"

	"todo-app/util"

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

//TestInsertTodoItem test inserting todo to database
func TestInsertTodoItem(t *testing.T) {
	err := database.Truncate()
	require.NoError(t, err)
	createRandomTodo(t)
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
