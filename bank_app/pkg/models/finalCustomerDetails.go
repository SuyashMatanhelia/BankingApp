package models

type CustBankMap struct {
	CustomerId    int    `sql:"customer_id"`
	FirstName     string `sql:"first_name"`
	LastName      string `sql:"last_name"`
	PanNumber     string `sql:"pan_number"`
	AadharNumber  int64  `sql:"aadhar_number"`
	DOB           string `sql:"dob"`
	Email         string `sql:"email"`
	ContactNumber int64  `sql:"contact_number"`
	Addr          string `sql:"addr"`
	BankId        int    `sql:"bank_id"`
	BankCode      string `sql:"bank_code"`
	BankName      string `sql:"bank_name"`
	BankAddress   string `sql:"bank_address"`
}
