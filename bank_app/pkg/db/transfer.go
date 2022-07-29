package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func Transfer(db *pg.DB, cust *models.TransferDetails) (n models.Transaction) {
	var account1 models.Account
	var trans1 models.Transaction
	var account2 models.Account
	var trans2 models.Transaction
	trans1.FkAccountId = cust.Id
	trans2.FkAccountId = cust.Beneficiery_Id
	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Printf("Error while opening tx for transaction, reason : %v\n", txErr.Error())
	}
	tx.Model(&account1).ColumnExpr("current_balance, account_number").Where("account_id = ?0", cust.Id).Select()
	tx.Model(&account2).ColumnExpr("current_balance,  account_number").Where("account_id = ?0", cust.Beneficiery_Id).Select()
	fmt.Println(account1)
	fmt.Println(account2)
	trans1.RunningBalance = account1.CurrentBalance - cust.Amount
	trans2.RunningBalance = account1.CurrentBalance + cust.Amount

	trans1.DebitedAmount = cust.Amount
	trans2.CreditedAmount = cust.Amount

	trans1.OtherPartyAccountNumber = cust.Beneficiery_Id
	trans2.OtherPartyAccountNumber = cust.Id
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
	trans1.OtherPartyAccountNumber = cust.Beneficiery_Id
	fmt.Println(trans1)
	_, err := tx.Model(&trans1).Insert()
	fmt.Println(trans1)
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	_, err = tx.Model(&account1).Set("current_balance = ?", trans1.RunningBalance).Where("account_id = ?", cust.Id).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
	_, err = tx.Model(&account2).Set("current_balance = ?", trans2.RunningBalance).Where("account_id = ?", cust.Beneficiery_Id).Update()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
	tx.Commit()
	return
}
