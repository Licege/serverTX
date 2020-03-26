package menu

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutDish(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requestBody := Dish{}
		c.BindJSON(&requestBody)

		_, err := db.Exec(`UPDATE dishes SET title = $1, description = $2, category_id = $3, price = $4, weight = $5, url = $6 WHERE id = $7`,
			requestBody.Title,
			requestBody.Description,
			requestBody.CategoryId,
			requestBody.Price,
			requestBody.Weight,
			requestBody.Url,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
