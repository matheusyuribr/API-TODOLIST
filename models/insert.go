package models

import (
	"api-postgresql/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := "INSERT INTO todos (tittle, description, done) VALUES ($1, $2, $3) RETURNING id"

	err = conn.QueryRow(sql, todo.Tittle, todo.Description, todo.Done).Scan(&id)

	return
}
