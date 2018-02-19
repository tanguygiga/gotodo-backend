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
		// /todo/
		Todo: handler.NewTodo(s.TodoDaoImpl),
	}
	log.Println("listening on " + s.Adress)
	log.Fatal(http.ListenAndServe(s.Adress, app))
}
