package customers

import (
	dbpkg "bank_app/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)
type AccDet struct{

	Id int `json:"id"`

}
func GetAccountDetails(db *pg.DB, r *gin.Engine) {
	// var Det AccDet
	r.POST("/customer/account", func(c *gin.Context) {
		Det := &AccDet{}
		if err := c.ShouldBindJSON(&Det); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		c.JSON(200, dbpkg.AccountDetails(db, Det.Id))
	})
}
