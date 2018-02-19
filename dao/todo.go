package dao

import m "gotodo-backend/model"

// Todo Data Access Object for Todo
type Todo interface {
	GetAll() ([]m.Todo, error)
	Get(id int) (m.Todo, error)
	Create(t *m.Todo) error
	Update(t *m.Todo) error
	Delete(id int) error
}
