package orders

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getOrders(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orders := []Order{}

		rows, err := db.Query(`select * from orders`)

		if err != nil {
			panic(err)
		}

		var id, count int
		var name, phone, date, comment string
		for rows.Next() {

			err := rows.Scan(&id, &name, &phone, &date, &count, &comment)

			if err != nil {
				panic(err)
			}

			newOrder := Order{
				Id: id,
				Name: name,
				Phone: phone,
				Date: date,
				Count: count,
				Comment: comment,
			}
			orders = append(orders, newOrder)
		}
		c.JSON(http.StatusOK, orders)
	}
}
