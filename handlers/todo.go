package handlers

import (
	"encoding/json"
	"net/http"
)

// Todo represents a simple todo item
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// CreateTodo handles the creation of a new todo item
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.ID = "1" // In a real app, you would generate a unique ID
	json.NewEncoder(w).Encode(todo)
}

// EditTodo handles the editing of an existing todo item
func EditTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.Status = "Updated" // Example of editing the todo status
	json.NewEncoder(w).Encode(todo)
}
