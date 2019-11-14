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
		err := db.QueryRow("INSERT INTO dishes (title, description, weight, price, category_id, file_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		requestBody.Title,
		requestBody.Description,
		requestBody.Weight,
		requestBody.Price,
		requestBody.Category,
		requestBody.FileId).Scan(&lastID)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, lastID)
	}
}
