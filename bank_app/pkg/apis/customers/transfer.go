package customers

import (
	dbpkg "bank_app/pkg/db"
	"bank_app/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Transfer(db *pg.DB, r *gin.Engine) {
	r.POST("/transfer", func(c *gin.Context) {
		testone := &models.TransferDetails{}
		if err := c.ShouldBindJSON(&testone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, dbpkg.Transfer(db, testone))
	})
}
