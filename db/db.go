package db

import (
	"database/sql"
	"todo-app/models"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func GetDB(dbName string) (*Database, error) {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/"+dbName)

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (this *Database) InsertTodoItem(item *models.TodoItem) (int32, error) {
	const query = "INSERT INTO todos (UserID, Todo) VALUES(?, ?);"
	result, err := this.db.Exec(query, item.UserID, item.Todo)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int32(id), err
}

func extractTodos(rows *sql.Rows) ([]*models.TodoItem, error) {
	defer rows.Close()
	todos := make([]*models.TodoItem, 0)
	for rows.Next() {
		item := &models.TodoItem{}
		err := rows.Scan(&item.TodoID, &item.UserID, &item.Todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, item)
	}
	return todos, nil
}

func (this *Database) GetAllTodos() ([]*models.TodoItem, error) {
	const query = "SELECT * FROM todos"
	rows, err := this.db.Query(query)

	if err != nil {
		return nil, err
	}

	return extractTodos(rows)
}

func (this *Database) GetUserTodos(userID int32) ([]*models.TodoItem, error) {
	const query = "SELECT * FROM todos WHERE UserID = ?"
	rows, err := this.db.Query(query, userID)

	if err != nil {
		return nil, err
	}

	return extractTodos(rows)
}

func (this *Database) DeleteUserTodos(userID int32) error {
	const query = "DELETE FROM todos WHERE UserID = ?"
	_, err := this.db.Exec(query, userID)
	return err
}

func (this *Database) Truncate() error {
	const query = "TRUNCATE TABLE todos;"
	_, err := this.db.Exec(query)
	return err
}
