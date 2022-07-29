package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func DeleteCustomer(db *pg.DB, cust *models.DelCust) (n []models.Customer) {
	fmt.Println(cust.I)
	var c models.Customer
	var a models.Account
	var l models.Loan
	var t models.Transaction
	db.Model(&a).Where("fk_customer_id = ?", cust.I).Delete()
	db.Model(&l).Where("fk_customer_id = ?", cust.I).Delete()
	db.Model(&t).Where("fk_customer_id = ?", cust.I).Delete()
	db.Model(&c).Where("customer_id = ?", cust.I).Delete()

	return
}
