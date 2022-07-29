package models

import "time"

type Branch struct {
	tableName     struct{}  `sql:"branch"`
	BranchId      int       `sql:"branch_id, type:bigserial PRIMARY KEY"`
	IfscCode      string    `sql:"ifsc_code, type:varchar(11) UNIQUE NOT NULL"`
	FkBankId      int       `sql:"fk_bank_id, type:bigint REFERENECES bank(bank_id)" `
	BranchName    string    `sql:"branch_name, type:text"`
	BranchAddress string    `sql:"branch_address, type:text"`
	CreatedAt     time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
