package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetAll(res http.ResponseWriter, _ *http.Request) {
	todos, err := h.Todos.FindAll()
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(todos); err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *Handler) GetById(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(res, "Invalid ID format", http.StatusBadRequest)
		log.Println(err)
		return
	}

	todo, err := h.Todos.FindByID(id)
	if err != nil {
		http.Error(res, "Todo not found", http.StatusNotFound)
		log.Println(err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(todo); err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *Handler) Add(res http.ResponseWriter, req *http.Request) {
	var body AddTodo
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	todo, err := h.Todos.Create(&body)
	if err != nil {
		http.Error(res, "Unable to add todo", http.StatusBadRequest)
		log.Println(err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(todo); err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *Handler) Patch(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(res, "Invalid ID format", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var body UpdatedTodo
	err = json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		log.Println(err)
		return
	}

	todo, err := h.Todos.Update(id, &body)
	if err != nil {
		http.Error(res, "Unable to update todo", http.StatusBadRequest)
		log.Println(err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(todo); err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *Handler) Delete(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(res, "Invalid ID format", http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = h.Todos.Remove(id)
	if err != nil {
		http.Error(res, "Unable to delete todo", http.StatusBadRequest)
		log.Println(err)
		return
	}
}
