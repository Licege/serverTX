package menu

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostDish(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Dish{}
		err := c.BindJSON(&requestBody)

		if err != nil {
			panic(err.Error())
		}

		var lastID int
		err = db.QueryRow("INSERT INTO dishes (title, description, category_id, price, weight, file_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		requestBody.Title,
		requestBody.Description,
		requestBody.CategoryId,
		requestBody.Price,
		requestBody.Weight,
		requestBody.File.Id).Scan(&lastID)

		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusOK, lastID)
	}
}
