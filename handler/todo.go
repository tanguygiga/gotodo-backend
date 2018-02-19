package handler

import (
	"encoding/json"
	"fmt"
	"gotodo-backend/dao"
	"gotodo-backend/model"
	"gotodo-backend/shared"
	"log"
	"net/http"
	"strconv"
)

const (
	contentType = "Content-Type"
	appJSON     = "application/json"
)

// Todo handle requests for Todo
type Todo struct {
	dao dao.Todo
}

// NewTodo creating handler with given implementation
func NewTodo(impl string) *Todo {
	h := &Todo{}
	h.dao = dao.TodoFactory(impl)
	return h
}

func (h *Todo) ServeHTTP(rsw http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shared.ShiftPath(req.URL.Path)
	if head == "" {
		switch req.Method {
		case http.MethodGet:
			h.getAll(rsw, req)
		case http.MethodPost:
			h.create(rsw, req)
		default:
			http.Error(rsw, fmt.Sprintf("%s is not allowed", req.Method), http.StatusMethodNotAllowed)
		}
	} else {
		id, e := strconv.Atoi(head)
		if e != nil {
			http.Error(rsw, fmt.Sprintf("Invalid todo id %q", head), http.StatusBadRequest)
			return
		}
		switch req.Method {
		case http.MethodGet:
			h.get(rsw, req, id)
		case http.MethodPut:
			h.update(rsw, req, id)
		case http.MethodDelete:
			h.delete(rsw, req, id)
		default:
			http.Error(rsw, fmt.Sprintf("%s is not allowed", req.Method), http.StatusMethodNotAllowed)
		}
	}
}

func (h *Todo) getAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	listTodo, e := h.dao.GetAll()
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	json, _ := json.Marshal(listTodo)
	w.Header().Set(contentType, appJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", json)
}

func (h *Todo) get(w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	t, e := h.dao.Get(id)
	json, _ := json.Marshal(t)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set(contentType, appJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", json)
}

func (h *Todo) create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	t := model.Todo{}
	json.NewDecoder(r.Body).Decode(&t)
	e := h.dao.Create(&t)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	json, _ := json.Marshal(t)
	w.Header().Set(contentType, appJSON)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", json)
}

func (h *Todo) update(w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	t := model.Todo{}
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = id
	e := h.dao.Update(&t)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	json, _ := json.Marshal(t)
	w.Header().Set(contentType, appJSON)
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "%s", json)
}

func (h *Todo) delete(w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	e := h.dao.Delete(id)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set(contentType, appJSON)
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "%s %d", "Removed Todo", id)
}
