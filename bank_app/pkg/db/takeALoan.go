package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func TakeALoan(db *pg.DB, cust *models.Details) (n []models.Loan) {
	fmt.Println(cust, "Hello")
	var loan models.Loan
	loan.FkCustomerId = cust.CustId
	loan.LoanAmount = cust.Amt
	loan.LoanTerm = cust.Term
	loan.LoanInterest = 12
	loan.TotalInterest = (loan.LoanAmount * loan.LoanTerm * loan.LoanInterest) / 100
	loan.Installments = loan.LoanTerm * 12
	loan.MonthlyAmount = loan.TotalInterest/loan.Installments + loan.LoanAmount/loan.Installments
	loan.MonthlyInterest = loan.TotalInterest / loan.Installments

	_, err := db.Model(&loan).Insert()
	if err != nil {
		fmt.Println(err)
	}
	return
}
