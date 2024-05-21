package router

import (
	"pkart/handlers"

	"github.com/gorilla/mux"
)

func PkartRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", handlers.CreateUserHandler).Methods("Post")
	return r
}
