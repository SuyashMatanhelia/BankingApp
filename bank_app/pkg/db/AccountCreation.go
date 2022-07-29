package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func AccountCreation(db *pg.DB, cust *models.Account) (n []models.Account) {
	fmt.Println(cust)
	_, err := db.Model(cust).Insert()
	if err != nil {
		fmt.Println(err)
	}
	return
}
