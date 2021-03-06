package dao

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	m "gotodo-backend/model"
	"gotodo-backend/shared"
)

// TodoTxtImpl text implementation of Todo
type todoTxtImpl struct {
	cache map[int]string
}

func (impl *todoTxtImpl) getCache() (map[int]string, error) {
	var err error
	if len(impl.cache) == 0 {
		impl.cache = make(map[int]string)
		f := openFile()
		s := bufio.NewScanner(f)
		for i := 1; s.Scan(); i++ {
			impl.cache[i] = s.Text()
		}
		err = s.Err()
	}
	return impl.cache, err
}

// GetAll return all Todo
func (impl *todoTxtImpl) GetAll() (listTodo []m.Todo, err error) {
	cache, err := impl.getCache()
	if err != nil {
		return nil, err
	}
	for k, v := range cache {
		var t m.Todo
		t.ID = k
		t.Task = v
		listTodo = append(listTodo, t)
	}
	return listTodo, nil
}

// Get return a Todo given an id
func (impl *todoTxtImpl) Get(id int) (t m.Todo, err error) {
	cache, err := impl.getCache()
	if err != nil {
		return t, err
	}
	t.ID = id
	found := false
	t.Task, found = cache[id]
	if !found {
		err = notInCollectionError(id)
	}
	return t, err
}

// Create a Todo
func (impl *todoTxtImpl) Create(t *m.Todo) (err error) {
	cache, err := impl.getCache()
	if err != nil {
		return err
	}
	n := getLastKey(cache) + 1
	cache[n] = t.Task
	list := extractTasks(cache)
	err = writeSortedLines(shared.Todotxt, list)
	return err
}

func extractTasks(m map[int]string) (list []string) {
	for _, v := range m {
		list = append(list, v+"\n")
	}
	return list
}

func getLastKey(m map[int]string) (lastKey int) {
	for k := range m {
		if k > lastKey {
			lastKey = k
		}
	}
	return lastKey
}

// Create a Todo
func create(t *m.Todo) error {
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
func (todoTxtImpl) Update(t *m.Todo) error {
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
func (impl *todoTxtImpl) Delete(id int) error {
	cache, err := impl.getCache()
	if err != nil {
		return err
	}
	err = notInCollectionError(id)
	present := false
	_, present = cache[id]
	if present {
		delete(cache, id)
		err = nil
	} else {
		return err
	}
	list := extractTasks(cache)
	err = writeSortedLines(shared.Todotxt, list)
	return err
}

func (impl todoTxtImpl) delete(id int) error {
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
		_, err = w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
