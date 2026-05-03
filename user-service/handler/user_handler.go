package handler

import (
	"encoding/json"
	"net/http"

	"user-service/config"
	"user-service/model"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetUser(w, r)

	case "POST":
		CreateUser(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// read JSON body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// insert into MySQL
	query := "INSERT INTO users (id, name) VALUES (?, ?)"
	_, err = config.DB.Exec(query, user.ID, user.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	var user model.User

	query := "SELECT id, name FROM users WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(&user.ID, &user.Name)

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
