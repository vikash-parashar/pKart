package controllers

import (
	"log"
	"pkart/database"
	"pkart/models"
	"pkart/utils"
	"time"
)



func InsertCustomer(models.Customer, models.Address) {
	db := database.DbInIt()
	var newCustomer models.Customer
	var id int

utils.

err := db.QueryRow(`INSERT INTO customers (name, email, mobile, password, gender, adult, created_at)VALUES ($1, $2, $3, $4, $5, $6,$7)RETURNING id`,customer.Name, customer.Email, customer.Mobile, customer.Password, customer.Gender, customer.Adult, time.Now()).Scan(&id)



if err != nil {
	log.Fatal(err)
}


	if err != nil {
		log.Fatal(err)
	}
	// Insert Address
	createAddressTable(newCustomer.Address)
	res, err = db.Exec(`INSERT INTO addresses(customer_id,house_no,town,district,state,country,pin_code)VALUES($1, $2, $3, $4, $5, $6,$7)`, id, newCustomer.Address.HouseNo, newCustomer.Address.Town, newCustomer.Address.District, newCustomer.Address.State, newCustomer.Address.Country, newCustomer.Address.PinCode)
	if err != nil {
		log.Fatalf("error while excute create address table query %v", err)
	}
	if res != nil {
		println("Table create succesfuly")
	}

}
