package controllers

import (
	"database/sql"
	"log"
	"pkart/database"
	"pkart/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserTable(*sql.DB) {
	db := database.DbInIt()
	query := `CREATE TABLE IF NOT EXISTS users(
	user_id SERIAL PRIMARY KEY,
	gmail_id TEXT NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error during execute database %v", err)
	}
}

func InsertUser(newUser models.User) int {
	db := database.DbInIt()
	defer db.Close()
	CreateUserTable(db)
	var userId int
	query := (`INSERT INTO users(gmail_id,password,created_at)VALUES($1,$2,$3)RETURNING user_id`)
	HashPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Unable to insert user execute the query. %v", err)
	}
	err = db.QueryRow(query, newUser.GmailId, string(HashPass), time.Now()).Scan(&userId)
	if err != nil {
		log.Fatalf("Unable to insert user execute the query. %v", err)
	}

	return userId
}

func DeleteUserDb(userid int) (result string) {
	db := database.DbInIt()
	defer db.Close()
	_, err := db.Exec("DELETE FROM users WHERE user_id = $1", userid)
	if err != nil {
		log.Fatal(err)
	}
	return "User deleted"
}
