package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func LoanDetails(db *pg.DB, custId int) (n []models.Loan) {
	db.Model(&n).Where("fk_customer_id = ?0", custId).Select()
	fmt.Println(n)
	return
}
