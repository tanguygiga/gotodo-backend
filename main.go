package main

import (
	"gotodo-backend/handler"
	"gotodo-backend/router"
	s "gotodo-backend/shared"
	"log"
	"net/http"
)

func main() {
	app := &router.App{
		Routes: map[string]http.Handler{
			"todo": handler.NewTodo(s.TodoDaoImpl),
		},
	}

	log.Println("listening on " + s.Address)
	log.Fatal(http.ListenAndServe(s.Address, app))
}
