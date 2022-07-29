package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func CreateCustomer(db *pg.DB, cust *models.Customer) (n []models.Account) {
	fmt.Println(cust)
	_, err := db.Model(cust).Insert()
	if err != nil {
		fmt.Println(err)
	}
	return
}
