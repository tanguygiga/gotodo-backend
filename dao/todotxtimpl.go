package dao

import (
	"bufio"
	"fmt"
	m "gotodo-backend/model"
	"gotodo-backend/shared"
	"log"
	"os"
	"sort"
)

type TodoMap map[int]string

// TodoTxtImpl text implementation of Todo
type TodoTxtImpl struct {
	Tm TodoMap
}

// GetAll return all Todo
func (impl TodoTxtImpl) GetAll() (listTodo []m.Todo, err error) {
	tm := impl.Tm
	if len(tm) == 0 {
		f := openFile()
		s := bufio.NewScanner(f)
		for i := 1; s.Scan(); i++ {
			tm[i] = s.Text()
		}
		if err = s.Err(); err != nil {
			return nil, err
		}
	}

	for k, v := range tm {
		var t m.Todo
		t.ID = k
		t.Task = v
		listTodo = append(listTodo, t)
	}
	return listTodo, nil
}

// Get return a Todo given an ID
func (impl TodoTxtImpl) Get(id int) (t m.Todo, err error) {
	tm := impl.Tm
	ok := false
	err = notInCollectionError(id)
	t.Task, ok = tm[id]
	if ok {
		err = nil
	}
	return t, err
}

// Create a Todo
func (impl TodoTxtImpl) Create(t *m.Todo) error {
	f := openFile()
	s := bufio.NewScanner(f)
	var list []string
	for s.Scan() {
		list = append(list, s.Text()+"\n")
	}
	list = append(list, t.Task+"\n")
	if err := s.Err(); err != nil {
		return err
	}
	err := writeSortedLines(shared.Todotxt, list)
	return err
}

// Update a Todo
func (impl TodoTxtImpl) Update(t *m.Todo) error {
	f := openFile()
	err := notInCollectionError(t.ID)
	s := bufio.NewScanner(f)
	var list []string
	for i := 1; s.Scan(); i++ {
		task := s.Text()
		if i == t.ID {
			task = t.Task
			err = nil
		}
		list = append(list, task+"\n")
	}
	if err != nil {
		return err
	}
	if err = s.Err(); err != nil {
		return err
	}
	err = writeSortedLines(shared.Todotxt, list)
	return err
}

// Delete a Todo given an id
func (impl TodoTxtImpl) Delete(id int) error {
	err := notInCollectionError(id)
	tm := impl.Tm
	_, present := tm[id]
	if present {
		delete(tm, id)
		err = nil
	}
	return err
}
func (impl TodoTxtImpl) delete(id int) error {
	err := notInCollectionError(id)
	f := openFile()
	s := bufio.NewScanner(f)
	var list []string
	for i := 1; s.Scan(); i++ {
		if i == id {
			err = nil
			continue
		}
		list = append(list, s.Text()+"\n")
	}
	if err != nil {
		return err
	}
	if err = s.Err(); err != nil {
		return err
	}
	err = writeSortedLines(shared.Todotxt, list)
	return err
}

func notInCollectionError(id int) error {
	return fmt.Errorf("id: %d not in collection", id)
}

func openFile() *os.File {
	f, err := os.Open(shared.Todotxt)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func writeSortedLines(file string, lines []string) error {
	sort.Strings(lines)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		if "\n" == line {
			continue
		}
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
