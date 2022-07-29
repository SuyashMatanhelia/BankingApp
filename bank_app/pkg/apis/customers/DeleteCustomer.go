package customers

import (
	dbpkg "bank_app/pkg/db"
	"bank_app/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func RemoveCustomer(db *pg.DB, r *gin.Engine) {

	var cust *models.DelCust
	r.POST("/customer/delete", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&cust); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		fmt.Println(cust)
		c.JSON(200, dbpkg.DeleteCustomer(db, cust))
	})
}
