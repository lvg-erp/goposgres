package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "server"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO todo_lists(title, description) VALUES ($1, $2) RETURNING id")
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	//row := r.db.QueryRow("INSERT INTO users(name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Username, user.Password)
	createUsersListQuery := fmt.Sprintf("INSERT INTO users_lists(user_id, list_id) VALUES ($1, $2) RETURNING id")
	_, err = tx.Exec(createUsersListQuery, userId, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description  " +
		"FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id WHERE ul.user_id = $1")

	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description  " +
		"FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2")

	err := r.db.Get(&list, query, userId, listId)
	return list, err
}