package models

type Address struct {
	HouseNo  string `json:"house_no"`
	Town     string `json:"town"`
	District string `json:"district"`
	State    string `json:"state"`
	Country  string `json:"country"`
	AreaCode string `json:"area_code"`
}
