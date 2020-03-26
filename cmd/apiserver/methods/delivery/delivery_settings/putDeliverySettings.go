package delivery_settings

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutSettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := Settings{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE delivery_settings SET city_id = $1, price_for_delivery = $2, free_delivery = $3, is_delivery = $4 WHERE id = $5`,
			requestBody.CityId,
			requestBody.PriceForDelivery,
			requestBody.FreeDelivery,
			requestBody.IsDelivery,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
