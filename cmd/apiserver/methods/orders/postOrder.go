package orders

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Order{}
		c.BindJSON(&requestBody)

		var id int
		err := db.QueryRow(`INSERT INTO orders (name, phone, date, count, comment, create_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
			requestBody.Name,
			requestBody.Phone,
			requestBody.Date,
			requestBody.Count,
			requestBody.Comment,
			requestBody.CreateAt).Scan(&id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, id)
	}
}
