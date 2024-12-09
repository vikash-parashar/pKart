package utils

import (
	"fmt"
	"log"

	"pkart/database"
	"pkart/models"
	"time"
)

func createAddressTable() {
	db := database.DbInIt()
	query := `CREATE TABLE  IF NOT EXISTS addresses(
	customer_id INT REFERENCES customers(customer_id),
	house_no VARCHAR(50),
	town VARCHAR(50),
	district VARCHAR(50),
	state VARCHAR(50),
	country VARCHAR(50),
	pin_code VARCHAR(10)
	
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error while excute create address table query %v", err)
	}

}

func createCustomerTable() {
	db := database.DbInIt()
	query := `CREATE  TABLE IF NOT EXISTS customers(
		customer_id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		date_of_birth VARCHAR(20),
		mobile_no VARCHAR(15),
		created_at TIMESTAMP
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error while excute create customer table query %v", err)
	}
}

func InsertCustomer(new models.Customer, addr models.Address) int {
	db := database.DbInIt()
	defer db.Close()
	// var newCustomer models.Customer
	var id int
	createCustomerTable()
	query := `INSERT INTO customers(first_name,last_name,date_of_birth,mobile_no,created_at)VALUES($1,$2,$3,$4,$5)RETURNING customer_id`
	err := db.QueryRow(query, new.FirstName, new.LastName, new.DateOfBirth, new.MobileNo, time.Now()).Scan(&id)
	if err != nil {
		log.Fatalf("error while excute insert customer table query %v", err)
	}
	// Insert Address

	createAddressTable()
	sqlStatement := `INSERT INTO addresses (customer_id, house_no, town, district, state, country, pin_code)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Exec(sqlStatement, id, addr.HouseNo, addr.Town, addr.District, addr.State, addr.Country, addr.PinCode)
	if err != nil {
		log.Fatalf("error while excute insert customer table query %v", err)

	}
	fmt.Println("Customer Details:", new)
	fmt.Println("Address Details:", addr)
	return id

}
