package todo

var _ Todos = TodosModel{}

func (m TodosModel) FindAll() ([]Todo, error) {
	rows, err := m.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
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

func (m TodosModel) FindByID(int) (*Todo, error) {
	return nil, nil
}

func (m TodosModel) Create(*AddTodo) (*Todo, error) {
	return nil, nil
}

func (m TodosModel) Update(int, *UpdatedTodo) (*Todo, error) {
	return nil, nil
}

func (m TodosModel) Remove(int) error {
	return nil
}
