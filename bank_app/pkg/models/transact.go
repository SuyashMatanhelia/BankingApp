package models

type TransDetails struct {
	Check  int `json:"transaction_type"`
	Id     int `json:"account_id"`
	Amount int `json:"amount"`
}
