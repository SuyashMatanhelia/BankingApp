package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func Transaction(db *pg.DB, cust *models.TransDetails) (n models.Transaction) {
	fmt.Println(cust, "Hello")
	var account1 models.Account
	var trans1 models.Transaction
	trans1.FkAccountId = cust.Id
	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Printf("Error while opening tx for transaction, reason : %v\n", txErr.Error())
	}
	tx.Model(&account1).ColumnExpr("current_balance, account_number").Where("account_id = ?0", cust.Id).Select()
	fmt.Println(account1)
	switch cust.Check {
	case 1:
		trans1.RunningBalance = account1.CurrentBalance - cust.Amount
		trans1.DebitedAmount = cust.Amount
	case 2:
		trans1.RunningBalance = account1.CurrentBalance + cust.Amount
		trans1.CreditedAmount = cust.Amount
	}
	if trans1.RunningBalance < 0 {
		fmt.Println("Balance too low for transaction")
		return trans1
	}
	trans1.OtherPartyAccountNumber = cust.Id
	fmt.Println(trans1)
	_, err := tx.Model(&trans1).Insert()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&account1).Set("current_balance = ?", trans1.RunningBalance).Where("account_id = ?", cust.Id).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
	tx.Commit()
	return
}
