package controllers

import (
	"net/http"
	"pkart/handlers"

	"github.com/gorilla/mux"
)

func PkartRoutes() *mux.Router {
	router := mux.NewRouter()
	http.HandleFunc("/user/create", handlers.CreateUerHandler)
	return router
}
