package repository

import (
	"github.com/jmoiron/sqlx"
	"server"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	//sqlStr := "INSERT INTO users(name, username, password_hash) values ($1, $2, $3) RETURNING id"
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	//row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	row := r.db.QueryRow("INSERT INTO users(name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	//query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2")

	err := r.db.Get(&user, "SELECT id FROM users WHERE username=$1 AND password_hash=$2", username, password)
	//if err !=nil {
	//	return nil, err
	//}
	return user, err
}