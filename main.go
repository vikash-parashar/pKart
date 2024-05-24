package main

import (
	"log"
	"net/http"
	"pkart/database"
	"pkart/router"
)

func main() {
	router := router.PkartRoutes()
	database.DbInIt()
	log.Fatal(http.ListenAndServe(":8080", router))
}
