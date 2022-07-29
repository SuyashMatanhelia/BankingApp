package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func AccountDetails(db *pg.DB, custId int) (n []models.Account) {
	db.Model(&n).Where("account_id = ?0", custId).Select()
	fmt.Println(n)
	return
}
