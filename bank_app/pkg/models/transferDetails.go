package models

type TransferDetails struct {
	Check          int `json:"transaction_type"`
	Id             int `json:"account_id"`
	Amount         int `json:"amount"`
	Beneficiery_Id int `json:"beneficiary_account_id"`
}
