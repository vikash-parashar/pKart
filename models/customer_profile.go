package models

type Customer struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	DateOfBirth string  `json:"date_of_birth"`
	MobileNo    string  `json:"mobile_no"`
	Address     Address `json:"address"`
}
