package customers

import (
	dbpkg "bank_app/pkg/db"
	"bank_app/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func PutCustomer(db *pg.DB, r *gin.Engine) {

	r.POST("/customer/create", func(c *gin.Context) {
		testone := &models.Customer{}
		if err := c.ShouldBindJSON(&testone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		fmt.Println(testone)
		c.JSON(200, dbpkg.CreateCustomer(db, testone))
	})
}
