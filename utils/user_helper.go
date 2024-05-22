package utils

import (
	"log"
	"pkart/database"
	"pkart/models"

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
func GetUserEmailId(user models.User) {
	db := database.DbInIt()
	query:=`SELECT email_id WHERE `
}
