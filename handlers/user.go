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
	w.WriteHeader(http.StatusCreated)
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "aplication/json")
	res, err := controllers.DeleteUserDb(userId)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	} else {
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

}
