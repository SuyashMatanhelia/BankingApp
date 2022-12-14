package models

import "time"

type Loan struct {
	tableName       struct{}  `sql:"loan"`
	LoanId          int       `sql:"loan_id, type:bigserial PRIMARY KEY"`
	FkCustomerId    int       `sql:"fk_customer_id, type:bigint REFERENCES customer(customer_id)"`
	LoanAmount      int       `sql:"loan_amount, type:bigint"`
	LoanTerm        int       `sql:"loan_term, type:int"`
	LoanInterest    int       `sql:"loan_interest, type:int"`
	TotalInterest   int       `sql:"total_interest, type:int"`
	Installments    int       `sql:"installments, type: int"`
	MonthlyAmount   int       `sql:"monthly_amount, type:int"`
	MonthlyInterest int       `sql:"monthly_interest, type:int"`
	CreatedAt       time.Time `sql:"created_at, type:timestamp NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time `sql:"updated_at, type:timestamp NOT NULL DEFAULT NOW()"`
}
