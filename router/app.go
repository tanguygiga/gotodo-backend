package router

import (
	"gotodo-backend/handler"
	"gotodo-backend/shared"
	"net/http"
)

// App register all the Handler of the application and their routes
type App struct {
	Todo *handler.Todo
}

func (h *App) ServeHTTP(rsw http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shared.ShiftPath(req.URL.Path)
	switch head {
	case "todo":
		h.Todo.ServeHTTP(rsw, req)
	default:
		http.Error(rsw, "Not Found", http.StatusNotFound)
	}
}
