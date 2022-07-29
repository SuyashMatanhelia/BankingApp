package customers

import (
	dbpkg "bank_app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func GetCustomerDetails(db *pg.DB, r *gin.Engine) {
	r.GET("/customer/details", func(c *gin.Context) {
		Det := &AccDet{}
		if err := c.ShouldBindJSON(&Det); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		c.JSON(200, dbpkg.CustomerDetails(db, Det.Id))
	})
}
