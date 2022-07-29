package models

import "time"

type Transaction struct {
	tableName               struct{}  `sql:"transactions"`
	TransactionId           int       `sql:"transaction_id,type:bigserial PRIMARY KEY"`
	FkAccountId             int       `sql:"fk_account_id, type:bigint REFERENCES account(account_id)"`
	CreditedAmount          int       `sql:"credited_amount, type:int"`
	DebitedAmount           int       `sql:"debited_amount, type:int"`
	RunningBalance          int       `sql:"running_balance, type:int"`
	OtherPartyAccountNumber int       `sql:"other_party_account_number, type:numeric(10) NOT NULL"`
	TransactionAt           time.Time `sql:"transaction_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
