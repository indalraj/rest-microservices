package handler

import (
	"encoding/json"
	"net/http"

	"user-service/model"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{ID: "1", Name: "Rahul"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
