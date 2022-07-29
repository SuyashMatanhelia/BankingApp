package models

import "time"

type Account struct {
	tableName      struct{}  `sql:"accounts"`
	AccountId      int       `json:"account_id" sql:"account_id, type:bigserial PRIMARY KEY"`
	FkCustomerId   int       `json:"fk_customer_id" sql:"fk_customer_id, type:bigint REFERENCES customer(customer_id) "`
	AccountNumber  int64     `json:"account_number" sql:"account_number, type:numeric(10) UNIQUE NOT NULL"`
	IsActive       bool      `json:"is_active" sql:"is_active, type:boolean"`
	AccountType    string    `json:"account_type" sql:"account_type, type:text NOT NULL"`
	CurrentBalance int       `json:"current_balance" sql:"current_balance, type:bigint NOT NULL"`
	FkIfscCode     string    `json:"fk_ifsc_code" sql:"fk_ifsc_code, type:varchar(11) REFERENCES branch(ifsc_code) NOT NULL"`
	CreatedAt      time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
