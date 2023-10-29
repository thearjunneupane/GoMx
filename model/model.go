package model

import (
	"gomx/db"
)

type Todo struct {
	ID   uint64
	Todo string
	Done bool
}

func CreateTodo(todo string) error {

	stmt := `INSERT INTO todos(todo, done) VALUES($1, $2)`

	_, err := db.DB.Exec(stmt, todo, false)
	return err
}

func GetAllTodos() ([]Todo, error) {
	todos := []Todo{}
	stmt := `SELECT id, todo, done FROM todos;`

	rows, err := db.DB.Query(stmt)
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id    uint64
			title string
			done  bool
		)

		if err := rows.Scan(&id, &title, &done); err != nil {
			return todos, err
		}

		todo := Todo{
			ID:   id,
			Todo: title,
			Done: done,
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func GetTodo(id uint64) (Todo, error) {
	stmt := `SELECT todo, done FROM todos WHERE id=$1`

	todo := Todo{}

	todo.ID = id

	rows, err := db.DB.Query(stmt, id)
	if err != nil {
		return todo, err
	}

	defer rows.Close()

	for rows.Next() {
		var title string
		var done bool

		err := rows.Scan(&title, &done)
		if err != nil {
			return todo, err
		}

		todo.Todo = title
		todo.Done = done

	}

	return todo, err

}

func MarkTodo(id uint64) (bool, error) {
	todo, err := GetTodo(id)
	if err != nil {
		return !todo.Done, err
	}

	stmt := `UPDATE todos SET done=$2 WHERE id=$1`

	_, err = db.DB.Query(stmt, id, !todo.Done)

	return !todo.Done, err
}

func DeleteTodo(id uint64) error {
	stmt := `DELETE FROM todos WHERE id=$1`

	_, err := db.DB.Exec(stmt, id)
	return err
}

func EditTodo(id uint64, newtodo string) error {

	oldtodo, err := GetTodo(id)
	if err != nil {
		return err
	}

	stmt := `UPDATE todos SET todo = REPLACE(todo, $2, $3) WHERE id=$1;`

	_, err = db.DB.Query(stmt, id, oldtodo.Todo, newtodo)

	return err
}
