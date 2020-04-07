package delivery_settings

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDeliverySettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var settings []Settings

		rows, err := db.Query(`SELECT * FROM delivery_settings`)

		if err != nil {
			panic(err)
		}

		var id, cityId, priceForDelivery, freeDelivery int
		var isDelivery bool
		for rows.Next() {
			err := rows.Scan(&id, &cityId, &priceForDelivery, &freeDelivery, &isDelivery)

			if err != nil {
				panic(err)
			}

			setting := Settings{
				Id:		id,
				CityId:           cityId,
				PriceForDelivery: priceForDelivery,
				FreeDelivery:     freeDelivery,
				IsDelivery: 	isDelivery,
			}

			settings = append(settings, setting)
		}
		c.JSON(http.StatusOK, settings)
	}
}
