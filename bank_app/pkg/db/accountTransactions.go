package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func AccountTransactions(db *pg.DB, custId int) (n []models.Transaction) {
	db.Model(&n).Where("fk_account_id = ?", custId).Select()
	fmt.Println(n)
	return
}
