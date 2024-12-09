package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserId    int    `json:"user_id"`
	GmailId   string `json:"gmail_id"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type Claims struct {
	GmailId string `json:"gmail_id"`
	jwt.StandardClaims
}
