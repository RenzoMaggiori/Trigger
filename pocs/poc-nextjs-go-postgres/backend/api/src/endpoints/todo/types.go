package todo

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AddTodo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type UpdatedTodo struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Status      *string    `json:"status,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type Todos interface {
	FindAll() ([]Todo, error)
	FindByID(int) (*Todo, error)
	Create(*AddTodo) (*Todo, error)
	Update(int, *UpdatedTodo) (*Todo, error)
	Remove(int) error
}

type Handler struct {
	Todos
}

type Model struct {
	db *sql.DB
}
