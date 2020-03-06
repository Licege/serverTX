package menu

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostDish(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Dish{}
		c.BindJSON(&requestBody)

		var lastID int
		err := db.QueryRow("INSERT INTO dishes (title, description, category_id, price, weight, url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		requestBody.Title,
		requestBody.Description,
		requestBody.CategoryId,
		requestBody.Price,
		requestBody.Weight,
		requestBody.Url).Scan(&lastID)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, lastID)
	}
}
