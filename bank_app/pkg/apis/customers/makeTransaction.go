package customers

import (
	dbpkg "bank_app/pkg/db"
	"bank_app/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Transact(db *pg.DB, r *gin.Engine) {
	r.POST("/transact", func(c *gin.Context) {
		testone := &models.TransDetails{}
		if err := c.ShouldBindJSON(&testone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(testone)
		c.JSON(200, dbpkg.Transaction(db, testone))
	})
}
