package models

type Address struct {
	HouseNo  string `json:"house_no"`
	Town     string `json:"town"`
	District string `json:"district"`
	State    string `json:"state"`
	Country  string `json:"country"`
	PinCode  string `json:"pin_code"`
}
