package models

import "api-postgresql/db"

func Get(id int64) (todo Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	err = row.Scan(&todo.ID, &todo.Tittle, &todo.Description, &todo.Done)

	return

}
