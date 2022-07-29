package db

import (
	"bank_app/pkg/models"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func CustomerDetails(db *pg.DB, custId int) (mod models.CustBankMap) {
	n := new(models.Customer)
	db.Model(n).ColumnExpr("customer_id,first_name,last_name,pan_number,aadhar_number,dob,email,contact_number,addr,banks.bank_id,banks.bank_code,banks.bank_name,banks.bank_address").Join("LEFT JOIN banks ON fk_bank_id = banks.bank_id").Where("customer_id = ?0", custId).Select(&mod)
	fmt.Println(mod)
	return
}
