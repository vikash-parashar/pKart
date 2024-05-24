package handlers

import (
	"encoding/json"
	"net/http"
	"pkart/controllers"
	"pkart/models"
	"pkart/utils"
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

// func GetUerByEmail(w http.ResponseWriter, r *http.Request) {
// 	paramsId := mux.Vars(r)
// 	gmailID := paramsId["id"]
// 	user := controllers.GetUserbyGmail(gmailID)
//   if err != nil {
//   	http.Error(w, "no user found with Gmail ID:", http.StatusNotFound)
//   }
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("content-type", "aplication/json")
// 	json.NewEncoder(w).Encode(user)

// }
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	authenticated, err := utils.AuthenticateUser(user.GmailId, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if authenticated {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	} else {
		http.Error(w, "Invalid gmail or password", http.StatusUnauthorized)
	}

}
