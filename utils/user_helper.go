package utils

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"pkart/database"
	"pkart/models"
	"time"

	"github.com/dgrijalva/jwt-go"
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
func AuthenticateUser(w http.ResponseWriter, gmailId, password string) (string, error) {
	var jwtKey = []byte("pkart")
	var storedHash string
	db := database.DbInIt()
	defer db.Close()
	query := "SELECT password FROM users WHERE gmail_id = $1"
	err := db.QueryRow(query, gmailId).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid gmail or password")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		GmailId: gmailId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		err = errors.New("error : failed to create jwt token")
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    gmailId,
		Value:   tokenString,
		Expires: expirationTime,
	})
	return tokenString, nil
}
