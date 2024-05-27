package router

import (
	"pkart/handlers"

	"github.com/gorilla/mux"
)

func PkartRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/deletebyuserid/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user/login", handlers.UserLogin).Methods("GET")
	// r.HandleFunc("/user/getbygmail/{gmail}", handlers.GetUerByEmail).Methods("GET")

	// Customer Routes

	r.HandleFunc("/customer/create/profile", handlers.CreateCustomerProfile).Methods("POST")

	return r
}
