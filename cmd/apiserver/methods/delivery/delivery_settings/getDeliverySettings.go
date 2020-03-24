package delivery_settings

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDeliverySettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		settings := []Settings{}

		rows, err := db.Query(`SELECT * FROM delivery_settings`)

		if err != nil {
			panic(err)
		}

		var id, cityId, priceForDelivery, freeDelivery int
		for rows.Next() {
			err := rows.Scan(&id, &cityId, &priceForDelivery, &freeDelivery)

			if err != nil {
				panic(err)
			}

			setting := Settings{
				CityId:           cityId,
				PriceForDelivery: priceForDelivery,
				FreeDelivery:     freeDelivery,
			}

			settings = append(settings, setting)
		}
		c.JSON(http.StatusOK, settings)
	}
}