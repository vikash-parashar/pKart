package utils

import (
	"database/sql"
	"errors"
	"log"
	"pkart/database"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	HashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Unable to generate hash password. %v", err)
		return "", err
	}
	return string(HashPass), nil
}
func AuthenticateUser(gmailID, password string) (bool, error) {
	var storedHash string
	db := database.DbInIt()
	query := "SELECT password FROM users WHERE gmail_id = $1"
	err := db.QueryRow(query, gmailID).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("invalid gmail or password")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return false, errors.New("invalid gmail or password")
	}

	return true, nil
}
