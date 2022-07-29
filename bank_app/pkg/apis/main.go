package apis

import (
	"bank_app/pkg/apis/customers"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func ConnectDB() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "password",
		Database: "bankdb",
	})
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	StartServer(db)
	defer db.Close()
}
func StartServer(db *pg.DB) {
	router := gin.Default()
	customers.GetCustomerDetails(db, router)
	customers.GetAccountDetails(db, router)
	customers.GetLoanDetails(db, router)
	customers.GetAccountTransactions(db, router)
	customers.PutCustomer(db, router)
	customers.TakeALoan(db, router)
	customers.RemoveCustomer(db, router)
	customers.Transact(db, router)
	customers.Transfer(db, router)
	customers.CreateAccount(db, router)
	router.Run(":8080")
}
