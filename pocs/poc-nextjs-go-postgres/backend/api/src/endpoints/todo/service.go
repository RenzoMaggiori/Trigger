package todo

import (
	"fmt"
	"reflect"
	"strings"
)

var _ Todos = Model{}

func (m Model) FindAll() ([]Todo, error) {
	rows, err := m.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]Todo, 0)
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (m Model) FindByID(id int) (*Todo, error) {
	query := "SELECT * FROM todo WHERE id = $1"
	row := m.db.QueryRow(query, id)

	var t Todo
	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m Model) Create(nt *AddTodo) (*Todo, error) {
	query := `
		INSERT INTO todo (title, description, status, due_date)
		VALUES ($1, $2, $3, $4)
		RETURNING *
    `
	row := m.db.QueryRow(query, nt.Title, nt.Description, nt.Status, nt.DueDate)

	var t Todo
	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m Model) Update(id int, ut *UpdatedTodo) (*Todo, error) {
	value := reflect.ValueOf(ut).Elem()
	tValue := reflect.TypeOf(*ut)

	var setClauses []string
	var args []interface{}
	argIndex := 1

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.IsNil() {
			continue
		}

		fieldName := tValue.Field(i).Tag.Get("db")
		if fieldName == "" {
			fieldName = strings.ToLower(strings.Replace(tValue.Field(i).Name, "_", "", -1))
		}

		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", fieldName, argIndex))
		args = append(args, field.Interface())
		argIndex++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf(
		"UPDATE todo SET %s, updated_at = NOW() WHERE id = $%d RETURNING *",
		strings.Join(setClauses, ", "),
		argIndex,
	)
	args = append(args, id)
	row := m.db.QueryRow(query, args...)

	var t Todo
	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m Model) Remove(id int) error {
	query := "DELETE FROM todo WHERE id = $1"
	_, err := m.db.Exec(query, id)
	return err
}
