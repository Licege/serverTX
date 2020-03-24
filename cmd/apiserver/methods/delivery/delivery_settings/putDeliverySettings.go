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

		err := db.QueryRow(`UPDATE delivery_settings SET city_id = $1, price_for_delivery = $2, free_delivery = $3 WHERE id = $4`,
			requestBody.CityId,
			requestBody.PriceForDelivery,
			requestBody.FreeDelivery,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
