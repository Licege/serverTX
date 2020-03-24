package delivery_global_settings

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutDeliveryGlobalSettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Settings{}
		c.BindJSON(&requestBody)

		err := db.QueryRow(`UPDATE delivery_global_settings SET is_delivery_settings = $1, phone_for_sms = $2 WHERE id = 1`,
			requestBody.IsDeliveryWorking,
			requestBody.PhoneForSms)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
