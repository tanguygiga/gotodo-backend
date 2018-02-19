package dao

import "fmt"

// TodoFactory return a Dao with the given implementation
func TodoFactory(impl string) Todo {
	var dao Todo
	switch impl {
	case "txt":
		dao = &TodoTxtImpl{
			Tm: make(map[int]string),
		}
	default:
		dao = nil
		fmt.Print("Not yet implemented !")
	}
	return dao
}
