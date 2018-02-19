package router

import (
	"gotodo-backend/shared"
	"net/http"
)

// App register all the Handler of the application and their routes
type App struct {
	Routes map[string]http.Handler
}

// ServeHTTP serve the corresponding handler given a path
func (a *App) ServeHTTP(rsw http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shared.ShiftPath(req.URL.Path)
	h, found := a.Routes[head]
	if found {
		h.ServeHTTP(rsw, req)
		return
	}
	http.Error(rsw, "Not Found", http.StatusNotFound)
	}
