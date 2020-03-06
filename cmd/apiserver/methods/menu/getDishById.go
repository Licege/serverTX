package menu

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDishById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		dish := Dish{}

		db.QueryRow(`SELECT * FROM dishes WHERE id = $1`, id).Scan(
			&dish.Id,
			&dish.Title,
			&dish.Description,
			&dish.CategoryId,
			&dish.Price,
			&dish.Weight,
			&dish.Url)

		c.JSON(http.StatusOK, dish)
	}
}
