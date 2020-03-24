package delivery_global_settings

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDeliveryGlobalSettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		settings := Settings{}

		var id int

		err := db.QueryRow(`SELECT * FROM delivery_global_settings WHERE id = 1`).Scan(
			&id,
			&settings.IsDeliveryWorking,
			&settings.PhoneForSms)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, settings)
	}
}
