package todo

import (
	"todo-app/models"
)

func toModelsTodoItem(item *TodoItem) *models.TodoItem {
	return &models.TodoItem{TodoID: item.TodoID, UserID: item.UserID, Todo: item.Todo}
}

func toProtoTodoItem(item *models.TodoItem) *TodoItem {
	return &TodoItem{TodoID: item.TodoID, UserID: item.UserID, Todo: item.Todo}
}
