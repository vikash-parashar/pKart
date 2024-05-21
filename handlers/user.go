package handlers

import (
	"encoding/json"
	"net/http"
	"pkart/controllers"
	"pkart/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("content-type", "aplication/json")
	res := controllers.InsertUser(newUser)
	json.NewEncoder(w).Encode(res)

}
