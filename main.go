package main

import (
	"log"
	"net/http"
	"pkart/controllers"
	"pkart/database"
)

func main() {
	router := controllers.PkartRoutes()
	database.DbInIt()
	log.Fatal(http.ListenAndServe(":8030", router))
}
