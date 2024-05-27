package utils

import (
	"log"
	"pkart/database"
	"pkart/models"
	"time"
)

func createAddressTable() {
	db := database.DbInIt()
	query := `CREATE TABLE  IF NOT EXISTS addresses(
		// id SERIAL PRIMARY KEY,
		customer_id INT,
		house_no VARCHAR(50) NOT NULL,
		town VARCHAR(100) NOT NULL,
		district VARCHAR(100) NOT NULL,
		state VARCHAR(100) NOT NULL,
		country  VARCHAR(100) NOT NULL,
		pin_code INT,
		FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
			)`
	res, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error while excute create address table query %v", err)
	}
	if res != nil {
		println("Table create succesfuly")
	}
}

func createCustomerTable() {
	db := database.DbInIt()
	stmt := `CREATE  TABLE IF NOT EXISTS customers(
		customer_id SERIAL PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		date_of_birth VARCHAR(20),
		mobile_no VARCHAR(255),
		created_at TIMESTAMP
    )`
	res, err := db.Exec(stmt)
	if err != nil {
		log.Fatalf("error while excute create customer table query %v", err)
	}
	if res != nil {
		println("Table create succesfuly")
	}

}
func InsertCustomer(models.Customer, models.Address) {
	db := database.DbInIt()
	var newCustomer models.Customer
	var id int
	createCustomerTable()
	query := (`INSERT INTO customers(customer_id,first_name,last_name,date_of_birth,mobile_no,created_at)VALUES($1,$2,$3,$4,$5,$6)RETURNING id`)
	err := db.QueryRow(query, newCustomer.CustomerId, newCustomer.FirstName, newCustomer.LastName, newCustomer.DateOfBirth, newCustomer.MobileNo, time.Now()).Scan(&id)
	if err != nil {
		log.Fatalf("error while excute insert customer table query %v", err)
	}
	// Insert Address

	createAddressTable()
	stmt := (`INSERT INTO customers(customer_id,first_name,last_name,date_of_birth,mobile_no,created_at)VALUES($1,$2,$3,$4,$5,$6)RETURNING id`)
	err := db.QueryRow(query, newCustomer.CustomerId, newCustomer.FirstName, newCustomer.LastName, newCustomer.DateOfBirth, newCustomer.MobileNo, time.Now()).Scan(&id)
	if err != nil {
		log.Fatalf("error while excute insert customer table query %v", err)
	}
}
