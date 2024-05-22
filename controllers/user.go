package controllers

import (
	"database/sql"
	"errors"
	"log"
	"pkart/database"
	"pkart/models"
	"pkart/utils"
	"time"
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
	HashPass, err := utils.GenerateHashPassword(newUser.Password)
	if err!=nil{
		log.Fatalf("Unable to generate hash password: %V",err)
	}
	err = db.QueryRow(query, newUser.GmailId, string(HashPass), time.Now()).Scan(&userId)
	if err != nil {
		log.Fatalf("Unable to insert user execute the query. %v", err)
	}

	return userId
}

func DeleteUserDb(userid int) (result string, err error) {
	db := database.DbInIt()
	defer db.Close()

	res, err := db.Exec("DELETE FROM users WHERE user_id = $1", userid)
	if err != nil {
		return "", err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("user ID does not exist")
	}

	return "User deleted", nil
}
