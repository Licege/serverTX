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

		err := db.QueryRow(`UPDATE dishes SET title = $1, description = $2, category_id = $3, price = $4, weight = $5, file_id = $6 WHERE id = $7`,
			requestBody.Title,
			requestBody.Description,
			requestBody.Category,
			requestBody.Price,
			requestBody.Weight,
			requestBody.FileId,
			id)

		if err != nil {
			panic(err)
		}

		c.Status(http.StatusNoContent)
	}
}
