package model

// Todo ID: int, Task: string
type Todo struct {
	ID   int
	Task string
}

// New Todo constructor
func New(id int, task string) *Todo {
	return &Todo{ID: id, Task: task}
}
