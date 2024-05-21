package handlers

import (
	"encoding/json"
	"net/http"
	"pkart/controllers"
	"pkart/models"
	"strconv"

	"github.com/gorilla/mux"
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
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
        
	paramsId := mux.Vars(r)
	id := paramsId["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	res := controllers.DeleteUserDb(userId)
	err=json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
