package models

type User struct {
	UserId    int    `json:"user_id"`
	GmailId   string `json:"gmail_id"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
