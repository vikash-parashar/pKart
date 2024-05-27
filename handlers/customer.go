package handlers

import (
	"encoding/json"
	"net/http"
	"pkart/controllers"
	"pkart/models"
)

func CreateCustomerProfile(w http.ResponseWriter, r *http.Request) {
	var newCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	customer := controllers.InsertCustomer(newCustomer)
}
