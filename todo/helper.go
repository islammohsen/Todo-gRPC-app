package todo

import (
	"todo-app/db"
)

func toDatabaseTodoItem(item *TodoItem) *db.TodoItem {
	return &db.TodoItem{TodoID: int(item.TodoID), UserID: int(item.UserID), Todo: item.Todo}
}

func toProtoTodoItem(item *db.TodoItem) *TodoItem {
	return &TodoItem{TodoID: int32(item.TodoID), UserID: int32(item.UserID), Todo: item.Todo}
}
