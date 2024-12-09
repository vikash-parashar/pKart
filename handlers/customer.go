package handlers

import (
	"encoding/json"
	"net/http"
	"pkart/models"
	"pkart/utils"
)

func CreateCustomerProfile(w http.ResponseWriter, r *http.Request) {
	var newCustomer models.Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	customerId := utils.InsertCustomer(newCustomer,newCustomer.Address)
	err = json.NewEncoder(w).Encode(customerId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
